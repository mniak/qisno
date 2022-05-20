package keepass

func (k *Keepass) Username(path string) (string, error) {
	return k.Attribute(path, "Username")
}

func (k *Keepass) Password(path string) (string, error) {
	return k.Attribute(path, "Password")
}

func (k *Keepass) Attribute(path string, attr string) (string, error) {
	db, err := k.loadDB()
	if err != nil {
		return "", err
	}
	e, err := findEntry(db, path)
	if err != nil {
		return "", err
	}
	return e.GetContent(attr), nil
}
