package main

import (
	"context"
	"fmt"

	"github.com/ayadnez/proto"
	"github.com/ayadnez/types"
)

var users = map[int32]types.User{
	1: {ID: 1, FName: "zahid", City: "SGR", Phone: "1234567890", Height: 5.9, Married: true},
	2: {ID: 2, FName: "John", City: "NY", Phone: "2345678901", Height: 5.9, Married: false},
	3: {ID: 3, FName: "ahmed", City: "Delhi", Phone: "234567345", Height: 5.7, Married: false},
	4: {ID: 4, FName: "George", City: "Srinagar", Phone: "234598777", Height: 5.9, Married: true},
	5: {ID: 5, FName: "Gautham", City: "Chennai", Phone: "2345678901", Height: 5.8, Married: false},
}

type Server struct {
	proto.UnimplementedUserServiceServer
}

func (s *Server) GetUser(ctx context.Context, in *proto.UserRequestId) (*proto.User, error) {

	user, exits := users[in.Id]

	if !exits {
		return nil, fmt.Errorf("the user  with id : &d is not found", in.Id)
		//return nil, grpc.Errorf(grpc.Code(codes.NotFound), "user not found")
	}

	return &proto.User{
		Id:      user.ID,
		Fname:   user.FName,
		City:    user.City,
		Phone:   user.Phone,
		Height:  user.Height,
		Married: user.Married,
	}, nil
}

func (s *Server) GetUsers(ctx context.Context, in *proto.UserRequestIds) (*proto.UserResponses, error) {
	var result []*proto.User

	for _, id := range in.Ids {
		if user, exits := users[id]; exits {
			result = append(result, &proto.User{
				Id:      user.ID,
				Fname:   user.FName,
				City:    user.City,
				Phone:   user.Phone,
				Height:  user.Height,
				Married: user.Married,
			})
		}
	}

	return &proto.UserResponses{Users: result}, nil

}

func (s *Server) SearchUsers(ctx context.Context, in *proto.SearchCriteria) (*proto.UserResponses, error) {
	var result []*proto.User

	for _, user := range users {
		if user.City == in.City && user.Married == in.Married {
			result = append(result, &proto.User{
				Id:      user.ID,
				Fname:   user.FName,
				City:    user.City,
				Phone:   user.Phone,
				Height:  user.Height,
				Married: user.Married,
			})
		}
	}

	return &proto.UserResponses{Users: result}, nil

}
