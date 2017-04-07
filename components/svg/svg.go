package svg

import (
	"github.com/gernest/cute/colors"
	"github.com/gernest/cute/style"
	"github.com/gopherjs/vecty"
)

type Icon struct {
	vecty.Core
	Children     vecty.Component
	Color        colors.Value
	Style        style.Object
	ViewBox      string
	OnMouseEnter func(*vecty.Event)
	OnMouseLeave func(*vecty.Event)
	hovered      bool
}

func (i *Icon) handleMouseLeave(e *vecty.Event) {
	i.hovered = false
	if i.OnMouseLeave != nil {
		i.OnMouseLeave(e)
	}
}
func (i *Icon) handleMouseEnter(e *vecty.Event) {
	i.hovered = true
	if i.OnMouseLeave != nil {
		i.OnMouseLeave(e)
	}
}
