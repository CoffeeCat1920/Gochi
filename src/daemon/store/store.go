package store

import (
	"daemon/objects"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

type AppStore struct {
	names   []string
	nameMap map[string]string
}

func (appStore *AppStore) Names() []string { return appStore.names }

var (
	appStoreInstance *AppStore
)

func New() *AppStore {
	if appStoreInstance != nil {
		return appStoreInstance
	}

	nameMap := make(map[string]string)
	var names []string

	fs := afero.NewOsFs()
	path := "/usr/share/applications"

	files, err := afero.ReadDir(fs, path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".desktop" {
			entry, err := objects.NewAppEntry(filepath.Join(path, file.Name()))
			if err != nil {
				continue
			}
			names = append(names, entry.Name())
			nameMap[entry.Name()] = entry.Exec()
		}
	}

	return &AppStore{
		names:   names,
		nameMap: nameMap,
	}
}
