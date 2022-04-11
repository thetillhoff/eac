package eac

import (
	"os"
	"path"
	"testing"
)

func TestCleanNoArgs(t *testing.T) {
	setUpClean()

	Clean()

	// Success if /tmp/eac/ doesn't exist any  more

	if _, err := os.Stat("/tmp/eac"); !os.IsNotExist(err) {
		t.Error("Failed to `eac clean`")
	}

	tearDownClean()
}

func TestCleanOneArg(t *testing.T) {
	setUpClean()

	Clean("a")

	// Success if /tmp/eac/a/ doesn't exist any more, but /tmp/eac/b/ and /tmp/eac/c/ still exist

	if _, err := os.Stat("/tmp/eac/a"); !os.IsNotExist(err) {
		t.Error("Failed to `eac clean a`; /tmp/eac/a still exists")
	}
	if _, err := os.Stat("/tmp/eac/b"); os.IsNotExist(err) {
		t.Error("Failed to `eac clean a`, /tmp/eac/b doesn't exist")
	}
	if _, err := os.Stat("/tmp/eac/c"); os.IsNotExist(err) {
		t.Error("Failed to `eac clean a`, /tmp/eac/c doesn't exist")
	}

	tearDownClean()
}

func TestCleanTwoArgs(t *testing.T) {
	setUpClean()

	Clean("a", "b")

	// Success if /tmp/eac/a/ and /tmp/eac/b/ don't exist any more, but /tmp/eac/c/ still exists

	if _, err := os.Stat("/tmp/eac/a"); !os.IsNotExist(err) {
		t.Error("Failed to `eac clean a`; /tmp/eac/a still exists")
	}
	if _, err := os.Stat("/tmp/eac/b"); !os.IsNotExist(err) {
		t.Error("Failed to `eac clean a`, /tmp/eac/b still exist")
	}
	if _, err := os.Stat("/tmp/eac/c"); os.IsNotExist(err) {
		t.Error("Failed to `eac clean a`, /tmp/eac/c doesn't exist")
	}

	tearDownClean()
}

// Prepare by clearing the workdir
// and by creating folders `workdir/a`, `workdir/b`, and `workdir/c`.
func setUpClean() error {
	var (
		err error
	)

	Verbose = true

	os.RemoveAll(workdir)

	err = os.MkdirAll(path.Join(workdir, "a"), os.ModeDir)
	if err != nil {
		return err
	}
	err = os.MkdirAll(path.Join(workdir, "b"), os.ModeDir)
	if err != nil {
		return err
	}
	err = os.MkdirAll(path.Join(workdir, "c"), os.ModeDir)
	if err != nil {
		return err
	}

	return nil
}

// Teardown by removing the workdir.
func tearDownClean() error {

	return os.RemoveAll(workdir)

}
