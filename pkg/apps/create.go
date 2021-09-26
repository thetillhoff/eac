package apps

import (
	"os"
	"path"

	"github.com/thetillhoff/eac/internal/templates"
	"github.com/thetillhoff/eac/pkg/logs"
)

func Create(appNames []string, flaggedPlatforms []string, appsDirPath string, verbose bool, createData map[string]string) {
	logs.Verbose = verbose
	platforms := ResolvePlatforms(flaggedPlatforms)
	logs.Info("Selected platforms:", platforms)

	if _, err := os.Stat(appsDirPath); os.IsNotExist(err) {
		logs.Err("Apps folder at '" + appsDirPath + "' doesn't exist.\nPlease run 'eac init' first.")
	} else if err != nil {
		logs.Err("There was an error while accessing the appsDir at '"+appsDirPath+"':", err)
	}

	for _, appName := range appNames {
		appPath := path.Join(appsDirPath, appName)
		if _, err := os.Stat(appPath); os.IsNotExist(err) { // if folder doesn't exist yet
			err := os.Mkdir(appPath, os.ModePerm)
			if err != nil {
				logs.Err("There was an error while creating the app at '"+appPath+"':", err)
			}
			logs.Success("Created app '" + appName + "'.")
		} else if err == nil { // if folder does exist
			logs.Warn("App '" + appName + "' does already exist.")
			return
		} else {
			logs.Err("There was an error while accessing the app at '"+appPath+"':", err)
		}

		for _, platform := range platforms {
			platformPath := path.Join(appPath, platform)
			if _, err := os.Stat(platformPath); os.IsNotExist(err) {
				err := os.Mkdir(platformPath, os.ModePerm) // ignore errors
				if err != nil {
					logs.Err("There was an error while creating the platform '"+platform+"' for app '"+appPath+"' at '"+platformPath+"':", err)
				}
				logs.Info("Created '" + platformPath + "' folder.")
			} else if err == nil {
				logs.Warn("Platform '" + platform + "' for app '" + appName + "' does already exist.")
			} else {
				logs.Err("There was an error while accessing the platform '"+platform+"' for app '"+appPath+"' at '"+platformPath+"':", err)
			}

			if githubUser, ok := createData["githubUser"]; ok { // add github files
				data := map[string]string{
					"name": appName,
					"user": githubUser,
				}
				switch platform {
				case "windows":
					templates.WriteTemplateToFile(path.Join(platform, "github-install.ps1"), data, path.Join(platformPath, "install.ps1"))
					templates.WriteTemplateToFile(path.Join(platform, "default-uninstall.ps1"), data, path.Join(platformPath, "uninstall.ps1"))
					templates.WriteTemplateToFile(path.Join(platform, "github-getLatestVersion.ps1"), data, path.Join(platformPath, "getLatestVersion.ps1"))
					templates.WriteTemplateToFile(path.Join(platform, "default-getLocalVersion.ps1"), data, path.Join(platformPath, "getLocalVersion.ps1"))
					templates.WriteTemplateToFile(path.Join(platform, "default-configure.ps1"), data, path.Join(platformPath, "configure.ps1"))
				default:
					templates.WriteTemplateToFile(path.Join(platform, "github-install.sh"), data, path.Join(platformPath, "install.sh"))
					templates.WriteTemplateToFile(path.Join(platform, "default-uninstall.sh"), data, path.Join(platformPath, "uninstall.sh"))
					templates.WriteTemplateToFile(path.Join(platform, "github-getLatestVersion.sh"), data, path.Join(platformPath, "getLatestVersion.sh"))
					templates.WriteTemplateToFile(path.Join(platform, "default-getLocalVersion.sh"), data, path.Join(platformPath, "getLocalVersion.sh"))
					templates.WriteTemplateToFile(path.Join(platform, "default-configure.sh"), data, path.Join(platformPath, "configure.sh"))
				}
			} else { // default files
				data := map[string]string{
					"name": appName,
				}
				switch platform {
				case "windows":
					templates.WriteTemplateToFile(path.Join(platform, "default-install.ps1"), data, path.Join(platformPath, "install.ps1"))
					templates.WriteTemplateToFile(path.Join(platform, "default-uninstall.ps1"), data, path.Join(platformPath, "uninstall.ps1"))
					templates.WriteTemplateToFile(path.Join(platform, "default-getLatestVersion.ps1"), data, path.Join(platformPath, "getLatestVersion.ps1"))
					templates.WriteTemplateToFile(path.Join(platform, "default-getLocalVersion.ps1"), data, path.Join(platformPath, "getLocalVersion.ps1"))
					templates.WriteTemplateToFile(path.Join(platform, "default-configure.ps1"), data, path.Join(platformPath, "configure.ps1"))
				default:
					templates.WriteTemplateToFile(path.Join(platform, "default-install.sh"), data, path.Join(platformPath, "install.sh"))
					templates.WriteTemplateToFile(path.Join(platform, "default-uninstall.sh"), data, path.Join(platformPath, "uninstall.sh"))
					templates.WriteTemplateToFile(path.Join(platform, "default-getLatestVersion.sh"), data, path.Join(platformPath, "getLatestVersion.sh"))
					templates.WriteTemplateToFile(path.Join(platform, "default-getLocalVersion.sh"), data, path.Join(platformPath, "getLocalVersion.sh"))
					templates.WriteTemplateToFile(path.Join(platform, "default-configure.sh"), data, path.Join(platformPath, "configure.sh"))
				}
			}

			logs.Success("Created platform '" + platform + "' for app '" + appName + "'.")
		}
	}
}
