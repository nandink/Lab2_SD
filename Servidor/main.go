package main

import (
	"fmt"
	"net"
	"context"
<<<<<<< HEAD
=======
	"log"
>>>>>>> 753bccc4a913771db1b8b0164cea6c5977691319
	"google.golang.org/grpc"
	"log"
	pb "Lab2_SD/Pb"	
)

type Server struct {
	pb.UnimplementedClienteSvServer
}

func (s *Server) DimeHola(ctx context.Context, in *pb.Mensaje) (*pb.Mensaje, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &pb.Mensaje{Body: "Hello From the Server!"}, nil
}
func main() {
	log.Printf("Holi estoy funcionando")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterClienteSvServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
	
