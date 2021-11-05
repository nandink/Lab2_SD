package main

import (
	"log"
)

type Server struct {
}

func (s *Server) DimeHola(ctx context.Context, in *Mensaje) (*Mensaje, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Mensaje{Body: "Hello From the Server!"}, nil
}
func main() {
	log.Printf("Holi estoy funcionando")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterClienteSvServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
