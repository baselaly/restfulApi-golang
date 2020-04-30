package user

type UserDto struct {
	ID       uint   `json:"id,string"`
	Name     string `json:"name,string"`
	Age      int    `json:"age,int"`
	Password string `json:"password,string"`
	Email    string `json:"email,string"`
}
