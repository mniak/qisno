package abstractions

type PasswordManager interface {
	Username(path string) (string, error)
	Password(path string) (string, error)
	Attribute(path, attr string) (string, error)
}
