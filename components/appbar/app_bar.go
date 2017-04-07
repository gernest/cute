package appbar

import (
	"strconv"

	"github.com/gernest/cute/components/paper"
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

func newStyle(t themes.Manager) appBarStyle {
	base := t.AppBar()
	root := make(style.Object)
	root.Set("position", "relative")
	root.Set("z-index", style.AppBar.String())
	root.Set("width", "100%")
	root.Set("display", "flex")
	root.Set("background-color", base.Color.String())
	root.Set("padding-left", itos(base.Padding))
	root.Set("padding-right", itos(base.Padding))

	title := make(style.Object)
	title.Set("white-space", "nowrap")
	title.Set("overflow", "hidden")
	title.Set("text-overflow", "ellipsis")
	title.Set("margin", "0")
	title.Set("padding-top", "0")
	title.Set("letter-spacing", "0")
	title.Set("font-size", "24")
	title.Set("font-weight", itos(base.TitleFontWeight))
	title.Set("color", base.TextColor.String())
	title.Set("height", itos(base.Height))
	title.Set("line-height", itos(base.Height)+"px")
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
	Style               appBarStyle
	Title               *Title
	ClassName           string
	IconClassNameLeft   string
	IconClassNameRight  string
	IconElementLeft     vecty.Component
	IconElementRight    vecty.Component
	IconStyleLeft       style.Object
	IconStyleRight      style.Object
	OnRightIconTouchTap func(*vecty.Event)
	OnTitleTouchTap     func(*vecty.Event)
	ShowMenuIconButton  bool
	TitleStyle          style.Object
	theme               themes.Manager
}

type Title struct {
	Text      string
	Component vecty.Component
}

type Opts struct {
	Theme themes.Manager
	Title Title
}

func New(o Opts) *AppBar {
	a := &AppBar{
		Style: newStyle(o.Theme),
		Title: &o.Title,
		theme: o.Theme,
	}
	return a
}

func (a *AppBar) Render() *vecty.HTML {
	var title *vecty.HTML
	var titleStyle style.Object
	if a.TitleStyle != nil {
		titleStyle = a.TitleStyle
	} else {
		titleStyle = a.Style.title
	}
	if a.Title.Component != nil {
		title = elem.Div(
			titleStyle.Style(),
			a.Title.Component,
		)
	} else {
		title = elem.Heading1(
			titleStyle.Style(),
			vecty.Text(a.Title.Text),
		)
	}
	p := paper.New(a.theme, paper.Opts{Rounded: true})
	return elem.Div(
		p,
		a.Style.root.Style(),
		title,
	)
}
