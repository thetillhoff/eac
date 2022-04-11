package eac

import (
	"testing"
)

func TestGetLatestVersionJson(t *testing.T) {
	var (
		err     error
		version string
	)

	version, err = GetLatestVersion("eac")

	// Success if err == nil and version != ""

	if err != nil {
		t.Error(err)
	}
	if version == "" {
		t.Error("Empty version returned.")
	}
}

func TestGetLatestVersionPlaintext(t *testing.T) {
	var (
		err     error
		version string
	)

	version, err = GetLatestVersion("kubectl")

	// Success if err == nil and version != ""

	if err != nil {
		t.Error(err)
	}
	if version == "" {
		t.Error("Empty version returned.")
	}
}
