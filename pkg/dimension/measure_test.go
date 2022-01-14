package dimension

import "testing"

func TestMeasuresUnique(t *testing.T) {
	track := make(map[Dimensions]string)
	for name, dim := range Measures {
		if oldname, seen := track[dim]; seen {
			t.Errorf("measure `%s` has the same dimension as `%s`", name, oldname)
		} else {
			track[dim] = name
		}
	}
}
