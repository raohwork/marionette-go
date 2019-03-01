package marionette

import (
	"encoding/json"
	"strconv"
	"testing"
)

func TestRefEncode(t *testing.T) {
	cases := map[string]ReftestRefList{
		`[["A",[["B",[],"!="]],"=="]]`: (ReftestRefList{}).
			Or("B", "!=").And("A", "=="),
		`[["A",[["B",[],"!="],["C",[],"!="]],"=="]]`: (ReftestRefList{}).
			Or("B", "!=").Or("C", "!=").And("A", "=="),
		`[["A",[["B",[],"!="]],"=="],["D",[],"=="]]`: (ReftestRefList{}).
			Or("B", "!=").And("A", "==").Or("D", "=="),
	}

	cnt := 0
	for expected, rules := range cases {
		cnt++
		t.Run("#"+strconv.Itoa(cnt), func(t *testing.T) {
			data, err := json.Marshal(rules)
			if err != nil {
				t.Fatalf("cannot marshal rule: %s", err)
			}
			str := string(data)
			if str != expected {
				t.Fatalf("unexpected result: %s", str)
			}
		})
	}
}
