package di

import "go/api-demo/internal/user"

type IStatRepository interface {
	AddClick(linkId uint)
}


type IUserRepository interface {
	Create(user *user.User) (*user.User, error)
	FindByEmail(email string) (*user.User, error)
}
