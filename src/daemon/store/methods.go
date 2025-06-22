package store

import (
	"fmt"
	"os/exec"
	"strings"
)

func (store *AppStore) Run(name string) error {
	exe, ok := store.nameMap[name]
	if !ok {
		return fmt.Errorf("can't find the app of name %v", name)
	}

	cmdStr := strings.ReplaceAll(exe, "%u", "")
	parts := strings.Fields(cmdStr)
	cmd := exec.Command(parts[0], parts[1:]...)

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("error launching app: %v", err)
	}

	return nil
}
