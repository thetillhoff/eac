package apps

import (
	"log"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/thetillhoff/eac/internal/config"
	"github.com/thetillhoff/eac/pkg/logs"
)

// Returns a list of all apps that are managed via `eac` (== contained in `versions.yaml`).
// Per default one app per line in the format `<app>[==<installedVersion>]`.
// The seperator can be edited with a flag and is returned as the second parameter.
// Created by trying to run the getInstalledVersion script.
func listRemote(conf config.Config) []string {
	logs.Verbose = conf.Verbose

	// clones the repository into memory - no persistency involved!
	// Sadly, sparse checkout is not supported yet with this library - https://github.com/go-git/go-git/issues/90
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL:   "https://github.com/thetillhoff/eac",
		Depth: 1,
		Tags:  git.NoTags,
	})
	if err != nil {
		log.Fatalln(err)
	}

	ref, err := r.Head()
	if err != nil {
		log.Fatalln(err)
	}

	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		log.Fatalln(err)
	}

	tree, err := commit.Tree()
	if err != nil {
		log.Fatalln(err)
	}

	uniqueAppNames := map[string]bool{}
	tree.Files().ForEach(func(f *object.File) error {
		if strings.HasPrefix(f.Name, "apps/") {
			potentialAppName := strings.TrimPrefix(f.Name, "apps/")
			potentialAppName = strings.Split(potentialAppName, "/")[0]
			uniqueAppNames[potentialAppName] = true
		}
		return nil
	})

	appNames := []string{}
	for key := range uniqueAppNames {
		appNames = append(appNames, key)
	}

	return appNames
}
