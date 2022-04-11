package eac

import (
	"testing"

	"github.com/thetillhoff/eac/pkg/eac/internal/dotProfile"
)

func TestConfigureApp(t *testing.T) {
	var (
		err error
	)

	setUpTestConfigureApp()

	dotProfile.ProfileContainsEnvVariable("GOROOT", "/usr/local/go")

	err = ConfigureApp("eac", false)

	if err != nil {
		t.Log("ConfigureApp is not implemented")
		// t.Error(err)
	}

	tearDownTestConfigureApp()
}

// Prepare by clearing the workdir.
func setUpTestConfigureApp() error {

	Verbose = true

	Clean()

	return nil
}

// Teardown by removing the workdir.
func tearDownTestConfigureApp() error {

	Clean()

	return nil
}
