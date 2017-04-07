package colors

import (
	"errors"
	"math"
	"strconv"

	"fmt"

	"strings"

	colorful "github.com/lucasb-eyer/go-colorful"
)

type Value struct {
	Type string
	colorful.Color
	A float64
}

func New(c Color) Value {
	v, err := Parse(string(c))
	if err != nil {
		panic(err)
	}
	return v
}

//Parse parses the color string to a maningful structure
func Parse(src string) (Value, error) {
	if src == "" {
		return Value{}, errors.New("empty color")
	}
	if src[0] != '#' {
		i := strings.Index(src, "(")
		if i == -1 {
			return Value{}, errors.New("bad value")
		}
		typ := src[:i]
		switch typ {
		case "rgb":
			p := parts(src[i:])
			return Value{Type: "rgb", Color: colorful.Color{R: p[0], G: p[1], B: p[2]}}, nil
		case "rgba":
			p := parts(src[i:])
			return Value{Type: "rgb", Color: colorful.Color{R: p[0], G: p[1], B: p[2]}, A: p[3]}, nil
		}
		return Value{}, errors.New("unsupported type")

	}
	c, err := colorful.Hex(src)
	if err != nil {
		return Value{}, err
	}
	return Value{Type: "rgb", Color: c}, nil
}

func parts(src string) []float64 {
	src = strings.TrimPrefix(src, "(")
	src = strings.TrimSuffix(src, ")")
	p := strings.Split(src, ",")
	e := make([]float64, len(p))
	for i, v := range p {
		f, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
		if err != nil {
			panic(err)
		}
		e[i] = f
	}
	return e
}

func (v *Value) Luminance() float64 {
	return roundPlus(0.299*v.R+0.567*v.G+0.114*v.B, 3)
}

func clamp(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func contrastRatio(fg, bg *Value) float64 {
	a := fg.Luminance()
	b := bg.Luminance()
	c := (math.Max(a, b) + 0.05) / (math.Max(a, b) + 0.05)
	return roundPlus(c, 2)
}

func roundPlus(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return round(f*shift) / shift
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}

func (v Value) Darken(coe float64) Value {
	c := clamp(coe, 0, 1)
	return Value{
		Type: v.Type,
		Color: colorful.Color{
			R: v.R * (1 - c),
			G: v.G * (1 - c),
			B: v.B * (1 - c),
		},
	}
}
func (v Value) Lighten(coe float64) Value {
	c := clamp(coe, 0, 1)
	return Value{
		Type: v.Type,
		Color: colorful.Color{
			R: (float64(255) - v.R) * c,
			G: (float64(255) - v.G) * c,
			B: (float64(255) - v.B) * c,
		},
	}
}

func (v Value) String() string {
	r, g, b := v.RGB255()
	if v.Type == "rgba" {
		return fmt.Sprintf("%s(%d,%d,%d,%.3f)", v.Type, r, g, b, v.A)
	}
	return fmt.Sprintf("%s(%d,%d,%d)", v.Type, r, g, b)
}

func (v Value) Emphasize(val float64) Value {
	l := v.Luminance()
	if l > 0.5 {
		return v.Darken(val)
	}
	return v.Lighten(val)
}

func (v Value) Fade(val float64) Value {
	n := v
	c := clamp(val, 0, 1)
	n.Type += "a"
	n.A = c
	return n
}
