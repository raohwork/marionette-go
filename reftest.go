package marionette

import "encoding/json"

const (
	// reftest screenshotoptions
	ReftestAlwaysScreenshot     = "always"
	ReftestScreenshotOnFailed   = "failed"
	ReftestScreenshotUnexpected = "unexpected"
)

const (
	// possible relations for reftest:run
	RelEQ = "=="
	RelNE = "!="
)

const (
	// possible expected result of reftest:run
	ReftestPass = "PASS"
	ReftestFail = "FAIL"
)

// RefestRefList represents "OR" relation between references
type ReftestRefList []*reftestRef

// And creates a new list that represnets "uri AND l"
func (l ReftestRefList) And(uri, rel string) (ret ReftestRefList) {
	return []*reftestRef{
		&reftestRef{
			URL: uri,
			Ref: l,
			Rel: rel,
		},
	}
}

// Or creates a new ReftestRefList represents "l OR uri"
func (l ReftestRefList) Or(uri, rel string) (ret ReftestRefList) {
	return append(l, &reftestRef{
		URL: uri,
		Rel: rel,
	})
}

// reftestRef represnets a reference
type reftestRef struct {
	URL string
	Ref ReftestRefList
	Rel string
}

func (r reftestRef) MarshalJSON() (ret []byte, err error) {
	data := []interface{}{
		r.URL,
		r.Ref,
		r.Rel,
	}

	if len(r.Ref) == 0 {
		data[1] = []interface{}{}
	}

	return json.Marshal(data)
}

// ReftestResult represents the result of executed test
type ReftestResult struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Stack   string                 `json:"stack"`
	Extra   map[string]interface{} `json:"extra"`
}
