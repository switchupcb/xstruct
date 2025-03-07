package config

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/switchupcb/xstruct/cli/models"
)

const (
	recursivePathEnd = "/..."
	recursivePathLen = 4
)

// LoadFiles loads all .go files from the given path.
func LoadFiles(relativepath string) (*models.Generator, error) {
	// determine whether sub-directories will be traversed.
	var recursive bool

	// check for `/...`
	if len(relativepath) >= recursivePathLen {
		end := relativepath[len(relativepath)-recursivePathLen:]
		if end == recursivePathEnd {
			relativepath = relativepath[:len(relativepath)-recursivePathLen]
			recursive = true
		}
	}

	// determine the absolute filepath.
	absfilepath, err := filepath.Abs(relativepath)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while determining the absolute file path of the loader file\n%v", relativepath)
	}

	// walk the entire directory (includes sub-directories).
	gen := new(models.Generator)
	err = filepath.WalkDir(absfilepath, func(path string, d fs.DirEntry, err error) error { //nolint:revive
		if filepath.Ext(path) == ".go" {
			if recursive {
				gen.GoFiles = append(gen.GoFiles, path)
			} else {
				// add the file if the directory == the given directory.
				if filepath.Dir(path) == absfilepath {
					gen.GoFiles = append(gen.GoFiles, path)
				}
			}
		}

		if err != nil {
			return fmt.Errorf("load files: error walking the directory: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("load files: %w", err)
	}

	return gen, nil
}
