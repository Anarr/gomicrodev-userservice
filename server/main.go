package server

import (
	u "github.com/Anarr/gomicrodev-userservice/proto/user"
	s "github.com/Anarr/gomicrodev-userservice"
	"log"
	"net"
	"google.golang.org/grpc"
)

var users []*u.User

func init() {
	u := []*u.User{
		&u.User{
			Id:       1,
			Name:     "Anar Rzayev",
			Username: "anar.rzayev",
		},
		&u.User{
			Id:       2,
			Name:     "Elmir Mahmudov",
			Username: "elmirm",
		},
		&u.User{
			Id:       3,
			Name:     "Elshad Yarmetov",
			Username: "elsh",
		},
	}

	users = append(users, u...)
}

func Serve() error {
	list, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Println("connection err:", err)
		return err
	}

	server := grpc.NewServer()
	u.RegisterUserServiceServer(server, new(s.UserService))

	if err = server.Serve(list); err != nil {
		log.Println("server err: ", err)
		return err
	}

	return nil
}
