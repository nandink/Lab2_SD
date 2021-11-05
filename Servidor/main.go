package main

import (
	"fmt"
	"net"
	"context"
	"google.golang.org/grpc"
	"log"
	"google.golang.org/grpc"
	pb "Lab2_SD/Pb"	
)

type Server struct {
	pb.UnimplementedClienteSvServer
}

func (s *Server) DimeHola(ctx context.Context, in *pb.Mensaje) (*pb.Mensaje, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Mensaje{Body: "Hello From the Server!"}, nil
}
func main() {
	log.Printf("Holi estoy funcionando")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	s.RegisterClienteSvServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
	
