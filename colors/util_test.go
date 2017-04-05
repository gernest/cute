package colors

import (
	"testing"
)

func TestHex(t *testing.T) {
	sample := []struct {
		src     string
		r, g, b uint16
		rgb     string
	}{
		{"#9f3", 153, 255, 51, "rgb(153,255,51)"},
		{"#A94FD3", 169, 79, 211, "rgb(169,79,211)"},
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
		ss := v.String()
		if ss != s.rgb {
			t.Errorf("expected %s got %s", s.rgb, ss)
		}
	}

}

func TestValue_Luminance(t *testing.T) {

	sample := []struct {
		src string
		e   float64
	}{
		{"rgb(0, 0, 0)", 0},
		{"rgb(255, 255, 255)", 1},
		{"rgb(127, 127, 127)", 0.212},
		{"rgb(255, 127, 0)", 0.364},
	}

	for _, s := range sample {
		v, err := Parse(s.src)
		if err != nil {
			t.Fatal(err)
		}
		// fmt.Println(*v)
		g := v.Luminance()
		if g != s.e {
			t.Errorf("expected %.3f got %.3f", s.e, g)
		}
	}

}
