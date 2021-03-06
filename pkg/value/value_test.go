package value

import (
	"testing"

	"github.com/nkanaev/numb/pkg/ratutils"
)

func TestFormat(t *testing.T) {
	testcases := []struct {
		have string
		prec int
		sep  string
		want string
	}{
		{"1", 2, ",", "1"},
		{"123", 2, ",", "123"},
		{"1234", 2, ",", "1,234"},
		{"1234567", 2, ",", "1,234,567"},
		{"1234567", 2, ",", "1,234,567"},
		{"-100", 2, ",", "-100"},
		{"-1000", 2, ",", "-1,000"},
		{"-100.1234", 2, ",", "-100.12"},
		{"-1000.1234", 2, ",", "-1,000.12"},
		{"1.01", 2, "", "1.01"},
		{"1.01", 4, "", "1.01"},
		{"1.1234", 2, "", "1.12"},
		{"1.001", 2, "", "1.0"},
		{"0.001", 2, "", "0.0"},
		{"0.001", 1, "", "0.0"},
		{"0.001", 0, "", "0"},
		{"100000.01", 2, ",", "100,000.01"},
		{"1234567890.01", 2, "_", "1_234_567_890.01"},
		{"1000000000.01", 2, "", "1000000000.01"},
	}
	for _, tc := range testcases {
		have := formatDec(ratutils.Must(tc.have), tc.sep, tc.prec)
		if have != tc.want {
			t.Errorf("val=%#v sep=%#v prec=%d\nwant: %s\nhave: %s", tc.have, tc.sep, tc.prec, tc.want, have)
		}
	}
}
