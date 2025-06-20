package loader

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

type AppEntry struct {
	Name       string
	Categories []string
	Exec       string
	Icon       string
}

func GetEntries() ([]AppEntry, error) {
	var entries []AppEntry
	fs := afero.NewOsFs()
	path := "/usr/share/applications"

	files, err := afero.ReadDir(fs, path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) == ".desktop" {
			entry, err := parseDesktopFile(filepath.Join(path, f.Name()))
			if err != nil {
				continue
			}
			entries = append(entries, *entry)
		}
	}

	return entries, nil
}
