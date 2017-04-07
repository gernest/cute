package paper

import (
	"time"

	"github.com/gernest/cute/style"
	"github.com/gernest/cute/themes"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Paper struct {
	vecty.Core
	Children          vecty.Component
	Circle            bool
	Rounded           bool
	TransitionEnabled bool
	Zdepth            int
	root              style.Object
	Style             style.Object
}

type Opts struct {
	EnableTrasition bool
	Circle          bool
	Rounded         bool
}

func New(t themes.Manager, o Opts) *Paper {
	p := t.Paper()
	r := make(style.Object)
	s := t.Style()
	r.Set("color", p.Color.String())
	if o.EnableTrasition {
		r.Set("transition", style.EaseOut(450*time.Millisecond, 0*time.Millisecond, "all"))
	}
	r.Set("box-sizing", "border-box")
	r.Set("forn-family", s["font-family"])
	r.Set("webkit-tap-highlight-color", "rgba(0,0,0,0)")

	//TODO missing
	// set box-shadow
	var rad string
	if o.Circle {
		rad = "50%"
	} else {
		if o.Rounded {
			rad = s["border-radius"]
		} else {
			rad = "0px"
		}
	}
	r.Set("border-radius", rad)
	return &Paper{
		root:              r,
		Circle:            o.Circle,
		TransitionEnabled: o.EnableTrasition,
		Rounded:           o.Rounded,
	}
}

func (p *Paper) Render() *vecty.HTML {
	if p.Style != nil {
		p.root.Merge(p.Style)
	}
	return elem.Div(
		p.root.Style(),
		p.Children,
	)
}
