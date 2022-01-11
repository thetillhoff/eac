package apps

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/thetillhoff/eac/internal/config"
)

// List the (remotely) available apps
func PrintList(conf config.Config) {

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

	for uniqueAppName := range uniqueAppNames {
		fmt.Println(uniqueAppName)
	}
}
