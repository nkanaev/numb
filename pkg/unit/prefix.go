package unit

import (
	"math/big"

	r "github.com/nkanaev/numb/pkg/ratutils"
)

type prefix struct {
	abbr  string
	name  string
	value *big.Rat
}

var metricPrefixes = []prefix{
	{"d", "deci", r.Exp(10, -1)},
	{"c", "centi", r.Exp(10, -2)},
	{"m", "milli", r.Exp(10, -3)},
	{"u", "micro", r.Exp(10, -6)},
	{"n", "nano", r.Exp(10, -9)},
	{"p", "pico", r.Exp(10, -12)},
	{"f", "femto", r.Exp(10, -15)},
	{"a", "atto", r.Exp(10, -18)},
	{"z", "zepto", r.Exp(10, -21)},
	{"y", "yocto", r.Exp(10, -24)},

	{"da", "deca", r.Exp(10, 1)},
	{"h", "hecto", r.Exp(10, 2)},
	{"k", "kilo", r.Exp(10, 3)},
	{"M", "mega", r.Exp(10, 6)},
	{"G", "giga", r.Exp(10, 9)},
	{"T", "tera", r.Exp(10, 12)},
	{"P", "peta", r.Exp(10, 15)},
	{"E", "exa", r.Exp(10, 18)},
	{"Z", "zetta", r.Exp(10, 21)},
	{"Y", "yotta", r.Exp(10, 24)},
}

var metricPrefixesTonne = []prefix{
	{"k", "kilo", r.Exp(10, 3)},
	{"M", "mega", r.Exp(10, 6)},
	{"G", "giga", r.Exp(10, 9)},
	{"T", "tera", r.Exp(10, 12)},
	{"P", "peta", r.Exp(10, 15)},
	{"E", "exa", r.Exp(10, 18)},
}

var digitalPrefixes = []prefix{
	{"K", "kilo", r.Exp(1000, 1)},
	{"M", "mega", r.Exp(1000, 2)},
	{"G", "giga", r.Exp(1000, 3)},
	{"T", "tera", r.Exp(1000, 4)},
	{"P", "peta", r.Exp(1000, 5)},
	{"E", "exa", r.Exp(1000, 6)},
	{"Z", "zetta", r.Exp(1000, 7)},
	{"Y", "yotta", r.Exp(1000, 8)},

	{"Ki", "kibi", r.Exp(1024, 1)},
	{"Mi", "mebi", r.Exp(1024, 2)},
	{"Gi", "gibi", r.Exp(1024, 3)},
	{"Ti", "tebi", r.Exp(1024, 4)},
	{"Pi", "pebi", r.Exp(1024, 5)},
	{"Ei", "exi", r.Exp(1024, 6)},
	{"Zi", "zebi", r.Exp(1024, 7)},
	{"Yi", "yobi", r.Exp(1024, 8)},
}
