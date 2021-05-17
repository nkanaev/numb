package unit

import (
	"github.com/nkanaev/numb/pkg/consts"
)

var angleUnits = []baseUnit{
	{d: ANGLE, name: "rad", long: "radian", value: f64(1), info: "SI derived unit"},
	{d: ANGLE, name: "°", long: "deg, degree", value: unitdiv(consts.PI, 180), info: "SI-accepted unit"},
	{d: ANGLE, name: "arcsec", value: unitdiv(consts.PI, 648000), info: "SI-accepted unit (pi / 648000)"},
	{d: ANGLE, name: "arcmin", value: unitdiv(consts.PI, 10800), info: "SI-accepted unit (pi / 10800)"},
	{d: ANGLE, name: "grad", long: "grade, gradian", value: unitdiv(consts.PI, 200)},
	{d: ANGLE, name: "cycle", value: unitdiv(consts.PI, 2)},
}
