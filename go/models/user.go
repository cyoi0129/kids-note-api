package models

type User struct {
	Id       uint
	Email    string
	Password string
	Name     string
	Gender   string
	Family   int
}

type NewUser struct {
	Id       uint
	Email    string
	Password string
	Name     string
	Gender   string
	Token    string
	Family   int
}

type UserResponse struct {
	Info  User
	Token string
}

type LoginUser struct {
	Email    string
	Password string
}

type TokenUser struct {
	Id    uint
	Token string
}
