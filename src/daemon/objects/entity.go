package objects

import (
	"errors"
	"strings"

	"github.com/go-ini/ini"
)

type AppEntry struct {
	name       string
	exec       string
	categories []string
}

func (app *AppEntry) Name() string         { return app.name }
func (app *AppEntry) Exec() string         { return app.exec }
func (app *AppEntry) Categories() []string { return app.categories }

func NewAppEntry(path string) (*AppEntry, error) {
	cfg, err := ini.Load(path)
	if err != nil {
		return nil, err
	}

	section := cfg.Section("Desktop Entry")

	name := section.Key("Name").String()
	exec := section.Key("Exec").String()
	categories := strings.Split(section.Key("Categories").String(), ";")

	if name == "" || exec == "" {
		return nil, errors.New("invalid entry")
	}

	return &AppEntry{
		name:       name,
		exec:       exec,
		categories: categories,
	}, err
}
