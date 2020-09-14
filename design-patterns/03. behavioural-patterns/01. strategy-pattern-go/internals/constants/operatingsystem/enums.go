package operatingsystem

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

type Detail struct {
	Os              OperatingSystem
	OsGroup         Group
	Architecture    Architecture
	RawArchitecture string
	NewLine         string
	PathSeparator   string
}

type Group int

const (
	WindowsGroup Group = iota
	UNIXGroup
	AndroidGroup
	JSGroup
)

type Architecture int

const (
	x86 Architecture = iota
	x64
)
