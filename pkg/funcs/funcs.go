package funcs

import "github.com/nkanaev/numb/pkg/value"

type Func func(...value.Value) value.Value

var db = map[string]Func{
	"sin": Sin,
}

func Get(x string) *Func {
	if fn, ok := db[x]; ok {
		return &fn
	}
	return nil
}
