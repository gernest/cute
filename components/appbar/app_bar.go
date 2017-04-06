package appbar

import (
	"strconv"

	"github.com/gernest/cute/style"
	"github.com/gernest/cute/themes"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type appBarStyle struct {
	root                style.Object
	title               style.Object
	mainElement         style.Object
	iconButtonStyle     style.Object
	iconButtonIconStyle style.Object
	flatButton          style.Object
}

func newStyle(base themes.AppBar) appBarStyle {
	root := make(style.Object)
	root.Set("position", "relative")
	root.Set("z-ndex", style.AppBar.String())
	root.Set("width", "100")
	root.Set("display", "flex")
	root.Set("background-color", base.Color.String())
	root.Set("padding-left", itos(base.Padding))
	root.Set("padding-right", itos(base.Padding))

	title := make(style.Object)
	title.Set("whitespace", "nowrap")
	title.Set("overflow", "hidden")
	return appBarStyle{
		root:  root,
		title: title,
	}
}

func itos(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

type AppBar struct {
	vecty.Core
	style appBarStyle
	title *Title
}

type Title struct {
	Text      string
	Component vecty.Component
}

type Opts struct {
	Theme themes.AppBar
	Title Title
}

func New(o Opts) *AppBar {
	a := &AppBar{
		style: newStyle(o.Theme),
		title: &o.Title,
	}
	return a
}

func (a *AppBar) Render() *vecty.HTML {
	var title *vecty.HTML
	if a.title.Component != nil {
		title = elem.Div(
			a.style.title.Style(),
			a.title.Component,
		)
	} else {
		title = elem.Heading1(
			a.style.title.Style(),
			vecty.Text(a.title.Text),
		)
	}
	return elem.Div(
		a.style.root.Style(),
		title,
	)
}