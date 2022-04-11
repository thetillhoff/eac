package dotProfile

import (
	"bufio"
	"errors"
	"log"
	"os"
	"os/user"
	"path"
	"runtime"
	"strings"
)

// ProfileContainsEnvVariables checks if `~/.profile` contains an environment variable named with key is already set to the value.
func ProfileContainsEnvVariable(key string, value string) (bool, error) {
	var (
		err             error
		home            string
		actualUser      *user.User
		profileContents = ""
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
	if _, err = os.Stat(path.Join(home, ".profile")); os.IsNotExist(err) {
		return false, errors.New("`~/.profile` does not exist")
	}

	proFile, err := os.Open(path.Join(home, ".profile")) // smol joke ;)
	if err != nil {
		return false, err
	}
	defer proFile.Close()

	scanner := bufio.NewScanner(proFile)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		profileContents = profileContents + scanner.Text()
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	profileLines := strings.Split(profileContents, "\n")

	for _, line := range profileLines {
		if strings.Contains(line, "export "+key+"=") { // Variable is set (unkown value)
			if strings.Contains(line, "export "+key+"="+value) { // Variable set correctly
				return true, nil
			} else { // Variable is set, but wrong value
				return false, nil
			}
		}
	}

	return false, nil // Variable is not set

}
