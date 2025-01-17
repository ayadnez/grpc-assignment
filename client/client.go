package main

import (
	"context"
	"log"
	"os"

	"github.com/ayadnez/proto"
	"google.golang.org/grpc"
)

func main() {

	// connect to the grpc server

	serverAddr := os.Getenv("GRPC_SERVER_ADDRESS")
	if serverAddr == "" {
		serverAddr = "grpc-server:3000" // Default to localhost
	}

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())

	if err != nil {
		log.Fatal("failed to connect : %v", err)
	}

	defer conn.Close()

	client := proto.NewUserServiceClient(conn)

	ctx := context.Background()

	// test GetUser

	testGetUser(ctx, client, 2)

	// test GetUsers

	testGetUsers(ctx, client, []int32{1, 2})

	// test SearchUsers criteria

	testSearchUsers(ctx, client, "SGR", true)
}

func testGetUser(ctx context.Context, client proto.UserServiceClient, userId int32) {

	user, err := client.GetUser(ctx, &proto.UserRequestId{Id: userId})

	if err != nil {
		log.Fatal("could not get the User : %v", err.Error())
	}

	log.Printf("User is : %v", user)
}

func testGetUsers(ctx context.Context, client proto.UserServiceClient, userIds []int32) {

	users, err := client.GetUsers(ctx, &proto.UserRequestIds{Ids: userIds})

	if err != nil {
		log.Fatal("could not find the users : %v", err)
	}

	log.Printf("users : ", users)
}

func testSearchUsers(ctx context.Context, client proto.UserServiceClient, city string, married bool) {

	users, err := client.SearchUsers(ctx, &proto.SearchCriteria{City: city, Married: married})

	if err != nil {
		log.Fatal("could not search the users : %v", err)
	}

	log.Printf("users based on the criteria are : %v", users)
}
