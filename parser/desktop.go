package parser

import (
	"errors"

	"github.com/go-ini/ini"
)

type AppEntry struct {
	Name string
	Exec string
	Icon string
}

func ParseDesktopFile(path string) (*AppEntry, error) {
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
