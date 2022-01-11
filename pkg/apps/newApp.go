package apps

import "github.com/thetillhoff/eac/pkg/apps/internal/app"

func newApp(appName string, options ...app.AppOption) app.App {
	return *app.New(appName, options...)
}
