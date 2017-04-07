package icon

import (
	"strconv"

	"github.com/gernest/cute/colors"
	"github.com/gernest/cute/style"
	"github.com/gernest/cute/themes"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

type Font struct {
	vecty.Core
	Color        *colors.Value
	HoverColor   colors.Value
	OnMouseEnter func(*vecty.Event)
	OnMouseLeave func(*vecty.Event)
	Style        style.Object
	hovered      bool
	root         style.Object
	offColor     colors.Value
}

func NewFont(t themes.Manager) *Font {
	s := make(style.Object)
	p := t.Spacings()
	s.Set("position", "relative")
	s.Set("font-size", p.Get(style.IconSize))
	s.Set("display", "inline-block")
	s.Set("user-select", "none")
	return &Font{root: s, offColor: t.TextColor()}
}
func itos(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func (f *Font) Render() *vecty.HTML {
	if f.hovered {
		if f.HoverColor.String() == "" {
			f.root.Set("color", f.offColor.String())
		} else {
			f.root.Set("color", f.HoverColor.String())
		}
	} else {
		f.root.Set("color", f.offColor.String())
	}
	var styles style.Object
	if f.Style != nil {
		styles = f.Style
	} else {
		styles = f.root
	}
	return elem.Span(
		styles,
		event.MouseLeave(f.handleMouseLeave),
		event.MouseEnter(f.handleMouseEnter),
	)
}

func (f *Font) handleMouseLeave(e *vecty.Event) {
	if f.HoverColor.String() == "" {
		f.hovered = false
	}
	if f.OnMouseLeave != nil {
		f.OnMouseLeave(e)
	}
}

func (f *Font) handleMouseEnter(e *vecty.Event) {
	if f.HoverColor.String() == "" {
		f.hovered = false
	}
	if f.OnMouseLeave != nil {
		f.OnMouseEnter(e)
	}
}
