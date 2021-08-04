package runtime

import (
	"bufio"
	_ "embed"
	"io"
	"log"
	"strings"

	"github.com/nkanaev/numb/pkg/parser"
	"github.com/nkanaev/numb/pkg/value"
)

//go:embed builtin.txt
var builtin string

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

func Clean(line string) string {
	return strings.TrimSpace(strings.SplitN(line, "|", 2)[0])
}

func (r *Runtime) Eval(line string) (string, error) {
	line = Clean(line)
	if line == "" {
		return "", nil
	}
	
	val, err := parser.Eval(Clean(line), r.Env)
	out := ""
	if err != nil {
		return "", err	
	} else {
		if val.Fmt == value.DEC {
			out = val.Format(r.Sep, r.Prec)
		} else {
			out = val.String()
		}
	}
	return out, nil
}

func (r *Runtime) LoadBuiltins() {
	r.Load(strings.NewReader(builtin), "<builtin>")
}

func (r *Runtime) Load(reader io.Reader, filename string) {
	scanner := bufio.NewScanner(reader)
	linenum := 0
	for scanner.Scan() {
		linenum += 1
		line := Clean(scanner.Text())
		if line == "" {
			continue
		}
		_, err := parser.Eval(line, r.Env)
		if err != nil {
			log.Fatalf("load %s (line %d): %s", filename, linenum, err)
		}
	}
}
