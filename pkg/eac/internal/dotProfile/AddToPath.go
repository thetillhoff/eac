package dotProfile

import (
	"errors"
	"log"
	"os"
	"os/user"
	"path"
	"runtime"
)

// AddtoPath adds the pathToAdd to the env $PATH.
// This will do nothing it is already part of it.
func AddToPath(pathToAdd string) error {
	var (
		err          error
		home         string
		actualUser   *user.User
		pathContains bool
	)

	// Fail if not on linux or darwin
	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		return errors.New("unsupported operating system")
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
		return errors.New("`~/.profile` does not exist")
	}

	pathContains, err = PathContains(pathToAdd)
	if err != nil {
		return err
	}

	if pathContains { // Nothing to do
		return nil
	} else { // Add pathToAdd to $PATH
		// TODO implement here
		return errors.New("not implemented")

	}
}
