package main

import (
	"github.com/gernest/cute/components/appbar"
	"github.com/gernest/cute/themes"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func main() {
	light := themes.NewLight()
	m := themes.NewManager(light)
	vecty.SetTitle("cute: beautiful material ui components for gopherjs")
	vecty.RenderBody(&app{theme: m})

}

func createAppBar(t themes.Manager) *appbar.AppBar {
	o := appbar.Opts{
		Theme: t,
		Title: appbar.Title{Text: "genesis"},
	}
	return appbar.New(o)
}

type app struct {
	vecty.Core
	theme themes.Manager
}

func (a *app) Render() *vecty.HTML {
	return elem.Body(
		createAppBar(a.theme),
	)
}
