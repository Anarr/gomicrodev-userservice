package userservice_grpc

import (
	"context"
	u "github.com/Anarr/gomicrodev-userservice/proto/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserService struct{}

var users []*u.User

func(us *UserService) All(ctx context.Context, req *u.AllRequest) (*u.AllResponse, error) {
	return nil, nil
}

func (us *UserService) Get(ctx context.Context, req *u.GetUserRequest) (*u.User, error) {
	return nil, nil
}

func (us *UserService) Upload(stream u.UserService_UploadServer) error {
	return nil
}
