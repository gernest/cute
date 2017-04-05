package colors

import (
	"testing"
)

func TestHex(t *testing.T) {
	sample := []struct {
		src     string
		r, g, b uint8
		rgb     string
	}{
		{"#9f3", 153, 255, 51, "rgb(153,255,51)"},
		{"#A94FD3", 169, 79, 211, "rgb(169,79,211)"},
	}
	for _, s := range sample {
		v, err := Parse(s.src)
		if err != nil {
			t.Fatal(err)
		}
		r, g, b := v.RGB255()
		if r != s.r {
			t.Errorf("expected %d got %d", s.r, r)
		}
		if g != s.g {
			t.Errorf("expected %d got %d", s.g, g)
		}
		if b != s.b {
			t.Errorf("expected %d got %d", s.b, b)
		}
		ss := v.String()
		if ss != s.rgb {
			t.Errorf("expected %s got %s", s.rgb, ss)
		}
	}

}
