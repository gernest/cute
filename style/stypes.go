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
	IconSize                    Space = 24
	DesktopGutter               Space = 24
	DesktopGutterMore           Space = 32
	DesktopGutterLess           Space = 16
	DesktopGutterMini           Space = 8
	DesktopKeyIncrement         Space = 64
	DesktopDropDownItemHeight   Space = 32
	DesktopDropDownMenuFontSize Space = 15
	DesktopDrawerMenuItemHeight Space = 48
	DesktopSubHeaderHeight      Space = 48
	DesktopToolBarheight        Space = 56
)
