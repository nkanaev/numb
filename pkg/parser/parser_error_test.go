package parser

import (
	"strings"
	"testing"
)

type poserror interface {
	error
	Pos() (int, int)
}

var errortestcases = map[string]struct {
	expr, err  string
	start, end int
}{
	"lsh": {
		expr: "1 <! 2", err: "expected <", start: 4,
	},
	"rsh": {
		expr: "1 >= 2", err: "expected >", start: 4,
	},
	"dangling_curly": {
		expr: "  {123", err: "date :: }", start: 3,
	},
	"notident": {
		expr: "!varname", err: "unexpected", start: 1,
	},
	"notassign": {
		expr: "x === 2", err: "unexpected", start: 4,
	},
}

func TestParserErrors(t *testing.T) {
	for name, tc := range errortestcases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			node, err := Parse(tc.expr)
			if err == nil {
				t.Fatalf("expected error, got result: %#v", node)
			}
			poserr, ok := err.(poserror)
			if !ok {
				t.Fatalf("error missing pos info: %s", err.Error())
			}
			start, end := poserr.Pos()

			// NOTE: error positions are 0-indexed, but tests use 1-indexed
			start, end = start+1, end+1
			if tc.start != 0 && tc.start != start {
				t.Fatalf(
					"error start position does not match\n"+
						"error: %s\nexpr: %#v\nwant: %d\nhave: %d",
					err, tc.expr, tc.start, start,
				)
			}
			if tc.end != 0 && tc.end != end {
				t.Fatalf(
					"error start position does not match\n"+
						"error: %s\nexpr: %s\nwant: %d\nhave: %d",
					err, tc.expr, tc.end, end,
				)
			}
			for _, mention := range strings.Split(tc.err, "::") {
				mention := strings.TrimSpace(mention)
				if strings.Index(err.Error(), mention) == -1 {
					t.Fatalf("error `%s` does not mention `%s`", err, mention)
				}
			}
		})
	}
}
