package main

import "flag"

var grpcAddr = flag.String("grpc", ":3000", "listen address for the grpc transport")
