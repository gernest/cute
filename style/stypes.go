package style

import "strconv"
import "github.com/gopherjs/vecty"
import "time"

// ZIndex css zindex value
type ZIndex int

//common z-index
const (
	Menu          ZIndex = 1000
	AppBar        ZIndex = 1100
	DrawerOverlay ZIndex = 1400
	Drawer        ZIndex = 1300
	DialogOverlay ZIndex = 1400
	Dialog        ZIndex = 1400
	Layer         ZIndex = 2000
	PopOver       ZIndex = 2100
	SnackBar      ZIndex = 2900
	ToolTip       ZIndex = 3000
)

func (z ZIndex) String() string {
	return strconv.FormatInt(int64(z), 10)
}

// Space css spacing
type Space int

//common spacings
const (
	IconSize Space = iota
	DesktopGutter
	DesktopGutterMore
	DesktopGutterLess
	DesktopGutterMini
	DesktopKeyIncrement
	DesktopDropDownItemHeight
	DesktopDropDownMenuFontSize
	DesktopDrawerMenuItemHeight
	DesktopSubHeaderHeight
	DesktopToolBarheight
)

func (s Space) String() string {
	switch s {
	case IconSize:
		return "iconSize"
	case DesktopGutter:
		return "desktopGutter"
	case DesktopGutterMore:
		return "desktopGutterMore"
	case DesktopGutterLess:
		return "desktopGutterLess"
	case DesktopGutterMini:
		return "desktopGutterMini"
	case DesktopKeyIncrement:
		return "desktopKeyIncrement"
	case DesktopDropDownItemHeight:
		return "desktopDropDownItemHeight"
	case DesktopDropDownMenuFontSize:
		return "desktopDropDownMenuFontSize"
	case DesktopDrawerMenuItemHeight:
		return "desktopDrawerMenuItemHeight"
	case DesktopSubHeaderHeight:
		return "desktopSubHeaderHeight"
	case DesktopToolBarheight:
		return "desktopToolBarheight"
	}
	return ""
}

// Object defines a group of key value pairs of css styles
type Object map[string]string

func (o Object) Set(k, v string) {
	o[k] = v
}
func (o Object) Merge(n Object) {
	for k, v := range n {
		o.Set(k, v)
	}
}

type markList []vecty.Markup

func (m markList) Apply(h *vecty.HTML) {
	for _, f := range m {
		f.Apply(h)
	}
}
func (o Object) Style() vecty.Markup {
	var m []vecty.Markup
	for k, v := range o {
		m = append(m, vecty.Style(k, v))
	}
	return markList(m)
}

// SpaceMap stores the spacing settings. Keys should be the the Space constants
// defined above
type SpaceMap map[Space]int

// Merge combines m with s
func (s SpaceMap) Merge(m SpaceMap) {
	for k, v := range m {
		s[k] = v
	}
}
func (s SpaceMap) Get(k Space) string {
	return strconv.FormatInt(int64(s[k]), 10)

}

// DefaultSpacing returns a map of default space settings
func DefaultSpacing() SpaceMap {
	m := make(SpaceMap)
	m[IconSize] = 24
	m[DesktopGutter] = 24
	m[DesktopGutterLess] = 18
	m[DesktopGutterMini] = 8
	m[DesktopKeyIncrement] = 64
	m[DesktopDropDownItemHeight] = 32
	m[DesktopDropDownMenuFontSize] = 15
	m[DesktopDrawerMenuItemHeight] = 48
	m[DesktopSubHeaderHeight] = 48
	m[DesktopToolBarheight] = 56
	return m
}

//common css properties
const (
	CSSColor = "color"
)

const (
	EaseOutFunction   = "cubic-bezier(0.23, 1, 0.32, 1)"
	EaseInOutFunction = "cubic-bezier(0.445, 0.05, 0.55, 0.95)"
)

func Create(duration, delay time.Duration, property, easeFunc string) string {
	return property + " " + mili(duration) + " " + easeFunc + " " + mili(delay)
}

func mili(d time.Duration) string {
	n := int64(d / time.Millisecond)
	return strconv.FormatInt(n, 10) + "ms"
}

func EaseOut(duration, delay time.Duration, property string) string {
	return Create(duration, delay, property, EaseOutFunction)
}
