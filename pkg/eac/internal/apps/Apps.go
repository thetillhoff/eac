package apps

import (
	"os"

	"github.com/thetillhoff/eac/pkg/eac/internal/version"
)

var (
	Apps map[string]App = map[string]App{}
)

type App struct {
	Destination           string
	DownloadUrlTemplate   string
	Checksum              *Checksum
	VersionProvider       version.VersionProvider
	TarProperties         *TarProperties
	ZipProperties         *ZipProperties
	PermissionAdjustments []PermissionAdjustment
	PathAddition          string
	ProfileVariables      []EnvironmentVariables
}

type Checksum struct { // sha256 only
	UrlTemplate      string
	ExtendedChecksum bool // checksum with filename, multiline possible
}

type TarProperties struct {
	Path string
}

type ZipProperties struct {
	Path string
}

type PermissionAdjustment struct {
	Path string
	Mode os.FileMode
}

type EnvironmentVariables struct {
	Key   string
	Value string
}
