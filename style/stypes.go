package style

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

// SpaceMap stores the spacing settings. Keys should be the the Space constants
// defined above
type SpaceMap map[Space]int

// Merge combines m with s
func (s SpaceMap) Merge(m SpaceMap) {
	for k, v := range m {
		s[k] = v
	}
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