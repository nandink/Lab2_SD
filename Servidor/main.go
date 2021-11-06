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

type Jugada struct {
	jugador [16]string
    ronda int32
    muertos [16]string
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

func ListenNN(){
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.40.171:9050", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
}

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

	log.Printf("Inicio Etapa 1: Luz Verde Luz Roja")

}
