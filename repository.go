package basego




type Repository interface {
	CreateNewUser(user *User) error
	GetUserByID(id string) (*User, error)
	AllUsers() ([]*User, error)
	UpdateUserByID(user *User) error
	DeleteUserByID(id string) error
}