package linux

import "fmt"

type colorCode uint
const (
	Black  colorCode = iota + 30
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	White
)
const (
	BgBlack  colorCode = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgPurple
	BgCyan
	BgWhite
)

type ShowMode uint
const (
	// R reset emphasis style
	R ShowMode = 0
	// B bold emphasis style
	B ShowMode = 1
	// D dim emphasis style
	D ShowMode = 2
	// I italic emphasis style
	I ShowMode = 3
	// U underline emphasis style
	U ShowMode = 4
	// In inverse emphasis style
	In ShowMode = 7
	// H hidden emphasis style
	H ShowMode = 8
	// S strikeout emphasis style
	S ShowMode = 9
)

type Color struct {
	showMode ShowMode
	color    colorCode
	bgcolor  colorCode
}

func Init(color, bgcolor colorCode, showMode ShowMode) *Color {
	return &Color{showMode: showMode, color: color, bgcolor: bgcolor}
}
func (c *Color) buildColor() string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, c.showMode, c.bgcolor, c.color, "%s", 0x1B)
}
func (c *Color) Println(s interface{}) {
	fmt.Println(fmt.Sprintf(c.buildColor(), s))
}
func (c *Color) Printf(format string, args ...interface{}) {
	fmt.Printf(fmt.Sprintf(c.buildColor(), fmt.Sprintf(format, args...)))
}

