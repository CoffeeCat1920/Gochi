package loader

import (
	"errors"

	"github.com/go-ini/ini"
)

func parseDesktopFile(path string) (*AppEntry, error) {
	cfg, err := ini.Load(path)
	if err != nil {
		return nil, err
	}

	section := cfg.Section("Desktop Entry")

	name := section.Key("Name").String()
	exec := section.Key("Exec").String()
	icon := section.Key("Icon").String()

	if name == "" || exec == "" {
		return nil, errors.New("invalid .desktop entry: missing name or exec")
	}

	return &AppEntry{
		Name: name,
		Exec: exec,
		Icon: icon,
	}, nil
}
