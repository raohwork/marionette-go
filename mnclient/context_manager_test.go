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
	"errors"
	"reflect"
	"sync"
	"testing"
	"time"

	marionette "github.com/raohwork/marionette-go/v3"
	"github.com/raohwork/marionette-go/v3/mncmd"
)

type fakeSenderForCtx struct {
	ctx  string
	sets int
	gets int
}

func (cl *fakeSenderForCtx) Start() (err error) {
	return
}
func (cl *fakeSenderForCtx) Close() {}
func (cl *fakeSenderForCtx) Wait()  {}

func (cl *fakeSenderForCtx) Async(cmd mncmd.Command) (ch chan *marionette.Message, err error) {
	ch = make(chan *marionette.Message, 1)
	switch c := cmd.(type) {
	case *mncmd.MozGetContext:
		ch <- &marionette.Message{
			Type:   1,
			Serial: 1,
			Data:   map[string]string{"Value": cl.ctx},
		}
		cl.gets++
	case *mncmd.MozSetContext:
		cl.ctx = c.Context
		ch <- &marionette.Message{
			Type:   1,
			Serial: 1,
		}
		cl.sets++
	default:
		panic(errors.New("wtf"))
	}

	return
}

func (cl *fakeSenderForCtx) Sync(cmd mncmd.Command) (msg *marionette.Message, err error) {
	ch, _ := cl.Async(cmd)
	msg = <-ch
	return
}

func TestSharedContext(t *testing.T) {
	cl := &fakeSenderForCtx{
		ctx: "content",
	}

	// can't be err as we use fake clent which doesn't return error
	ctx, _ := NewSharedContext(&Commander{Sender: cl})

	// 5+5 workers, generate 20 messages
	wg := &sync.WaitGroup{}
	wg.Add(10)
	logs := make(chan string, 20)
	workIn := func(target string) {
		ctx.Enter(target)
		logs <- target
		// do some work for 0.5s
		time.Sleep(500 * time.Millisecond)
		ctx.Leave()
	}
	jobA := func() {
		workIn("chrome")
		workIn("content")
		wg.Done()
	}
	jobB := func() {
		workIn("content")
		workIn("chrome")
		wg.Done()
	}

	// first job kicks in (type A)
	go jobA()

	// 0.1s later, five more jobs (type B)
	time.Sleep(100 * time.Millisecond)
	for x := 0; x < 5; x++ {
		go jobB()
	}

	// another four jobs (type A), each wait 0.2s
	//
	// what if
	for x := 0; x < 4; x++ {
		time.Sleep(200 * time.Millisecond)
		go jobA()
	}

	expected := []string{
		// A1~5 in random order
		"chrome", "chrome", "chrome", "chrome", "chrome",
		// all jobs in random order
		"content", "content", "content", "content", "content",
		"content", "content", "content", "content", "content",
		// B1~5 in random order
		"chrome", "chrome", "chrome", "chrome", "chrome",
	}
	actual := make([]string, 0, 20)
	wg.Wait()
	close(logs)
	for str := range logs {
		actual = append(actual, str)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("unexpected order: %+v", actual)
	}

	// expected 3 MozSetContext calls
	// - first chrome worker switch from content to chrome
	// - all chrome done first part job, switch to content
	// - chrome worker done, content worker finished first part, back to chrome
	if cl.sets != 3 {
		t.Errorf("expected 3 MozSetContext calls, got %d", cl.sets)
	}

	// expected 1 MozGetContext call
	if cl.gets != 1 {
		t.Errorf("expected 1 MozGetContext calls, got %d", cl.gets)
	}
}
