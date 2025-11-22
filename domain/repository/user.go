package repository

type UserRepository interface {
	Create(username, email, passwordHash string) error
	ExistsByUsername(username string) (bool, error)
}
