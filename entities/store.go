package entities

import "github.com/sahilm/fuzzy"

type AppStore struct {
	Names      []string
	Categories []string
	NameMap    map[string]string   // Name to Exec map
	CateoryMap map[string][]string // Category to name map
}

func NewAppStore(entries []AppEntry) *AppStore {
	nameMap := make(map[string]string)
	categoryMap := make(map[string][]string)
	var names []string
	var categories []string

	for _, entry := range entries {
		names = append(names, entry.Name)
		nameMap[entry.Name] = entry.Exec
		for _, category := range entry.Categories {
			categoryMap[category] = append(categoryMap[category], entry.Name)
		}
	}

	for category := range categoryMap {
		categories = append(categories, category)
	}

	return &AppStore{
		Names:      names,
		Categories: categories,
		NameMap:    nameMap,
		CateoryMap: categoryMap,
	}
}

func (store AppStore) Search(term string) []string {
	seen := make(map[string]struct{})
	var results []string

	nameMatches := fuzzy.Find(term, store.Names)
	for _, match := range nameMatches {
		if _, ok := seen[match.Str]; !ok {
			seen[match.Str] = struct{}{}
			results = append(results, match.Str)
		}
	}

	// Match categories → get apps from those categories
	categoryMatches := fuzzy.Find(term, store.Categories)
	for _, match := range categoryMatches {
		apps := store.CateoryMap[match.Str]
		for _, name := range apps {
			if _, ok := seen[name]; !ok {
				seen[name] = struct{}{}
				results = append(results, name)
			}
		}
	}

	return results
}
