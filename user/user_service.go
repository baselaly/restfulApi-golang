package user

// UserService for connect repository with api
type UserService struct {
	UserRepository UserRepository
}

func ProvideUserService(u UserRepository) UserService {
	return UserService{UserRepository: u}
}

func (userService *UserService) GetByID(id int) (User, error) {
	return userService.UserRepository.FindByID(id)
}

func (userService *UserService) GetAll() []User {
	return userService.UserRepository.FindAll()
}

func (userService *UserService) Delete(id int) (bool, error) {
	return userService.UserRepository.Delete(id)
}

func (userService *UserService) Create(user User) (User, error) {
	return userService.UserRepository.Create(user)
}
