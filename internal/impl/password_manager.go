package impl

import (
	"fmt"
	"os"
	"strings"

	"github.com/tobischo/gokeepasslib/v3"
)

type KeepassPasswordManager struct {
	Filename     string
	FilePassword string
}

func loadDB(filename, password string) (*gokeepasslib.Database, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	db := gokeepasslib.NewDatabase()
	db.Credentials = gokeepasslib.NewPasswordCredentials(password)
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
		if strings.EqualFold(e.GetTitle(), name) {
			return e, nil
		}
	}
	return gokeepasslib.Entry{}, fmt.Errorf("entry %s not found", name)
}

func findEntry(db *gokeepasslib.Database, path string) (gokeepasslib.Entry, error) {
	segments := strings.Split(path, "/")
	var group *gokeepasslib.Group
	for len(segments) >= 2 {
		var groups []gokeepasslib.Group
		name := segments[0]
		segments = segments[1:]
		if group == nil {
			groups = db.Content.Root.Groups
		} else {
			groups = group.Groups
		}
		g, err := findGroupOnSlice(groups, name)
		if err != nil {
			return gokeepasslib.Entry{}, err
		}
		group = &g
	}
	if group != nil {
		return findEntryOnSlice(group.Entries, segments[0])
	}
	return gokeepasslib.Entry{}, fmt.Errorf("path %s is invalid", path)
}

func (k *KeepassPasswordManager) Username(path string) (string, error) {
	return k.Attribute(path, "Username")
}

func (k *KeepassPasswordManager) Password(path string) (string, error) {
	return k.Attribute(path, "Password")
}

func (k *KeepassPasswordManager) Attribute(path, attr string) (string, error) {
	db, err := loadDB(k.Filename, k.FilePassword)
	if err != nil {
		return "", err
	}
	e, err := findEntry(db, path)
	if err != nil {
		return "", err
	}
	return e.GetContent(attr), nil
}
