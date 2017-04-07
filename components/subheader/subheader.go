package subheader

import (
	"strconv"

	"github.com/gernest/cute/style"
	"github.com/gernest/cute/themes"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type SubHeader struct {
	vecty.Core
	Children vecty.Component
	Inset    bool
	Style    style.Object
	root     style.Object
}

func New(s themes.Manager, inset bool) *SubHeader {
	sub := s.SubHeader()
	root := make(style.Object)
	root.Set("box-sizing", "border-box")
	root.Set("color", sub.Color.String())
	root.Set("font-size", itos(14))
	root.Set("font-weight", itos(sub.FontWeight))
	root.Set("line-height", "48px")
	if inset {
		root.Set("padding-left", itos(72))
	} else {
		root.Set("padding-left", itos(16))
	}
	root.Set("width", "100%")
	return &SubHeader{
		Inset: inset,
		root:  root,
	}
}

func itos(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func (s *SubHeader) Render() *vecty.HTML {
	if s.Style != nil {
		s.root.Merge(s.Style)
	}
	return elem.Div(
		s.root.Style(),
		s.Children,
	)
}
