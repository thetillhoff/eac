package eac

import (
	"testing"
)

func TestInstallApp(t *testing.T) {

	setUpInstallApp()

	var err = InstallApp("eac", "0.1.0")

	// Success if err == nil and version != ""

	if err != nil {
		t.Error(err)
	}

	// TODO test whether correct version is installed now. This requires to remove it beforehand in the setUp func.

	tearDownInstallApp()

}

// Prepare by clearing the workdir.
func setUpInstallApp() error {

	Verbose = true

	Clean()

	return nil
}

// Teardown by removing the workdir.
func tearDownInstallApp() error {

	Clean()

	return nil
}
