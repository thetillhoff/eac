package dotProfile

import (
	"errors"
	"log"
	"os"
	"os/user"
	"path"
	"runtime"
)

// PathContains checks whether the pathToCheck is already part of $PATH.
func PathContains(pathToCheck string) (bool, error) {
	var (
		err        error
		home       string
		actualUser *user.User
	)

	// Fail if not on linux or darwin
	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		return false, errors.New("unsupported operating system")
	}

	// Sometimes, we need to work around that to get the actual user home
	sudoUser := os.Getenv("SUDO_USER")
	if sudoUser == "" { // If no sudo is involved
		home, err = os.UserHomeDir() // Get user's $HOME.
		if err != nil {
			log.Fatalln(err)
		}
	} else { // `eac` was ran with sudo
		actualUser, err = user.Lookup(sudoUser)
		if err != nil {
			log.Fatalln(err)
		}
		home = actualUser.HomeDir
	}

	// Fail if $HOME/.profile does not exist
	if _, err := os.Stat(path.Join(home, ".profile")); os.IsNotExist(err) {
		return false, errors.New("`~/.profile` does not exist")
	}

	return false, errors.New("not implemented")

}
