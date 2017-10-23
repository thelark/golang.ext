package console

import (
	"runtime"
)

var flag = false

func init() {
	if runtime.GOOS == "windows" {
		// Windows 操作系统
		flag = true
	}
}

type colorInterface interface {
	Println(s interface{})
	Printf(format string, args ...interface{})
}

func Black() colorInterface {
	if flag {
		return Init(color["Black"], color["BgBlack"], R)
	}
	return Init(color["Black"], color["BgBlack"], R)
}
func Red() colorInterface {
	if flag {
		return Init(color["RedLight"], color["BgBlack"], R)
	}
	return Init(color["Red"], color["BgBlack"], B)
}
func Green() colorInterface {
	if flag {
		return Init(color["GreenLight"], color["BgBlack"], R)
	}
	return Init(color["Green"], color["BgBlack"], B)
}
func Yellow() colorInterface {
	if flag {
		return Init(color["YellowLight"], color["BgBlack"], R)
	}
	return Init(color["Yellow"], color["BgBlack"], B)
}
func Blue() colorInterface {
	if flag {
		return Init(color["BlueLight"], color["BgBlack"], R)
	}
	return Init(color["Blue"], color["BgBlack"], B)
}
func Magenta() colorInterface {
	if flag {
		return Init(color["PurpleLight"], color["BgBlack"], R)
	}
	return Init(color["Purple"], color["BgBlack"], B)
}
func Cyan() colorInterface {
	if flag {
		return Init(color["CyanLight"], color["BgBlack"], R)
	}
	return Init(color["Cyan"], color["BgBlack"], B)
}
func White() colorInterface {
	if flag {
		return Init(color["White"], color["BgBlack"], R)
	}
	return Init(color["White"], color["BgBlack"], B)
}
func Gray() colorInterface {
	if flag {
		return Init(color["Gray"], color["BgBlack"], R)
	}
	return Init(color["Black"], color["BgBlack"], B)
}
