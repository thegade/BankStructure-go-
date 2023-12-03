package domain

import "github.com/gofrs/uuid"

type User struct {
	id       uuid.UUID
	login    string
	password string
	fullname string
}

func (u *User) ID() uuid.UUID    { return u.id }
func (u *User) Login() string    { return u.login }
func (u *User) Password() string { return u.password }
func (u *User) Fullname() string { return u.fullname }

func NewUser(id uuid.UUID, login string, password string, fullname string) *User {
	return &User{
		id:       id,
		login:    login,
		password: password,
		fullname: fullname,
	}
}

type UserRepository interface {
	Save(u *User) (uuid.UUID, error)
	FindUser(login string, password string) (User, error)
}
