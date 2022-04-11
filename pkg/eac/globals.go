package eac

import (
	"os"
	"path"
)

var (
	workdir      = path.Join(os.TempDir() + "/eac")
	Verbose bool = false
	DryRun  bool = false // Only used by InstallApp and Clean
)
