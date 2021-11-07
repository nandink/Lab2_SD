package main

import (
	"time"
	pb "Lab2_SD/Pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"math/rand"
)

type Server struct {
	pb.UnimplementedClienteSvServer
}

var montoActual int64
func (s *Server) PedirMonto(ctx context.Context, in *pb.Mensajito2) (*pb.Monto, error) {
	log.Printf("Lider necesita saber cual es el monto acumulado hasta ahora")
	log.Printf("Enviando monto actual...")
	return &pb.Monto{Monto: montoActual}, nil
}

func ServerLider(){
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9001))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterLiderPozoServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func main(){
	//pozo debería utilizar goroutine para tener conexión grpc y rabbitmq al mismo tiempo
	montoActual = 0
	log.Printf("Bienvenido a la consola del Pozo.")
	ServerLider()

}