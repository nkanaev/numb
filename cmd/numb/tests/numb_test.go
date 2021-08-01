package tests

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nkanaev/numb/pkg/parser"
	"github.com/nkanaev/numb/pkg/value"
)

type Spec struct {
	Name  string
	Exprs []string
	Wants []string
}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func listSpecs() []Spec {
	files, err := filepath.Glob("*.txt")
	if err != nil {
		log.Fatal(err)
	}
	specs := make([]Spec, 0)
	spec := Spec{}
	for _, file := range files {
		lines := readLines(file)
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if len(line) == 0 {
				continue
			}
			if strings.HasPrefix(line, "#") {
				if spec.Name != "" {
					specs = append(specs, spec)
				}
				spec = Spec{Name: file + ":" + line[2:]}
			} else {
				chunks := strings.Split(line, "|")
				var expr, want string
				log.Println("$$", chunks)
				if len(chunks) == 2 {
					expr = strings.TrimSpace(chunks[0])
					want = strings.TrimSpace(chunks[1])
				} else if len(chunks) == 1 {
					expr = strings.TrimSpace(chunks[0])
				} else {
					log.Fatal("failed to parse:", line)
				}
				spec.Exprs = append(spec.Exprs, expr)
				spec.Wants = append(spec.Wants, want)
			}
		}	
	}
	specs = append(specs, spec)
	return specs
}

func TestSpecs(t *testing.T) {
	for _, spec := range listSpecs() {
		spec := spec
		t.Run(spec.Name, func(t *testing.T) {
			env := make(map[string]value.Value)
			for i := 0; i < len(spec.Exprs); i++ {
				expr := spec.Exprs[i]
				want := spec.Wants[i]
				have, err := parser.Eval(expr, env)

				if err != nil && (len(want) == 0 || want[0] != '!') {
					t.Fatalf("unexpected error\nexpr: %s\n err: %s", expr, err.Error())
				}
				if len(want) == 0 {
					continue
				}
				if want[0] == '!' {
					if err == nil || !strings.Contains(err.Error(), want[2:]) {
						t.Fatalf("expected error\nexpr: %s\n err: %s", expr, err)
					}
				} else {
					have := have.Format(",", 2)
					if have != want {
						t.Logf("%+v", spec)
						t.Fatalf("wrong answer\nexpr: %s\nwant: %s\nhave: %s", expr, want, have)
					}
				}
			}
		})
	}
}
