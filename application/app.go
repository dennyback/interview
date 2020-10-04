package application

import (
	"context"
	"github.com/dennyback/interview/contract"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Application struct {
	Ctx context.Context
}

type Service struct {
	Name string
}

func (s Service) GetUser(ctx context.Context, request *contract.GetUserRequest) (*contract.GetUserResponse, error) {
	db := NewDB()
	user := db.FindByID(request.User.Id, ctx)

	return &contract.GetUserResponse{User: &contract.GetUserResponse_User{Id: uint32(user.ID), FirstName: user.FirstName, LastName: user.LastName}}, nil
}

func (s Service) AddUser(ctx context.Context, request *contract.AddUserRequest) (*contract.AddUserResponse, error) {
	db := NewDB()
	user := db.CreateNewUser(request.User.FirstName, request.User.LastName)

	return &contract.AddUserResponse{User: &contract.AddUserResponse_User{Id: uint32(user.ID), FirstName: user.FirstName, LastName: user.LastName}}, nil
}

func (app *Application) RunAndServeApplication(name string) error {
	service := &Service{name}
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	contract.RegisterUsersServer(grpcServer, service)

	return grpcServer.Serve(lis)
}
