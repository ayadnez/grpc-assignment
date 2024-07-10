# gRPC User Service

This project is a Golang gRPC service that manages user details and includes a search capability. It simulates a database using an in-memory map and provides gRPC endpoints for fetching user details based on user IDs and search criteria.

## Project Structure

- `main.go`: Starts the gRPC server.
- `server.go`: Contains the gRPC server implementation.
- `types.go`: Contains the user type and simulated database.
- `client/client.go`: Implements a client to test the gRPC server.
- `proto/service.proto`: Protobuf file defining the gRPC service and messages.

## User Model

```json
{
    "id": 1,
    "fname": "Steve",
    "city": "LA",
    "phone": 1234567890,
    "height": 5.8,
    "married": true
}

### gRPC Endpoints

   - GetUser: Fetches a user by ID.
   - GetUsers: Fetches multiple users by their IDs.
 - SearchUsers: Searches users based on city and marital status.


Prerequisites

    Go 1.16+
    Protobuf compiler (protoc)
    gRPC Go plugin



Install gRPC and Protobuf Go Plugins:

go get google.golang.org/grpc
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc


Running the Grpc Server : 
  command : go run main.go server.go 

Running the Grpc Client : 

   command : go run client.go
