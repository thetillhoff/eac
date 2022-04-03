package helpers

import "errors"

func EnvPathExists(pathToCheck string) (bool, error) {

	return false, errors.New("not implemented")

	// Fail if not on linux or darwin

	// Fail if $HOME/.profile does not exist

}
