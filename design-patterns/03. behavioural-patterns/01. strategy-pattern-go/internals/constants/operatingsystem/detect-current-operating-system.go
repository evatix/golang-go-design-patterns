package operatingsystem

import (
	"fmt"
	"runtime"
)

// https://bit.ly/34EaYV2
func WhichOperatingSystem() OperatingSystem {
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
