package user

// ToUserDto change user model to userDTO
func ToUserDto(user User) UserDto {
	return UserDto{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Age:      user.Age,
	}
}

// ToUser change userdto model to user
func ToUser(userdto UserDto) User {
	return User{
		Name:     userdto.Name,
		Email:    userdto.Email,
		Age:      userdto.Age,
		Password: userdto.Password,
	}
}

// ToUsers change userdtos model to users array
func ToUsers(userdtos []UserDto) []User {
	users := make([]User, len(userdtos))

	for i, user := range userdtos {
		users[i] = ToUser(user)
	}
	return users
}

// ToUserDtos change users model to userdtos array
func ToUserDtos(users []User) []UserDto {
	userdtos := make([]UserDto, len(users))

	for i, user := range users {
		userdtos[i] = ToUserDto(user)
	}
	return userdtos
}
