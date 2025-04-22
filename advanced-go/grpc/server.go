package main

import (
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
)

func serve(network, host, port string) {
	address := strings.Join([]string{host, port}, ":")
	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatal("Failed to listen on: ", address, err)
		return
	}

	grpcServer := grpc.NewServer()
	log.Println("Grpc Server listening on", address)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Failed to serve on: ", address, err)
		return
	}
	
}

func main() {
	network, host, port := "tcp", "localhost", "50051"
	serve(network, host, port)
}