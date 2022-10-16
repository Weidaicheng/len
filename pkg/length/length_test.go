package length

import "testing"

func TestGetLen(t *testing.T) {
	str := "abc"
	want := 3
	len := GetLen(str)
	if want != len {
		t.Fatalf(`GetLen("abc") = %v, want equal to %v`, len, want)
	}
}
