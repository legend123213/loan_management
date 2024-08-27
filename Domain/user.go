package domain


type User struct{
	ID string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	IsAdmin bool `json:"is_admin" bson:"is_admin"`
	IsActive bool `json:"is_active" bson:"is_active"`
	
}

type UserRepo interface{
	CreateUser(user *User) (*User, error)
	GetUser(id string) (*User, error)	
	GetUsers() ([]*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(id string) error
}

type UserUsecase interface{
	AddeUser(user *User) (*User, error)
	GetUser(id string) (*User, error)
	GetUsers() ([]*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(id string) error
}