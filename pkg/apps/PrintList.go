package apps

import (
	"fmt"

	"github.com/thetillhoff/eac/internal/config"
)

// List the (remotely) available apps
func PrintList(conf config.Config) {
	for uniqueAppName := range listRemote(conf) {
		fmt.Println(uniqueAppName)
	}
}
