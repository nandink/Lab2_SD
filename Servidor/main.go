package main

import (
<<<<<<< HEAD
	"google.golang.org/grpc"
=======
	"fmt"
	"net"
	"context"
	"io"
>>>>>>> 8954c044745fbe94ea168b0724fc373d2842a5d8
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

	s.RegisterClienteSvServer(s, &Server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
	
