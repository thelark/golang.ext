package windows

import (
	"syscall"
	"fmt"
)

type Color struct {
	colorCode colorCode
}

type colorCode uintptr

const (
	Black       colorCode = iota
	BlueDark
	GreenDark
	CyanDark
	RedDark
	PurpleDark
	YellowDark
	WhiteDark
	Gray
	BlueLight
	GreenLight
	CyanLight
	RedLight
	PurpleLight
	YellowLight
	White
)
const (
	BgBlack       colorCode = iota * 0x10
	BgBlueDark
	BgGreenDark
	BgCyanDark
	BgRedDark
	BgPurpleDark
	BgYellowDark
	BgWhiteDark
	BgGray
	BgBlueLight
	BgGreenLight
	BgCyanLight
	BgRedLight
	BgPurpleLight
	BgYellowLight
	BgWhite
)

var kernel32 *syscall.LazyDLL
var proc *syscall.LazyProc

func init() {
	kernel32 = syscall.NewLazyDLL("kernel32.dll")
	proc = kernel32.NewProc("SetConsoleTextAttribute")
}

func Restore() {
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(WhiteDark|BgBlack))
	CloseHandle := kernel32.NewProc("CloseHandle")
	CloseHandle.Call(handle)
}
func SetColor(c colorCode) {
	proc.Call(uintptr(syscall.Stdout), uintptr(c))
}
func Init(color, bgColor colorCode) *Color {
	return &Color{colorCode: color | bgColor}
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
