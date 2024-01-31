package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/errors"
)

func (s *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (rely *v1.UserReply, err error) {
	if len(req.User.Email) == 0 {
		return nil, errors.NewHTTPError(400, "email", "can't be empty")
	}
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

func (s *RealWorldService) Register(ctx context.Context, req *v1.RegisterRequest) (rely *v1.UserReply, err error) {
	u, err := s.uc.Register(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Email:    u.Email,
			Username: u.UserName,
			Token:    u.Token,
		},
	}, nil
}

func (s *RealWorldService) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserRequest) (rely *v1.UserReply, err error) {
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
