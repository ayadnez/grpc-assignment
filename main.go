package main

import (
	"flag"
	"log"
	"net"

	"github.com/ayadnez/proto"
	"google.golang.org/grpc"
)

var grpcAddr = flag.String("grpc", ":3000", "listen address for the grpc transport")

func main() {
	// parse the flag
	flag.Parse()

	listenAddr, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, &Server{})

	log.Printf("Server listening at %v", listenAddr.Addr())

	if err := s.Serve(listenAddr); err != nil {
		log.Printf("failed to serve: %v", err)
	}

}
