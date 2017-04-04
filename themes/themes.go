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
	Primary1Color() colors.Color
	Primary2Color() colors.Color
	Primary3Color() colors.Color
	Accent1Color() colors.Color
	Accent2Color() colors.Color
	Accent3Color() colors.Color
	TextColor() colors.Color
	SecondaryTextColor() colors.Color
	AlternateTextColor() colors.Color
	CanvasColor() colors.Color
	DisableColor() colors.Color
	PickerHeaderColor() colors.Color
	ClockCircleColor() colors.Color
	ShadowColor() colors.Color
	Style() style.Object
	Spacings() style.SpaceMap
}

type baseTheme struct {
	primary1Color      colors.Color
	primary2Color      colors.Color
	primary3Color      colors.Color
	accent1Color       colors.Color
	accent2Color       colors.Color
	accent3Color       colors.Color
	textColor          colors.Color
	secondaryTextColor colors.Color
	alternateTextColor colors.Color
	canvasColor        colors.Color
	borderColor        colors.Color
	disabledColor      colors.Color
	pickerHeaderColor  colors.Color
	clockCircleColor   colors.Color
	shadowColor        colors.Color
	style              style.Object
	spacings           style.SpaceMap
}

func (b baseTheme) Primary1Color() colors.Color {
	return b.primary1Color
}

func (b baseTheme) Primary2Color() colors.Color {
	return b.primary2Color
}

func (b baseTheme) Primary3Color() colors.Color {
	return b.primary3Color
}

func (b baseTheme) Accent1Color() colors.Color {
	return b.accent1Color
}

func (b baseTheme) Accent2Color() colors.Color {
	return b.accent2Color
}

func (b baseTheme) Accent3Color() colors.Color {
	return b.accent2Color
}

func (b baseTheme) TextColor() colors.Color {
	return b.textColor
}

func (b baseTheme) SecondaryTextColor() colors.Color {
	return b.secondaryTextColor
}

func (b baseTheme) AlternateTextColor() colors.Color {
	return b.alternateTextColor
}

func (b baseTheme) CanvasColor() colors.Color {
	return b.canvasColor
}

func (b baseTheme) DisableColor() colors.Color {
	return b.disabledColor
}

func (b baseTheme) PickerHeaderColor() colors.Color {
	return b.pickerHeaderColor
}

func (b baseTheme) ClockCircleColor() colors.Color {
	return b.clockCircleColor
}
func (b baseTheme) ShadowColor() colors.Color {
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
			Color: theme.CanvasColor(),
		},
	}

	return m
}

// AppBar AppBar style
type AppBar struct {
	Color           colors.Color
	TextColor       colors.Color
	Height          int
	TitleFontWeight int
	Padding         int
}

// Avatar Avatar style
type Avatar struct {
	Color           colors.Color
	BackgroundColor colors.Color
}

// Badge badge style
type Badge struct {
	Color              colors.Color
	TextColor          colors.Color
	PrimaryColor       colors.Color
	PrimaryTextColor   colors.Color
	SecondaryColor     colors.Color
	SecondaryTextColor colors.Color
	FontWeight         string
}

// BottomNav Bottom Navigation style
type BottomNav struct {
	BackgroundColor    colors.Color
	UnselectedColor    colors.Color
	SelectedColor      colors.Color
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
	TitleColor    colors.Color
	SubtitleColor colors.Color
	FontWeight    int
}

// CardMedia Button style
type CardMedia struct {
	Color         colors.Color
	TitleColor    colors.Color
	SubtitleColor colors.Color
}

// CheckBox Button style
type CheckBox struct {
	BoxColor           colors.Color
	CheckedColor       colors.Color
	RequiredColor      colors.Color
	DisabledColor      colors.Color
	LabelColor         colors.Color
	LabelDisabledColor colors.Color
}

// CardText Button style
type CardText struct {
	TextColor colors.Color
}

// Chip Button style
type Chip struct {
	BackgroundColor colors.Color
	DeleteIconColor colors.Color
	TextColor       colors.Color
	FontSize        int
	FontWeight      int
	Shadow          string
}

// DatePicker Button style
type DatePicker struct {
	Color                       colors.Color
	TextColor                   colors.Color
	CalendarTextColor           colors.Color
	SelectColor                 colors.Color
	SelectTextColor             colors.Color
	CalendarYearBackgroundColor colors.Color
}

// Dialog Button style
type Dialog struct {
	TitleFontSize int
	BodyFontSize  int
	BodyColor     colors.Color
}

// FlatButton Button style
type FlatButton struct {
	Color              colors.Color
	ButtonFilterColor  colors.Color
	DisabledTextColor  colors.Color
	TextColor          colors.Color
	PrimaryTextColor   colors.Color
	SecondaryTextColor colors.Color
	FontSize           int
	FontWeight         int
}

// DropDownMenu Button style
type DropDownMenu struct {
	AccentColor colors.Color
}

// EnhancedButton Button style
type EnhancedButton struct {
	TapHighlightColor colors.Color
}

// FloatingActionButton Button style
type FloatingActionButton struct {
	ButtonSize         int
	MiniSize           int
	Color              colors.Color
	IconColor          colors.Color
	SecondaryColor     colors.Color
	SecondaryIconColor colors.Color
	DisabledTextColor  colors.Color
	DisabledColor      colors.Color
}

// GridTitle Button style
type GridTitle struct {
	TextColor colors.Color
}

// Icon Button style
type Icon struct {
	Color           colors.Color
	BackgroundColor colors.Color
}

// InkBar Button style
type InkBar struct {
	BackgroundColor colors.Color
}

// Drawer Button style
type Drawer struct {
	Color colors.Color
	Width style.Space
}

// ListItem Button style
type ListItem struct {
	NestedLevelDepth   int
	SecondaryTextColor colors.Color
	LeftIconColor      colors.Color
	RightIconColor     colors.Color
}

// Menu Button style
type Menu struct {
	BackgroundColor          colors.Color
	ContainerBackgroundColor colors.Color
}

// MenuItem Button style
type MenuItem struct {
	DataHeight           style.Space
	Height               style.Space
	HoverColor           colors.Color
	Padding              style.Space
	SelectedTextColor    colors.Color
	RightIconDesktopFill colors.Color
}

// MenuSubHeader Button style
type MenuSubHeader struct {
	Padding     style.Space
	BorderColor colors.Color
	TextColor   colors.Color
}

// Overlay Button style
type Overlay struct {
	BackgroundColor colors.Color
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
