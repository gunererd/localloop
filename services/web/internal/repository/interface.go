package repository

// UserRepository defines the interface for user-related operations
type UserRepository interface {
	Login(email, password string) (string, error)
	Register(email, password, name string) error
}
