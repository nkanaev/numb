package runtime

import (
	"github.com/nkanaev/numb/pkg/parser"
	"github.com/nkanaev/numb/pkg/value"
)

type Runtime struct {
	Prec int
	Sep string
	Env map[string]value.Value
}

func NewRuntime() *Runtime {
	return &Runtime{
		Prec: 2,
		Sep: ",",
		Env: make(map[string]value.Value),
	}
}

func (r *Runtime) Eval(line string) string {
	val, err := parser.Eval(line, r.Env)
	out := ""
	if err != nil {
		out = err.Error()
	} else {
		if val.Fmt == value.DEC {
			out = val.Format(r.Sep, r.Prec)
		} else {
			out = val.String()
		}
	}
	return out
}
