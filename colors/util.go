package colors

import (
	"errors"
	"strconv"
	"strings"
)

type Value struct {
	Type string
	R    uint16
	G    uint16
	B    uint16
}

//Parse parses the color string to a maningful structure
func Parse(src string) (*Value, error) {
	return nil, nil
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

func (v *Value) HexString() string {
	r := strconv.FormatUint(uint64(v.R), 16)
	g := strconv.FormatUint(uint64(v.G), 16)
	b := strconv.FormatUint(uint64(v.B), 16)
	return "#" + r + g + b
}
