package themes

import (
	"github.com/gernest/cute/colors"
	"github.com/gernest/cute/style"
	"github.com/gernest/cute/typography"
)

var (
	_ Theme = baseTheme{}
)

//Theme defines material theme
type Theme interface {
	Primary1Color() colors.Value
	Primary2Color() colors.Value
	Primary3Color() colors.Value
	Accent1Color() colors.Value
	Accent2Color() colors.Value
	Accent3Color() colors.Value
	TextColor() colors.Value
	SecondaryTextColor() colors.Value
	AlternateTextColor() colors.Value
	CanvasColor() colors.Value
	DisableColor() colors.Value
	PickerHeaderColor() colors.Value
	ClockCircleColor() colors.Value
	ShadowColor() colors.Value
	Style() style.Object
	Spacings() style.SpaceMap
}

type baseTheme struct {
	primary1Color      colors.Value
	primary2Color      colors.Value
	primary3Color      colors.Value
	accent1Color       colors.Value
	accent2Color       colors.Value
	accent3Color       colors.Value
	textColor          colors.Value
	secondaryTextColor colors.Value
	alternateTextColor colors.Value
	canvasColor        colors.Value
	borderColor        colors.Value
	disabledColor      colors.Value
	pickerHeaderColor  colors.Value
	clockCircleColor   colors.Value
	shadowColor        colors.Value
	style              style.Object
	spacings           style.SpaceMap
}

func (b baseTheme) Primary1Color() colors.Value {
	return b.primary1Color
}

func (b baseTheme) Primary2Color() colors.Value {
	return b.primary2Color
}

func (b baseTheme) Primary3Color() colors.Value {
	return b.primary3Color
}

func (b baseTheme) Accent1Color() colors.Value {
	return b.accent1Color
}

func (b baseTheme) Accent2Color() colors.Value {
	return b.accent2Color
}

func (b baseTheme) Accent3Color() colors.Value {
	return b.accent2Color
}

func (b baseTheme) TextColor() colors.Value {
	return b.textColor
}

func (b baseTheme) SecondaryTextColor() colors.Value {
	return b.secondaryTextColor
}

func (b baseTheme) AlternateTextColor() colors.Value {
	return b.alternateTextColor
}

func (b baseTheme) CanvasColor() colors.Value {
	return b.canvasColor
}

func (b baseTheme) DisableColor() colors.Value {
	return b.disabledColor
}

func (b baseTheme) PickerHeaderColor() colors.Value {
	return b.pickerHeaderColor
}

func (b baseTheme) ClockCircleColor() colors.Value {
	return b.clockCircleColor
}
func (b baseTheme) ShadowColor() colors.Value {
	return b.shadowColor
}

func (b baseTheme) Style() style.Object {
	return b.style
}
func (b baseTheme) Spacings() style.SpaceMap {
	return b.spacings
}

// Manager is a theme Manager
type Manager struct {
	baseTheme
	appBar               AppBar
	avatar               Avatar
	badge                Badge
	bottomNav            BottomNav
	button               Button
	card                 Card
	cardMedia            CardMedia
	cardText             CardText
	checkBox             CheckBox
	chip                 Chip
	datePicker           DatePicker
	dialog               Dialog
	enhancedButtom       EnhancedButton
	flatButton           FlatButton
	floatingActionButton FloatingActionButton
	gridTitle            GridTitle
	icon                 Icon
	inkBar               InkBar
	drawer               Drawer
	listItem             ListItem
	menu                 Menu
	menuItem             MenuItem
	menuSubHeader        MenuSubHeader
	overlay              Overlay
}

// NewManager returns a new theme manager
func NewManager(theme Theme) Manager {
	s := theme.Spacings()
	m := Manager{
		appBar: AppBar{
			Color:           theme.Primary1Color(),
			TextColor:       theme.AlternateTextColor(),
			Height:          s[style.DesktopKeyIncrement],
			TitleFontWeight: typography.FontWeightNormal,
			Padding:         s[style.DesktopGutter],
		},
		avatar: Avatar{
			Color:           theme.CanvasColor(),
			BackgroundColor: theme.CanvasColor().Emphasize(0.26),
		},
		badge: Badge{
			Color:              theme.AlternateTextColor(),
			TextColor:          theme.TextColor(),
			PrimaryColor:       theme.Primary1Color(),
			PrimaryTextColor:   theme.AlternateTextColor(),
			SecondaryColor:     theme.Accent1Color(),
			SecondaryTextColor: theme.AlternateTextColor(),
			FontWeight:         typography.FontWeightMedium,
		},
	}

	return m
}

// AppBar AppBar style
type AppBar struct {
	Color           colors.Value
	TextColor       colors.Value
	Height          int
	TitleFontWeight int
	Padding         int
}

// Avatar Avatar style
type Avatar struct {
	Color           colors.Value
	BackgroundColor colors.Value
}

// Badge badge style
type Badge struct {
	Color              colors.Value
	TextColor          colors.Value
	PrimaryColor       colors.Value
	PrimaryTextColor   colors.Value
	SecondaryColor     colors.Value
	SecondaryTextColor colors.Value
	FontWeight         int
}

// BottomNav Bottom Navigation style
type BottomNav struct {
	BackgroundColor    colors.Value
	UnselectedColor    colors.Value
	SelectedColor      colors.Value
	Height             int
	UnselectedFontSize int
	SelectedFontSize   int
}

// Button Button style
type Button struct {
	Height         int
	MinWidth       int
	IconButtonSize int
}

// Card Button style
type Card struct {
	TitleColor    colors.Value
	SubtitleColor colors.Value
	FontWeight    int
}

// CardMedia Button style
type CardMedia struct {
	Color         colors.Value
	TitleColor    colors.Value
	SubtitleColor colors.Value
}

// CheckBox Button style
type CheckBox struct {
	BoxColor           colors.Value
	CheckedColor       colors.Value
	RequiredColor      colors.Value
	DisabledColor      colors.Value
	LabelColor         colors.Value
	LabelDisabledColor colors.Value
}

// CardText Button style
type CardText struct {
	TextColor colors.Value
}

// Chip Button style
type Chip struct {
	BackgroundColor colors.Value
	DeleteIconColor colors.Value
	TextColor       colors.Value
	FontSize        int
	FontWeight      int
	Shadow          string
}

// DatePicker Button style
type DatePicker struct {
	Color                       colors.Value
	TextColor                   colors.Value
	CalendarTextColor           colors.Value
	SelectColor                 colors.Value
	SelectTextColor             colors.Value
	CalendarYearBackgroundColor colors.Value
}

// Dialog Button style
type Dialog struct {
	TitleFontSize int
	BodyFontSize  int
	BodyColor     colors.Value
}

// FlatButton Button style
type FlatButton struct {
	Color              colors.Value
	ButtonFilterColor  colors.Value
	DisabledTextColor  colors.Value
	TextColor          colors.Value
	PrimaryTextColor   colors.Value
	SecondaryTextColor colors.Value
	FontSize           int
	FontWeight         int
}

// DropDownMenu Button style
type DropDownMenu struct {
	AccentColor colors.Value
}

// EnhancedButton Button style
type EnhancedButton struct {
	TapHighlightColor colors.Value
}

// FloatingActionButton Button style
type FloatingActionButton struct {
	ButtonSize         int
	MiniSize           int
	Color              colors.Value
	IconColor          colors.Value
	SecondaryColor     colors.Value
	SecondaryIconColor colors.Value
	DisabledTextColor  colors.Value
	DisabledColor      colors.Value
}

// GridTitle Button style
type GridTitle struct {
	TextColor colors.Value
}

// Icon Button style
type Icon struct {
	Color           colors.Value
	BackgroundColor colors.Value
}

// InkBar Button style
type InkBar struct {
	BackgroundColor colors.Value
}

// Drawer Button style
type Drawer struct {
	Color colors.Value
	Width style.Space
}

// ListItem Button style
type ListItem struct {
	NestedLevelDepth   int
	SecondaryTextColor colors.Value
	LeftIconColor      colors.Value
	RightIconColor     colors.Value
}

// Menu Button style
type Menu struct {
	BackgroundColor          colors.Value
	ContainerBackgroundColor colors.Value
}

// MenuItem Button style
type MenuItem struct {
	DataHeight           style.Space
	Height               style.Space
	HoverColor           colors.Value
	Padding              style.Space
	SelectedTextColor    colors.Value
	RightIconDesktopFill colors.Value
}

// MenuSubHeader Button style
type MenuSubHeader struct {
	Padding     style.Space
	BorderColor colors.Value
	TextColor   colors.Value
}

// Overlay Button style
type Overlay struct {
	BackgroundColor colors.Value
}

// AppBar returns the AppBar
func (m Manager) AppBar() AppBar {
	return m.appBar
}

// Avatar return Avatar
func (m Manager) Avatar() Avatar {
	return m.avatar
}

// Badge returns Badge
func (m Manager) Badge() Badge {
	return m.badge
}

// BottomNav returns BottomNav
func (m Manager) BottomNav() BottomNav {
	return m.bottomNav
}

// Button returns Button
func (m Manager) Button() Button {
	return m.button
}

// Card returns Card
func (m Manager) Card() Card {
	return m.card
}

// CardMedia returns CardMedia
func (m Manager) CardMedia() CardMedia {
	return m.cardMedia
}

// CardText returns CardText
func (m Manager) CardText() CardText {
	return m.cardText
}

// CheckBox returns CheckBox
func (m Manager) CheckBox() CheckBox {
	return m.checkBox
}

// Chip returns Chip
func (m Manager) Chip() Chip {
	return m.chip
}

// DatePicker returns DatePicker
func (m Manager) DatePicker() DatePicker {
	return m.datePicker
}

// Dialog returns Dialog
func (m Manager) Dialog() Dialog {
	return m.dialog
}

// EnhancedButton returns EnhancedButton
func (m Manager) EnhancedButton() EnhancedButton {
	return m.enhancedButtom
}

// FlatButton returns FlatButton
func (m Manager) FlatButton() FlatButton {
	return m.flatButton
}

// FloatingActionButton returns FloatingActionButton
func (m Manager) FloatingActionButton() FloatingActionButton {
	return m.floatingActionButton
}

// Icon returns Icon
func (m Manager) Icon() Icon {
	return m.icon
}

// InkBar returns InkBar
func (m Manager) InkBar() InkBar {
	return m.inkBar
}

// Drawer returns Drawer
func (m Manager) Drawer() Drawer {
	return m.drawer
}

// ListItem returns ListItem
func (m Manager) ListItem() ListItem {
	return m.listItem
}

// Menu returns Menu
func (m Manager) Menu() Menu {
	return m.menu
}

// MenuItem returns Menu
func (m Manager) MenuItem() MenuItem {
	return m.menuItem
}

// MenuSubHeader returns MenuSubHeader
func (m Manager) MenuSubHeader() MenuSubHeader {
	return m.menuSubHeader
}

// Overlay returns Overlay
func (m Manager) Overlay() Overlay {
	return m.overlay
}
