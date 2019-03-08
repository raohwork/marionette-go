// This file is part of marionette-go
//
// marionette-go is distributed in two licenses: The Mozilla Public License,
// v. 2.0 and the GNU Lesser Public License.
//
// marionette-go is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE.
//
// See License.txt for further information.

package mnclient

import (
	"sync"
)

// ContextManager helps you to ensure your code is running in correct context
//
// It ensures codes between Enter() and Leave() will be executed in selected
// context.
//
// By using ContextManager, you MUST NOT SWITCH TO OTHER CONTEXT in your code.
type ContextManager interface {
	// Enter tries to acquire the lock, blocks until context is switched to
	// correct one
	Enter(string) error
	// Leave releases the lock
	Leave()
}

// NewSharedContext creates a ContextManager letting goroutines share same context
//
// It manages a virtual shared lock (something like R-part in sync.RWMutex). User
// requests entering same context with current one can acquire the lock without
// being blocked.
//
// As it stores current context in local variable, IT CANNOT COOPERATE WITH ANY
// OTHER CONTEXT SWITCHING CALLS, including plain mnclient.Commander.MozSetContext()
// call.
//
// Concurrency
//
// It works well ONLY IN MANAGED ENVIRONMENT. In random, heavy loading environment,
// there would be great possibility that some goroutines getting blocked for long
// time, and even draining your resources til OOM in extreme case.
//
// See test case for example.
func NewSharedContext(cl *Commander) (ret ContextManager, err error) {
	ctx, err := cl.MozGetContext()
	if err != nil {
		return
	}

	s := &sharedContext{
		cl:      cl,
		current: ctx,
	}
	s.cond = sync.NewCond(&s.lock)
	ret = s
	return
}

type sharedContext struct {
	cl      *Commander
	current string
	lock    sync.Mutex
	cond    *sync.Cond
	running int
}

func (ctx *sharedContext) Enter(c string) (err error) {
	ctx.lock.Lock()
	defer ctx.lock.Unlock()

	for ctx.current != c && ctx.running > 0 {
		ctx.cond.Wait()
	}

	err = ctx.toggle(c)
	if err != nil {
		return
	}

	ctx.running++
	return
}

func (ctx *sharedContext) Leave() {
	ctx.lock.Lock()
	defer ctx.lock.Unlock()
	ctx.running--

	ctx.cond.Broadcast()
}

func (ctx *sharedContext) toggle(c string) (err error) {
	if ctx.current == c {
		return
	}

	err = ctx.cl.MozSetContext(c)
	if err != nil {
		return
	}

	ctx.current = c
	ctx.cond.Broadcast()
	return
}
