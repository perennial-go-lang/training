package user

type User struct {
	Id int `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Age int `json:"age"`
}

type Users []User

func (u *User) getName() string {
	return u.FirstName + " " +u.LastName
}


