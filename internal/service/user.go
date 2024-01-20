package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (rely *v1.UserReply, err error) {

	return &v1.UserReply{
		User: &v1.UserReply_User{
			Email:    "email",
			Token:    "token",
			Username: "username",
			Bio:      "bio",
			Image:    "image",
		},
	}, nil
}