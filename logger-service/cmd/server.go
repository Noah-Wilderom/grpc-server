package main

import (
	"fmt"
	pb "github.com/Noah-Wilderom/grpc-server/shared/logs"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
)

type Server struct {
	port int
}

func NewServer() *Server {
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		panic(err)
	}

	return &Server{
		port: port,
	}
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalln("Failed to listen for gRPC:", err)
	}

	grpc := grpc.NewServer()

	pb.Register(grpc, &QueueListenerServer{})
	log.Println("gRPC Server started on port", gRPCPort)

	if err = grpc.Serve(lis); err != nil {
		log.Fatalln("Failed to listen for gRPC:", err)
	}
}
