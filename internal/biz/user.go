package biz

import (
	"context"
	"errors"
	"kratos-realworld/internal/conf"
	"kratos-realworld/internal/pkg/middleware/auth"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email        string
	UserName     string
	Token        string
	Bio          string
	Image        string
	Password     string
	PasswordHash string
}

type UserLogin struct {
	Email    string
	UserName string
	Token    string
	Bio      string
	Image    string
	Password string
}

func hashPassword(password string) string {
	if bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return ""
	} else {
		return string(bytes)
	}
}

func verifyPassword(password, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash)); err != nil {
		return false
	}
	return true
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type ProfileRepo interface {
}

type UesrUsecase struct {
	ur   UserRepo
	pr   ProfileRepo
	jwtc *conf.JWT
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewUesrUsecase(ur UserRepo,
	pr ProfileRepo, jwtc *conf.JWT, logger log.Logger) *UesrUsecase {
	return &UesrUsecase{
		ur:   ur,
		pr:   pr,
		jwtc: jwtc,
		log:  log.NewHelper(logger),
	}
}

func (uc *UesrUsecase) generateToken(username string) string {
	return auth.GenerateToke(uc.jwtc.Token, username)
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *UesrUsecase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	u := &User{
		Email:    email,
		UserName: username,
		Password: password,
	}
	if err := uc.ur.CreateUser(ctx, u); err != nil {
		return nil, err
	}
	return &UserLogin{
		Email:    u.Email,
		UserName: u.UserName,
		Bio:      u.Bio,
		Image:    u.Image,
		Token:    uc.generateToken(username),
	}, nil
}

func (uc *UesrUsecase) Login(ctx context.Context, email, password string) (rely *UserLogin, err error) {
	u, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !verifyPassword(password, u.PasswordHash) {
		return nil, errors.New("login error")
	}
	return &UserLogin{
		Email:    email,
		UserName: u.UserName,
		Token:    uc.generateToken(u.UserName),
	}, nil

}
