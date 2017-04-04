package colors

import (
	"testing"
)

func TestHex(t *testing.T) {
	sample := []struct {
		src     string
		r, g, b uint16
	}{
		{"#9f3", 153, 255, 51},
		{"#A94FD3", 169, 79, 211},
	}
	for _, s := range sample {
		v, err := Hex(s.src)
		if err != nil {
			t.Fatal(err)
		}
		if v.R != s.r {
			t.Errorf("expected %d got %d", s.r, v.R)
		}
		if v.G != s.g {
			t.Errorf("expected %d got %d", s.g, v.G)
		}
		if v.B != s.b {
			t.Errorf("expected %d got %d", s.b, v.B)
		}
	}

}
