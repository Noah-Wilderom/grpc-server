package main

import (
	"context"
	"fmt"
	"github.com/Noah-Wilderom/grpc-server/logger-service/database/models"
	"github.com/Noah-Wilderom/grpc-server/logger-service/handlers"
	pb "github.com/Noah-Wilderom/grpc-server/shared/logs"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
)

type Server struct {
	pb.UnimplementedLogServiceServer
	port    int
	handler *handlers.LogHandler
}

func NewServer(handler *handlers.LogHandler) *Server {
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		panic(err)
	}

	return &Server{
		port:    port,
		handler: handler,
	}
}

func (s *Server) WriteLog(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error) {
	data := req.GetLogEntry()
	logEntry := &models.Log{
		Type: data.Type,
		Data: data.Data,
	}

	err := s.handler.SaveLog(logEntry)
	if err != nil {
		return nil, err
	}

	res := &pb.LogResponse{
		Status: 200,
	}

	return res, nil
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalln("Failed to listen for gRPC:", err)
		return err
	}

	g := grpc.NewServer()

	pb.RegisterLogServiceServer(g, s)
	log.Println("gRPC Server started on port", s.port)

	if err = g.Serve(lis); err != nil {
		log.Fatalln("Failed to listen for gRPC:", err)
		return err
	}

	return nil
}
