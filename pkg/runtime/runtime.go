package runtime

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/nkanaev/numb/pkg/parser"
	"github.com/nkanaev/numb/pkg/value"
)

//go:embed builtin.txt
var builtin string

type Runtime struct {
	Prec int
	Tsep string
	Env  map[string]value.Value
}

func NewRuntime() *Runtime {
	return &Runtime{
		Prec: 2,
		Tsep: ",",
		Env:  make(map[string]value.Value),
	}
}

func Clean(line string) string {
	return strings.SplitN(line, "|", 2)[0]
}

func (r *Runtime) Eval(line string) (string, error) {
	line = Clean(line)
	
	trimmedline := strings.TrimSpace(line)
	if len(trimmedline) == 0 {
		return line, nil
	}

	firstchar := trimmedline[0]
	switch firstchar {
	case '.':
		return r.EvalConfig(line[1:])
	case '#':
		return line, nil
	default:
		val, err := parser.Eval(line, r.Env)
		if err != nil {
			return "", err
		}
		return val.String(), nil //val.Format(r.Tsep, r.Prec), nil
	}
}

func (r *Runtime) EvalConfig(line string) (string, error) {
	parts := strings.SplitN(line, " ", 2)
	cmd := parts[0]
	switch cmd {
	case "load":
		for _, path := range parts[1:] {
			r.LoadFile(path)
		}
	case "prec":
		if len(parts) != 2 {
			return strconv.Itoa(r.Prec), nil
		}
		prec, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", err
		}
		r.Prec = prec
	case "tsep":
		r.Tsep = ","
	case "notsep":
		r.Tsep = ""
	}
	return "", nil
}

func (r *Runtime) LoadBuiltins() {
	r.Load(strings.NewReader(builtin), "<builtin>")
}

func (r *Runtime) LoadFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fail("loadfile: %s", err.Error())
	}
	defer file.Close()
	r.Load(file, path)
}

func (r *Runtime) Load(reader io.Reader, filename string) {
	scanner := bufio.NewScanner(reader)
	linenum := 0
	for scanner.Scan() {
		linenum += 1
		_, err := r.Eval(scanner.Text())
		if err != nil {
			fail("load %s (line %d): %s", filename, linenum, err)
		}
	}
}

func fail(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format, v...)
	os.Stderr.WriteString("\n")
	os.Exit(1)
}
