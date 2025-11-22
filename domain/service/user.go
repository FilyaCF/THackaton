package service

type RegisterRequest struct {
	Username string
	Email    string
	Password string
}

type UserService interface {
	Register(req RegisterRequest) error
}
