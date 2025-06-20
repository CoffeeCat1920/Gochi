package entities

type AppStore struct {
	Entries    []AppEntry
	NameMap    map[string]string   // Name to Exec map
	CateoryMap map[string][]string // Category to name map
}

func NewAppStore(entries []AppEntry) *AppStore {
	nameMap := make(map[string]string)
	categoryMap := make(map[string][]string)

	for _, entry := range entries {
		nameMap[entry.Name] = entry.Exec
		for _, category := range entry.Categories {
			categoryMap[category] = append(categoryMap[category], entry.Name)
		}
	}

	return &AppStore{
		Entries:    entries,
		NameMap:    nameMap,
		CateoryMap: categoryMap,
	}
}

func (store AppStore) Search(term string) {

}
