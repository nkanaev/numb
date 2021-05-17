package unit

import (
	"math/big"
)

type prefix struct {
	abbr  string
	name  string
	value *big.Rat
}

func exp(b, n int64) *big.Rat {
	if n < 0 {
		int := big.NewInt(b)
		int.Exp(int, big.NewInt(-n), nil)
		rat := big.NewRat(1, 1)
		rat.Denom().Set(int)
		return rat
	}
	int := big.NewInt(b)
	int.Exp(int, big.NewInt(n), nil)
	rat := big.NewRat(1, 1)
	rat.Num().Set(int)
	return rat
}

var metricPrefixes = []prefix{
	{"d", "deci", exp(10, -1)},
	{"c", "centi", exp(10, -2)},
	{"m", "milli", exp(10, -3)},
	{"u", "micro", exp(10, -6)},
	{"n", "nano", exp(10, -9)},
	{"p", "pico", exp(10, -12)},
	{"f", "femto", exp(10, -15)},
	{"a", "atto", exp(10, -18)},
	{"z", "zepto", exp(10, -21)},
	{"y", "yocto", exp(10, -24)},

	{"da", "deca", exp(10, 1)},
	{"h", "hecto", exp(10, 2)},
	{"k", "kilo", exp(10, 3)},
	{"M", "mega", exp(10, 6)},
	{"G", "giga", exp(10, 9)},
	{"T", "tera", exp(10, 12)},
	{"P", "peta", exp(10, 15)},
	{"E", "exa", exp(10, 18)},
	{"Z", "zetta", exp(10, 21)},
	{"Y", "yotta", exp(10, 24)},
}

var digitalPrefixes = []prefix{
	{"K", "kilo", exp(1000, 1)},
	{"M", "mega", exp(1000, 2)},
	{"G", "giga", exp(1000, 3)},
	{"T", "tera", exp(1000, 4)},
	{"P", "peta", exp(1000, 5)},
	{"E", "exa", exp(1000, 6)},
	{"Z", "zetta", exp(1000, 7)},
	{"Y", "yotta", exp(1000, 8)},

	{"Ki", "kibi", exp(1024, 1)},
	{"Mi", "mebi", exp(1024, 2)},
	{"Gi", "gibi", exp(1024, 3)},
	{"Ti", "tebi", exp(1024, 4)},
	{"Pi", "pebi", exp(1024, 5)},
	{"Ei", "exi", exp(1024, 6)},
	{"Zi", "zebi", exp(1024, 7)},
	{"Yi", "yobi", exp(1024, 8)},
}
