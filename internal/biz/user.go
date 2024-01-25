package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	UserName string
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
}

type ProfileRepo interface {
}

type UesrUsecase struct {
	ur  UserRepo
	pr  ProfileRepo
	log *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewUesrUsecase(ur UserRepo,
	pr ProfileRepo, logger log.Logger) *UesrUsecase {
	return &UesrUsecase{
		ur:  ur,
		pr:  pr,
		log: log.NewHelper(logger),
	}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *UesrUsecase) Register(ctx context.Context, u *User) error {
	uc.ur.CreateUser(ctx, u)
	return nil
}
