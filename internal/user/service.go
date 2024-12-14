package user

type UserService struct {
	UserRepo *Repository
}

func NewUserService(userRepo *Repository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (us *UserService) CreateUser(user *User) (*User, error) {
	panic("implement me")
}
