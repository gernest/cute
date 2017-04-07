package themes

import (
	"github.com/gernest/cute/colors"
	"github.com/gernest/cute/style"
)

type Light struct {
	baseTheme
}

func NewLight() Light {
	o := make(style.Object)
	o.Set("font-family", "Roboto, sans-serif")
	o.Set("border-radius", "2")
	s := style.DefaultSpacing()
	b := baseTheme{
		primary1Color:      colors.New(colors.Cyan500),
		primary2Color:      colors.New(colors.Cyan700),
		primary3Color:      colors.New(colors.Grey400),
		accent1Color:       colors.New(colors.PinkA200),
		accent2Color:       colors.New(colors.Grey100),
		accent3Color:       colors.New(colors.Grey500),
		textColor:          colors.New(colors.DarkBlack),
		secondaryTextColor: colors.New(colors.DarkBlack).Fade(0.54),
		alternateTextColor: colors.New(colors.White),
		canvasColor:        colors.New(colors.White),
		borderColor:        colors.New(colors.Grey300),
		disabledColor:      colors.New(colors.DarkBlack).Fade(0.3),
		pickerHeaderColor:  colors.New(colors.Cyan500),
		clockCircleColor:   colors.New(colors.DarkBlack).Fade(0.07),
		shadowColor:        colors.New(colors.FullBlack),
		spacings:           s,
		style:              o,
	}
	return Light{baseTheme: b}
}
