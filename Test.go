package main

type Database interface {
	GetUser(id int) (string, error)
}
type UserService struct {
	db Database
}

func NewUserService(db Database) *UserService {
	return &UserService{db: db}
}
func (us *UserService) GetUserByID(id int) (string, error) {
	return us.db.GetUser(id)
}
