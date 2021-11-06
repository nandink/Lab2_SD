package main

import (
	pb "Lab2_SD/Pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.UnimplementedClienteSvServer
}

func (s *Server) DimeHola(ctx context.Context, in *pb.Mensaje) (*pb.Mensaje, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &pb.Mensaje{Body: "Estás inscrito en el juego!"}, nil
}

func (s *Server) MandarJugadores(ctx context.Context, in *pb.Mensajito2) (*pb.Mensajito2, error) {
	log.Printf("Jugador %d inscrito!", in.Id)
	return &pb.Mensajito2{Id: 1}, nil
}

//funcion conectar_jugador

func ServerJugador(){
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterClienteSvServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	pb.RegisterClienteSvServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}


func main() {
	log.Printf("** Bienvenido al Juego del Calamar **")
	ServerJugador()
}
