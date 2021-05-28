package app

type App struct {
	Name                   string // name of the app (must be equal to folder name)
	getLocalVersionScript  string // script that checks whether the app is installed and if yes, in which version
	installScript          string // script that runs the installation of the app
	configureScript        string // script that configures the app (if installed)
	getLatestVersionScript string // script that checks for new versions of the app
	uninstallScript        string // script that uninstalls the app
	continueOnError        bool   // if true, only this app will be skipped
	shell                  string // set specific shell to use for the scripts of this app
	wantedVersion          string // version of this app, which is desired to be installed
	localVersion           string // version of this app, which is currently installed
}

type AppOption func(*App)

func Shell(s string) AppOption {
	return func(a *App) {
		if s != "" {
			a.Name = s
		}
	}
}
func GetLocalVersionScript(s string) AppOption {
	return func(a *App) {
		if s != "" {
			a.getLocalVersionScript = s
		}
	}
}
func GetLatestVersionScript(s string) AppOption {
	return func(a *App) {
		if s != "" {
			a.getLatestVersionScript = s
		}
	}
}
func InstallScript(s string) AppOption {
	return func(a *App) {
		if s != "" {
			a.installScript = s
		}
	}
}
func UnstallScript(s string) AppOption {
	return func(a *App) {
		if s != "" {
			a.uninstallScript = s
		}
	}
}
func ConfigureScript(s string) AppOption {
	return func(a *App) {
		if s != "" {
			a.configureScript = s
		}
	}
}
func ContinueOnError(b bool) AppOption {
	return func(a *App) {
		if b {
			a.continueOnError = b
		}
	}
}
func WantedVersion(s string) AppOption {
	return func(a *App) {
		if s != "" {
			a.wantedVersion = s
		}
	}
}

// New creates a new App with some default settings.
func New(appName string, options ...AppOption) *App {

	appItem := &App{ // initialize with default values
		shell:                  "/bin/sh",
		Name:                   appName,
		getLocalVersionScript:  "get-local-version.sh",
		installScript:          "install.sh",
		configureScript:        "configure.sh",
		getLatestVersionScript: "get-latest-version.sh",
		uninstallScript:        "uninstall.sh",
		continueOnError:        false,
		wantedVersion:          "",
	}

	for _, opt := range options { // set custom properties
		opt(appItem)
	}

	return appItem
}
