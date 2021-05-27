package app

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

var (
	platform               = runtime.GOOS // name of the current platform. Folder or symlink with equal name must exist in apps/<appname>/.
	shell           string = "/bin/sh"    // default
	continueOnError bool   = false        // default
)

type App struct {
	name                   string // name of the app (must be equal to folder name)
	getLocalVersionScript  string // script that checks whether the app is installed and if yes, in which version
	installScript          string // script that runs the installation of the app
	configureScript        string // script that configures the app (if installed)
	getLatestVersionScript string // script that checks for new versions of the app
	uninstallScript        string // script that uninstalls the app
	continueOnError        bool   // if true, only this app will be skipped
	shell                  string // set specific shell to use for the scripts of this app
}

// New creates a new App with some default settings.
func NewApp(name string) App {
	return App{
		name:                   name,
		getLocalVersionScript:  "get-local-version.sh",
		installScript:          "install.sh",
		configureScript:        "configure.sh",
		getLatestVersionScript: "get-latest-version.sh",
		uninstallScript:        "uninstall.sh",
		continueOnError:        continueOnError,
		shell:                  shell,
	}
}

func GetLocalVersion(app App) (string, error) {
	return RunScript(app.shell, app.name, app.getLocalVersionScript, app.continueOnError)
}

func GetLatestVersion(app App) (string, error) {
	return RunScript(app.shell, app.name, app.getLatestVersionScript, app.continueOnError)
}

func Install(app App) (string, error) {
	return RunScript(app.shell, app.name, app.installScript, app.continueOnError)
}

func Configure(app App) (string, error) {
	return RunScript(app.shell, app.name, app.configureScript, app.continueOnError)
}

func Uninstall(app App) (string, error) {
	return RunScript(app.shell, app.name, app.uninstallScript, app.continueOnError)
}

func Update(app App) (string, error) {
	fmt.Println("Checking " + app.name + " version ...")
	localVersion, err := GetLocalVersion(app)
	if err != nil {
		return "", err
	}
	if localVersion == "" {
		return "", errors.New("Retrieval of local version of app '" + app.name + "' failed.")
	}
	localVersion = strings.TrimSuffix(localVersion, "\n") // remove line-break after content

	latestVersion, err := GetLatestVersion(app)
	if err != nil {
		return "", err
	}
	if latestVersion == "" {
		return "", errors.New("Retrieval of latest version of app '" + app.name + "' failed.")
	}
	latestVersion = strings.TrimSuffix(latestVersion, "\n") // remove line-break after content

	if localVersion == latestVersion {
		return "Already installed with the latest version (v" + localVersion + ").", nil
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Do you want to update from v" + localVersion + " to v" + latestVersion + "? [y/n] ")
		text, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		text = strings.TrimSuffix(text, "\n") // remove line-break after input
		if text == "y" {
			//TODO: parse version file, update version of app, save yaml back to file
			return "Updated from v" + localVersion + " to v" + latestVersion + ".", nil
		} else {
			return "Skipped " + app.name + ".", nil
		}
	}
}

func Validate(app App) (string, error) {
	out := ""

	scriptpaths := []string{}
	scriptpaths = append(scriptpaths, path.Join("apps", app.name, platform, app.getLocalVersionScript))
	scriptpaths = append(scriptpaths, path.Join("apps", app.name, platform, app.installScript))
	scriptpaths = append(scriptpaths, path.Join("apps", app.name, platform, app.configureScript))
	scriptpaths = append(scriptpaths, path.Join("apps", app.name, platform, app.getLatestVersionScript))
	scriptpaths = append(scriptpaths, path.Join("apps", app.name, platform, app.uninstallScript))

	missingFiles := testFiles(scriptpaths...)

	if len(missingFiles) == 0 {
		out = out + "The scripts for '" + app.name + "' exist.\n"
	} else {
		err := "The following files for '" + app.name + "' don't seem to exist:\n"
		for _, missingFile := range missingFiles {
			err = err + missingFile + "\n"
		}
		return "", errors.New(err)
	}

	//TODO check whether get-latest-version and get-local-version return single line strings

	return out, nil
}

func RunScript(appshell string, appname string, script string, appcontinueOnError bool) (string, error) {
	cmd := exec.Command(appshell, path.Join("apps", appname, platform, script))
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func testFiles(filepaths ...string) []string {

	missingFiles := []string{}

	for _, filepath := range filepaths {
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			missingFiles = append(missingFiles, filepath)
		}
	}

	return missingFiles
}
