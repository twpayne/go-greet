package greet

import "testing"

func TestGreet(t *testing.T) {
	for _, tc := range []struct {
		name string
		want string
	}{
		{name: "twpayne", want: "Hello twpayne"},
	} {
		if got := Greet(tc.name); got != tc.want {
			t.Errorf("Greet(%q) == %q, want %q", tc.name, got, tc.want)
		}
	}
}
