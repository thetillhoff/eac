package apps

import "github.com/thetillhoff/eac/internal/app"

func newApp(appName string, options ...app.AppOption) app.App {
	return *app.New(appName, options...)
}
