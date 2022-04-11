package eac

import (
	"testing"
)

func TestListAvailableApps(t *testing.T) {
	var (
		err     error
		appList []string
	)

	setUpTestListAvailableApps()

	appList, err = ListAvailableApps()

	// Success if err == nil and version != ""

	if err != nil {
		t.Error(err)
	}
	if len(appList) == 0 {
		t.Error("Number of available apps is 0.")
	}

	tearDownTestListAvailableApps()
}

// Prepare by clearing the workdir.
func setUpTestListAvailableApps() error {

	Verbose = true

	Clean()

	return nil
}

// Teardown by removing the workdir.
func tearDownTestListAvailableApps() error {

	Clean()

	return nil
}
