package unit

import (
	"github.com/nkanaev/numb/pkg/consts"
)

var angleUnits = []baseUnit{
	{name: "rad", long: "radian", value: f64(1), dimension: ANGLE, info: "SI derived unit"},
	{name: "°", long: "deg, degree", value: unitdiv(consts.PI, 180), dimension: ANGLE, info: "SI-accepted unit"},
	{name: "arcsec", value: unitdiv(consts.PI, 648000), dimension: ANGLE, info: "SI-accepted unit (pi / 648000)"},
	{name: "arcmin", value: unitdiv(consts.PI, 10800), dimension: ANGLE, info: "SI-accepted unit (pi / 10800)"},
	{name: "grad", long: "grade, gradian", value: unitdiv(consts.PI, 200), dimension: ANGLE},
	{name: "cycle", value: unitdiv(consts.PI, 2), dimension: ANGLE},
}
