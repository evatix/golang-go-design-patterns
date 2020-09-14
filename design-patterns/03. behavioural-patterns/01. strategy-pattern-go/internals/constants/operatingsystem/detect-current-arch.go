package operatingsystem

import (
	"fmt"
	"runtime"
	"strings"
)

var x32Architectures = []string{"386", "arm", "armbe", "mips"}
var x64Architectures = []string{"amd64", "arm", "arm64", "ppc64", "mips64le", "sparc64"}

func GetOperatingSystemArchitecture() Architecture {
	arch := runtime.GOARCH
	switch arch {
	case strings.Index()x32Architectures:
		return x64
	case "386":
		return x86
	case "arm64":
		return operatingsystem.Architec

	default:
		message := fmt.Sprintf("Not supported Operating System, Os=%s!", os)
		panic(message)
	}
}