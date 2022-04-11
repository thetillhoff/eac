package eac

import (
	"testing"
)

func TestListDownloadedAppsWithoutVersions(t *testing.T) {
	var (
		err     error
		appList []string
	)

	setUpTestListDownloadedApps()

	appList, err = ListDownloadedApps(false)

	// Success if err == nil and version != ""

	if err != nil {
		t.Error(err)
	}
	if len(appList) != 1 {
		t.Error("Number of downloaded apps is", len(appList), "expected 1.")
	}

	tearDownTestListDownloadedApps()
}

func TestListDownloadedAppsWithVersions(t *testing.T) {
	var (
		err     error
		appList []string
	)

	setUpTestListDownloadedApps()

	appList, err = ListDownloadedApps(true)

	// Success if err == nil and version != ""

	if err != nil {
		t.Error(err)
	}
	if len(appList) != 1 {
		t.Error("Number of downloaded apps is", len(appList), "expected 1.")
	}

	tearDownTestListDownloadedApps()
}

// Prepare by clearing the workdir.
func setUpTestListDownloadedApps() error {
	var (
		err     error
		version string
	)

	Verbose = true

	Clean()

	version, err = GetLatestVersion("eac")
	if err != nil {
		return err
	}
	err = InstallApp("eac", version)

	return err
}

// Teardown by removing the workdir.
func tearDownTestListDownloadedApps() error {

	Clean()

	return nil
}
