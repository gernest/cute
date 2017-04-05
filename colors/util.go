package colors

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/cathalgarvey/fmtless"
)

type Value struct {
	Type string
	R    uint16
	G    uint16
	B    uint16
}

//Parse parses the color string to a maningful structure
func Parse(src string) (*Value, error) {
	if src == "" {
		return nil, errors.New("empty color")
	}
	if src[0] == '#' {
		return Hex(src)
	}
	i := strings.Index(src, "(")
	if i == -1 {
		return nil, errors.New("bad color scheme")
	}
	typ := src[:i]
	switch typ {
	case "rgb":
		return RGB(src)
	}
	return nil, errors.New("not supported")
}

func Hex(src string) (*Value, error) {
	if src == "" {
		return nil, errors.New("empty color")
	}
	if src[0] != '#' {
		return nil, errors.New("invalid hex value")
	}
	src = src[1:]
	v := &Value{Type: "rgb"}
	parts := strings.Split(src, "")
	var rs, gs, bs string
	switch len(parts) {
	case 3:
		rs = parts[0]
		rs += rs
		gs = parts[1]
		gs += gs
		bs = parts[2]
		bs += bs
	case 6:
		rs = strings.Join(parts[0:2], "")
		gs = strings.Join(parts[2:4], "")
		bs = strings.Join(parts[4:6], "")
	default:
		return nil, errors.New("invalid hex value")

	}
	r, err := strconv.ParseInt(rs, 16, 32)
	if err != nil {
		return nil, err
	}
	v.R = uint16(r)
	g, err := strconv.ParseInt(gs, 16, 32)
	if err != nil {
		return nil, err
	}
	v.G = uint16(g)
	b, err := strconv.ParseInt(bs, 16, 32)
	if err != nil {
		return nil, err
	}
	v.B = uint16(b)
	return v, nil
}

func RGB(src string) (*Value, error) {
	i := strings.Index(src, "(")
	if i == -1 {
		return nil, errors.New("bad color scheme" + src)
	}
	typ := src[:i]
	if typ != "rgb" {
		return nil, errors.New("bad color scheme")
	}
	p, err := partsUint(parts(src[i:]))
	if err != nil {

		return nil, err
	}
	if len(p) != 3 {
		return nil, errors.New("bad color scheme")
	}
	return &Value{Type: "rgb", R: p[0], G: p[1], B: p[1]}, nil

}
func parts(src string) []string {
	src = strings.TrimPrefix(src, "(")
	src = strings.TrimSuffix(src, ")")
	return strings.Split(src, ",")
}
func partsUint(src []string) ([]uint16, error) {
	var v []uint16
	for _, val := range src {
		u, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			return nil, err
		}
		v = append(v, uint16(u))
	}
	return v, nil
}

func (v *Value) String() string {
	return fmt.Sprintf("%s(%d,%d,%d)", v.Type, int(v.R), int(v.G), int(v.B))
}

const factor = 1.0 / 255.0

func normalize(v uint16) float64 {
	val := float64(v) * factor
	if val <= 0.03928 {
		return val / 12.92
	}
	return math.Pow(val+0.055/1.055, 2.4)
}

func (v *Value) normalized() (float64, float64, float64) {
	return normalize(v.R),
		normalize(v.G),
		normalize(v.B)
}

func (v *Value) lum(r, g, b float64) float64 {
	return 0.2126*r + 0.7152*g + 0.0722*b
}

func (v *Value) Luminance() float64 {
	return v.lum(v.normalized())
}
