package eac

import (
	"testing"
)

func TestDownloadAppVersion(t *testing.T) {
	var (
		err      error
		filepath string
	)

	setUpTestDownloadAppVersion()

	filepath, err = DownloadAppVersion("eac", "0.0.1")

	// Success if err == nil and filepath == /tmp/eac/eac/0.0.1/eac_linux_amd64

	if err != nil {
		t.Error(err)
	}
	if filepath != "/tmp/eac/eac/0.0.1/eac_linux_amd64" { // Don't get confused with the /eac/eac part: The first is part of the workdir and the second the appName
		t.Error("Bad filename returned:", filepath)
	}

	tearDownTestDownloadAppVersion()
}

func TestDownloadAppVersionTar(t *testing.T) {
	var (
		err      error
		filepath string
	)

	setUpTestDownloadAppVersion()

	filepath, err = DownloadAppVersion("eac", "0.1.0")

	// Success if err == nil and filepath == /tmp/eac/eac/0.0.1/eac_linux_amd64

	if err != nil {
		t.Error(err)
	}
	if filepath != "/tmp/eac/eac/0.1.0/eac_linux_amd64" { // Don't get confused with the /eac/eac part: The first is part of the workdir and the second the appName
		t.Error("Bad filename returned. Expected /tmp/eac/eac/0.1.0/eac_linux_amd64, got", filepath)
	}

	tearDownTestDownloadAppVersion()
}

func TestDownloadAppVersionTarGzAndChecksum(t *testing.T) {
	var (
		err      error
		filepath string
	)

	setUpTestDownloadAppVersion()

	filepath, err = DownloadAppVersion("helm", "3.8.1")

	// Success if err == nil and filepath == /tmp/eac/helm/3.8.1/linux-amd64/helm

	if err != nil {
		t.Error(err)
	}
	if filepath != "/tmp/eac/helm/3.8.1/linux-amd64/helm" {
		t.Error("Bad filename returned. Expected /tmp/eac/helm/3.8.1/linux-amd64/helm, got", filepath)
	}

	tearDownTestDownloadAppVersion()
}

func TestDownloadAppVersionZipAndChecksum(t *testing.T) {
	var (
		err      error
		filepath string
	)

	setUpTestDownloadAppVersion()

	filepath, err = DownloadAppVersion("terraform", "1.1.7")

	// Success if err == nil and filepath == /tmp/eac/terraform/1.1.7/terraform

	if err != nil {
		t.Error(err)
	}
	if filepath != "/tmp/eac/terraform/1.1.7/terraform" {
		t.Error("Bad filename returned. Expected /tmp/eac/terraform/1.1.7/terraform, got", filepath)
	}

	tearDownTestDownloadAppVersion()
}

// Prepare by clearing the workdir.
func setUpTestDownloadAppVersion() error {

	Verbose = true

	Clean()

	return nil
}

// Teardown by removing the workdir.
func tearDownTestDownloadAppVersion() error {

	Clean()

	return nil
}
