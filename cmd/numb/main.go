package main

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/nkanaev/numb/pkg/parser"
	"github.com/nkanaev/numb/pkg/unit"
	"github.com/nkanaev/numb/pkg/value"
	"golang.org/x/term"
)

const CURRENCYFILE = "currency.txt"

var prompt = "> "
var prefix = "  "

var rates = ""
var sep = ","
var prec = 2

func eval(expr string, env map[string]value.Value) (val value.Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown error")
			}
		}
	}()
	val = parser.Parse(expr).Eval(env)
	return
}

func repl() {
	env := make(map[string]value.Value)

	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
			panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), state)

	screen := struct{
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}
	terminal := term.NewTerminal(screen, prompt)
	terminal.Write([]byte("enter `q` to quit\n"))
	for {
		line, err := terminal.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			os.Stderr.WriteString(line + "\n")
			os.Exit(1)
		}
		line = strings.TrimSpace(line)
		if line == "q" {
			break
		}
		if line == "" {
			continue
		}
		val, err := eval(line, env)
		out := ""
		if err != nil {
			out = err.Error()
		} else {
			if val.Fmt == value.DEC {
				out = val.Format(sep, prec)
			} else {
				out = val.String()
			}
		}
		terminal.Write([]byte(prefix + out + "\n"))
	}
}

func loadCurrencies(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
	defer file.Close()

	currencies := make([]unit.Currency, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), "\t", " ")
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			continue
		}
		code := parts[0]
		rate, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			continue
		}
		currencies = append(currencies, unit.Currency{code, rate})
	}
	unit.AddExchangeRates(currencies)
}

func main() {
	flag.IntVar(&prec, "prec", prec, "decimal precision")
	flag.StringVar(&sep, "sep", sep, "thousand separator")
	flag.StringVar(&rates, "rates", "", "path to exchange rates file")
	flag.Parse()

	if rates != "" {
		loadCurrencies(rates)
	}

	repl()
}
