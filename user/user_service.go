package user

// UserService for connect repository with api
type UserService struct {
	UserRepository UserRepository
}

func ProvideUserService(u UserRepository) UserService {
	return UserService{UserRepository: u}
}

func (userService *UserService) GetByID(id int) User {
	var user = userService.UserRepository.FindByID(id)
	return user
}

func (userService *UserService) GetAll() []User {
	var users = userService.UserRepository.FindAll()
	return users
}

func (userService *UserService) Delete(id int) bool {
	return userService.UserRepository.Delete(id)
}

func (userService *UserService) Create(user User) (User, error) {
	return userService.UserRepository.Create(user)
}
