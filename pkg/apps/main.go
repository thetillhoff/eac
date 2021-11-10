package apps

import "github.com/thetillhoff/eac/internal/appVersions"

var (
	AppsDirPath      string
	VersionsFilePath string
)

func OverrideDefaults(appsDirPath string, versionsFilePath string) {
	if appsDirPath != "" && appsDirPath != AppsDirPath {
		AppsDirPath = appsDirPath
		appVersions.AppsDirPath = appsDirPath
	}

	if versionsFilePath != "" && versionsFilePath != VersionsFilePath {
		VersionsFilePath = versionsFilePath
		appVersions.VersionsFilePath = versionsFilePath
	}
}
