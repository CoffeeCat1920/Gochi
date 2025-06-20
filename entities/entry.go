package entities

import (
	"errors"
	"strings"

	"github.com/go-ini/ini"
)

type AppEntry struct {
	Name       string
	Categories []string
	Exec       string
}

func NewAppEntry(path string) (*AppEntry, error) {
	cfg, err := ini.Load(path)
	if err != nil {
		return nil, err
	}

	section := cfg.Section("Desktop Entry")

	name := section.Key("Name").String()
	exec := section.Key("Exec").String()
	categories := strings.Split(section.Key("Categories").String(), ";")

	categories = append(categories, name)

	if name == "" || exec == "" {
		return nil, errors.New("invalid .desktop entry: missing name or exec")
	}

	return &AppEntry{
		Name:       name,
		Categories: categories,
		Exec:       exec,
	}, nil
}
