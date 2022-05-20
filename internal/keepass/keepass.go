package keepass

import (
	"fmt"
	"os"
	"strings"

	"github.com/tobischo/gokeepasslib/v3"
)

type Keepass struct {
	config Config
}

func (k *Keepass) loadDB() (*gokeepasslib.Database, error) {
	file, err := os.Open(k.config.Database)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	db := gokeepasslib.NewDatabase()
	db.Credentials = gokeepasslib.NewPasswordCredentials(k.config.Password)
	err = gokeepasslib.NewDecoder(file).Decode(db)
	if err != nil {
		return nil, err
	}
	err = db.UnlockProtectedEntries()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func findGroupOnSlice(groups []gokeepasslib.Group, name string) (gokeepasslib.Group, error) {
	for _, g := range groups {
		if strings.EqualFold(g.Name, name) {
			return g, nil
		}
	}
	return gokeepasslib.Group{}, fmt.Errorf("group %s not found", name)
}

func findEntryOnSlice(entries []gokeepasslib.Entry, name string) (gokeepasslib.Entry, error) {
	for _, e := range entries {
		entryTitle := e.GetTitle()
		if strings.EqualFold(entryTitle, name) {
			return e, nil
		}
	}
	return gokeepasslib.Entry{}, fmt.Errorf("entry %s not found", name)
}

func findEntry(db *gokeepasslib.Database, path string) (gokeepasslib.Entry, error) {
	segments := strings.Split(path, "/")
	group := db.Content.Root.Groups[0]
	for len(segments) >= 2 {
		var groups []gokeepasslib.Group
		name := segments[0]
		segments = segments[1:]
		groups = group.Groups

		g, err := findGroupOnSlice(groups, name)
		if err != nil {
			return gokeepasslib.Entry{}, err
		}
		group = g
	}
	return findEntryOnSlice(group.Entries, segments[0])

	return gokeepasslib.Entry{}, fmt.Errorf("path %s is invalid", path)
}
