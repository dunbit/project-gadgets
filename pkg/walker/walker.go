package walker

import (
	"os"
	"path/filepath"

	"github.com/bmatcuk/doublestar"
	"github.com/dunbit/project-gadgets/pkg/config"
	"github.com/dunbit/project-gadgets/pkg/file"
)

// Walk ...
func Walk(c *config.Config, root string) ([]*file.File, error) {
	files := []*file.File{}

	_, err := os.Stat(root)
	if os.IsNotExist(err) {
		return files, os.ErrNotExist
	}

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			for _, e := range c.Excludes {
				match, err := doublestar.PathMatch(filepath.FromSlash(e), path)
				if err != nil {
					return err
				}
				if match {
					return filepath.SkipDir
				}
			}
		}

		for _, m := range c.Files {
			match, err := doublestar.PathMatch(filepath.FromSlash(m.Match), path)
			if err != nil {
				return err
			}
			if match {
				files = append(files, file.New(path, m.Comment))
			}
		}

		return err
	})
	return files, err
}
