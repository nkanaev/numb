package ratutils

import "testing"

func TestRatMod(t *testing.T) {
	tcases := []struct{
		a, n, want string
	}{
		{"10", "2.4", "0.4"},
		{"10", "2.2", "1.2"},
		{"10", "2", "0"},
		{"10", "3", "1"},
	}
	for _, tcase := range tcases {
		want := Num(tcase.want)
		have := ModRat(Num(tcase.a), Num(tcase.n))
		if want.Cmp(have) != 0 {
			t.Errorf("%s %% %s != %s (must be %s)", tcase.a, tcase.n, have, want)
		}
	}	
}
