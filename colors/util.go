package colors

import (
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
	v := &Value{Type: "rgb"}
	parts := strings.Split(src, "")
	if len(parts) == 3 {
		rs := parts[0]
		rs += rs
		r, err := strconv.ParseInt(rs, 16, 32)
		if err != nil {
			return nil, err
		}
		v.R = uint16(r)
		gs := parts[1]
		gs += gs
		g, err := strconv.ParseInt(gs, 16, 32)
		if err != nil {
			return nil, err
		}
		v.G = uint16(g)
		bs := parts[2]
		bs += bs
		b, err := strconv.ParseInt(bs, 16, 32)
		if err != nil {
			return nil, err
		}
		v.B = uint16(b)
	}
	return v, nil
}
