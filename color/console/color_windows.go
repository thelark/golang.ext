// +build windows

package console

import (
	"fmt"
	"syscall"
)

var (
	kernel32DLL                 = syscall.NewLazyDLL("kernel32.dll")
	setConsoleTextAttributeProc = kernel32DLL.NewProc("SetConsoleTextAttribute")
)

type Color struct {
	colorCode colorCode
}

type colorCode uintptr

var color = map[string]colorCode{
	"Black":       0,
	"BlueDark":    1,
	"GreenDark":   2,
	"CyanDark":    3,
	"RedDark":     4,
	"PurpleDark":  5,
	"YellowDark":  6,
	"WhiteDark":   7,
	"Gray":        8,
	"BlueLight":   9,
	"GreenLight":  10,
	"CyanLight":   11,
	"RedLight":    12,
	"PurpleLight": 13,
	"YellowLight": 14,
	"White":       15,


	"BgBlack":       0x10 * 0,
	"BgBlueDark":    0x10 * 1,
	"BgGreenDark":   0x10 * 2,
	"BgCyanDark":    0x10 * 3,
	"BgRedDark":     0x10 * 4,
	"BgPurpleDark":  0x10 * 5,
	"BgYellowDark":  0x10 * 6,
	"BgWhiteDark":   0x10 * 7,
	"BgGray":        0x10 * 8,
	"BgBlueLight":   0x10 * 9,
	"BgGreenLight":  0x10 * 10,
	"BgCyanLight":   0x10 * 11,
	"BgRedLight":    0x10 * 12,
	"BgPurpleLight": 0x10 * 13,
	"BgYellowLight": 0x10 * 14,
	"BgWhite":       0x10 * 15,
}

func Restore() {
	handle, _, _ := setConsoleTextAttributeProc.Call(uintptr(syscall.Stdout), uintptr(color["WhiteDark"]|color["BgBlack"]))
	CloseHandle := kernel32DLL.NewProc("CloseHandle")
	CloseHandle.Call(handle)
}
func SetColor(c colorCode) {
	setConsoleTextAttributeProc.Call(uintptr(syscall.Stdout), uintptr(c))
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

func Init(color, bgColor colorCode, showMode ShowMode) *Color {
	return &Color{colorCode: color | bgColor}
}
func (c *Color) Print(s interface{}) {
	SetColor(c.colorCode)
	fmt.Print(s)
	Restore()
}
func (c *Color) Println(s interface{}) {
	SetColor(c.colorCode)
	fmt.Println(s)
	Restore()
}
func (c *Color) Printf(format string, args ...interface{}) {
	SetColor(c.colorCode)
	fmt.Printf(format, args...)
	Restore()
}
