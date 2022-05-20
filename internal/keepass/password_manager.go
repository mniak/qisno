package keepass

func (k *Keepass) Username(path string) (string, error) {
	result, err := k.Attribute(path, "UserName")
	return result, err
}

func (k *Keepass) Password(path string) (string, error) {
	result, err := k.Attribute(path, "Password")
	return result, err
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
	result := e.GetContent(attr)
	return result, nil
}
