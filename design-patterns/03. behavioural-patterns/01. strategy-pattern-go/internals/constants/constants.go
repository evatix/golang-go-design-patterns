package constants

import (
	"fmt"
	"runtime"
)

const (
	// need to optimize this to https://stackoverflow.com/questions/14493867/what-is-the-most-portable-cross-platform-way-to-represent-a-newline-in-go-golang
	WindowsNewLineSeparator = "\r\n"
	NewLineSeparator        = "\n"
	Coma                    = ","
	ComaSpace               = ", "
	Dot                     = "."
	DotSpace                = ". "
	SemiColon               = ";"
	SemiColonSpace          = "; "
	SemiColonNewLine        = ";\n"
	ComaNewLine             = ",\n"
	// |
	Pipe                 = "|"
	PipeSpace            = " | "
	DoublePipe           = "||"
	WindowsPathSeparator = "\\"
	UnixPathSeparator    = "/"
)

type OperatingSystem int

const (
	Debian OperatingSystem = iota
	Linux
	// Darwin is Mac OS or iOS
	DarwinOrMacOrIOS
	Ubuntu
	Android
	Dragonfly
	FreeBSD
	Nacl
	OpenBSD
	Js
	NetBSD
	Solaris
	Windows
)

type OperatingSystemDetail struct {
	Os              OperatingSystem
	OsGroup         OperatingSystemGroup
	Architecture    OperatingSystemArchitecture
	RawArchitecture string
	NewLine         string
	PathSeparator   string
}

type OperatingSystemGroup int

const (
	WindowsGroup OperatingSystemGroup = iota
	UNIXGroup
	AndroidGroup
	JSGroup
)

type OperatingSystemArchitecture int

const (
	x86 OperatingSystemArchitecture = iota
	x64
)

// https://bit.ly/34EaYV2
func RevealCurrentOperatingSystem() OperatingSystem {
	os := runtime.GOOS
	switch os {
	case "windows":
		return Windows
	case "darwin":
		return DarwinOrMacOrIOS
	case "linux":
		return Linux
	case "dragonfly":
		return Dragonfly
	case "freebsd":
		return FreeBSD
	case "js":
		return Js
	case "netbsd":
		return NetBSD
	case "openbsd":
		return OpenBSD
	default:
		message := fmt.Sprintf("Not supported Operating System, Os=%s!", os)
		panic(message)
	}
}

//func GetOperatingSystemDetail() OperatingSystemDetail {
//
//}

//func EnvNewLine() string {
//	if(os.)
//}
