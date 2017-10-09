package cmd

import (
	"runtime"
	"thelark.cn/golang.ext/color/cmd/windows"
	"thelark.cn/golang.ext/color/cmd/linux"
)

var flag = false

func init() {
	if runtime.GOOS == "windows" {
		// Windows 操作系统
		flag = true
	}
}

type color interface {
	Println(s interface{})
	Printf(format string, args ...interface{})
}

func Black() color {
	if flag {
		return windows.Init(windows.Black, windows.BgBlack)
	}
	return linux.Init(linux.Black, linux.BgBlack, linux.R)
}
func Red() color {
	if flag {
		return windows.Init(windows.RedLight, windows.BgBlack)
	}
	return linux.Init(linux.Red, linux.BgBlack, linux.B)
}
func Green() color {
	if flag {
		return windows.Init(windows.GreenLight, windows.BgBlack)
	}
	return linux.Init(linux.Green, linux.BgBlack, linux.B)
}
func Yellow() color {
	if flag {
		return windows.Init(windows.YellowLight, windows.BgBlack)
	}
	return linux.Init(linux.Yellow, linux.BgBlack, linux.B)
}
func Blue() color {
	if flag {
		return windows.Init(windows.BlueLight, windows.BgBlack)
	}
	return linux.Init(linux.Blue, linux.BgBlack, linux.B)
}
func Magenta() color {
	if flag {
		return windows.Init(windows.PurpleLight, windows.BgBlack)
	}
	return linux.Init(linux.Purple, linux.BgBlack, linux.B)
}
func Cyan() color {
	if flag {
		return windows.Init(windows.CyanLight, windows.BgBlack)
	}
	return linux.Init(linux.Cyan, linux.BgBlack, linux.B)
}
func White() color {
	if flag {
		return windows.Init(windows.White, windows.BgBlack)
	}
	return linux.Init(linux.White, linux.BgBlack, linux.B)
}
func Gray() color {
	if flag {
		return windows.Init(windows.Gray, windows.BgBlack)
	}
	return linux.Init(linux.Black, linux.BgBlack, linux.B)
}
