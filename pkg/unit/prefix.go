package unit

import (
	"math/big"
)

type prefix struct {
	abbr  string
	name  string
	value *big.Int
}

func exp(b, n int64) *big.Int {
	int := big.NewInt(b)
	return int.Exp(int, big.NewInt(n), nil)
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
	{"k", "kilo", exp(1000, 1)},
	{"m", "mega", exp(1000, 2)},
	{"g", "giga", exp(1000, 3)},
	{"t", "tera", exp(1000, 4)},
	{"p", "peta", exp(1000, 5)},
	{"e", "exa", exp(1000, 6)},
	{"z", "zetta", exp(1000, 7)},
	{"y", "yotta", exp(1000, 8)},

	{"ki", "kibi", exp(1024, 1)},
	{"mi", "mebi", exp(1024, 2)},
	{"gi", "gibi", exp(1024, 3)},
	{"ti", "tebi", exp(1024, 4)},
	{"pi", "pebi", exp(1024, 5)},
	{"ei", "exi", exp(1024, 6)},
	{"zi", "zebi", exp(1024, 7)},
	{"yi", "yobi", exp(1024, 8)},
}
