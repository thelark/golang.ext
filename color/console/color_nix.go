// +build !windows

package console

import "fmt"

type Color struct {
	showMode ShowMode
	color    colorCode
	bgcolor  colorCode
}
type colorCode uint

var color = map[string]colorCode{
	"Black":  30,
	"Red":    31,
	"Green":  32,
	"Yellow": 33,
	"Blue":   34,
	"Purple": 35,
	"Cyan":   36,
	"White":  37,

	"BgBlack":  40,
	"BgRed":    41,
	"BgGreen":  42,
	"BgYellow": 43,
	"BgBlue":   44,
	"BgPurple": 45,
	"BgCyan":   46,
	"BgWhite":  47,
}

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

func Init(color, bgcolor colorCode, showMode ShowMode) *Color {
	return &Color{showMode: showMode, color: color, bgcolor: bgcolor}
}
func (c *Color) buildColor() string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, c.showMode, c.bgcolor, c.color, "%s", 0x1B)
}
func (c *Color) Print(s interface{}) {
	fmt.Print(fmt.Sprintf(c.buildColor(), s))
}
func (c *Color) Println(s interface{}) {
	fmt.Println(fmt.Sprintf(c.buildColor(), s))
}
func (c *Color) Printf(format string, args ...interface{}) {
	fmt.Printf(fmt.Sprintf(c.buildColor(), fmt.Sprintf(format, args...)))
}
