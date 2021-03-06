package dimension

var Measures = map[string]Dimension{
	"dimensionless":                     {},
	"length":                            {Length: 1},
	"temperature":                       {Temperature: 1},
	"area":                              {Length: 2},
	"volume":                            {Length: 3},
	"mass":                              {Mass: 1},
	"time":                              {Time: 1},
	"digital":                           {Digital: 1},
	"currency":                          {Currency: 1},
	"frequency/radioactivity":           {Time: -1},
	"electric current":                  {Current: 1},
	"amount of substance":               {AmountOfSubstance: 1},
	"power":                             {Mass: 1, Length: 2, Time: -3},
	"force":                             {Mass: 1, Length: 1, Time: -2},
	"energy":                            {Mass: 1, Length: 2, Time: -2},
	"angular momentum":                  {Mass: 1, Length: 2, Time: -1},
	"moment of intertia":                {Length: 2, Mass: 1},
	"electric charge":                   {Time: 1, Current: 1},
	"electric potential":                {Mass: 1, Length: 2, Time: -3, Current: -1},
	"electric capacitance":              {Mass: -1, Length: -2, Time: 4, Current: 2},
	"electric conductance":              {Mass: -1, Length: -2, Time: 3, Current: 2},
	"magnetic flux":                     {Mass: 1, Length: 2, Time: -2, Current: -1},
	"magnetic flux density":             {Mass: 1, Time: -2, Current: -1},
	"electric inductance":               {Mass: 1, Length: 2, Time: -2, Current: -2},
	"electric resistance":               {Mass: 1, Length: 2, Time: -3, Current: 2},
	"pressure":                          {Mass: 1, Length: -1, Time: -2},
	"ionizing radiation/radiation dose": {Length: 2, Time: -2},
	"catalyctic activity":               {AmountOfSubstance: 1, Time: -1},
	"luminous flux/intensity":           {LuminousIntensity: 1},
	"illuminance":                       {LuminousIntensity: 1, Length: -1},
	"speed/velocity":                    {Length: 1, Time: -1},
	"data rate":                         {Digital: 1, Time: -1},
	"density":                           {Mass: 1, Length: -3},
	"flow":                              {Length: 3, Time: -1},
	"acceleration":                      {Length: 1, Time: -2},
	"momentum":                          {Length: 1, Mass: 1, Time: -1},
	"electric current density":          {Current: 1, Mass: -2},
	"price per mass":                    {Currency: 1, Mass: -1},
    /* TODO: more dimensions?
       surface tension?
       surface charge density
       electric charge density
       electric field strength
       magnetic field strength
       heat capacity
       specific heat capacity
       magnetic reluctance
       permeability
    */
}
