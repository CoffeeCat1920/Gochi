package loader

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/CoffeeCat1920/Gochi/entities"
	"github.com/spf13/afero"
)

func GetEntries() ([]entities.AppEntry, error) {
	var entries []entities.AppEntry
	fs := afero.NewOsFs()
	path := "/usr/share/applications"

	files, err := afero.ReadDir(fs, path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) == ".desktop" {
			entry, err := entities.NewAppEntry(filepath.Join(path, f.Name()))
			if err != nil {
				continue
			}
			entries = append(entries, *entry)
		}
	}

	return entries, nil
}
