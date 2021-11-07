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

type Jugada struct {
	jugador [16]string
    ronda int32
    muertos [16]string
}

var ronda int32

func (s *Server) DimeHola(ctx context.Context, in *pb.Mensaje) (*pb.Mensaje, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &pb.Mensaje{Body: "Estás inscrito en el juego!"}, nil
}

func (s *Server) MandarJugadores(ctx context.Context, in *pb.Mensajito2) (*pb.Mensajito2, error) {
	log.Printf("Jugador %d inscrito!", in.Id)
	return &pb.Mensajito2{Id: 1}, nil
}

func (s *Server) MandarJugada(ctx context.Context, in *pb.Jugada) (*pb.Jugada, error) {
	log.Printf("Jugadas Recibidas")
	rand.Seed(time.Now().UnixNano())
	var n_azar int32 = rand.Int31n(4) + 6
	log.Printf("El Lider escogió: %d",n_azar)
	var cont int32 = 0
	ronda = 0
	for {
		if in.Ronda < 5{
			if cont < 16{
				if in.Jugador[cont] >= n_azar{
					 in.Muertos[cont] = 1
					 log.Printf("Jugador %d ha muerto", cont)
				}
			}else{
				break
			}
			
		} else{
			if cont < 16 {
				if in.Jugador[cont] < 21 {
					in.Muertos[cont] = 1
					log.Printf("Jugador %d ha muerto", cont)
				}
			} else{
				break
			}
		}
		cont = cont + 1
	} 
	ronda = ronda + 1
	log.Printf("Devolviendo Jugadas")
	
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.40.171:9050", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := pb.NewLiderNNClient(conn)

	response, err := c.EnviarJugadas(context.Background(), &pb.Jugada{Jugador: in.Jugador, Ronda: in.Ronda, Muertos: in.Muertos})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Respuesta: %d", response.Id)


	var opcioncita int32
	if ronda == 5{  //Termino de ronda, ver si el lider desea ver el pozo (acá también le enviaría los muertos al pozo para que calcule el nuevo monto, pero no lo hemos implementado)
		log.Printf("\n¿Qué acción desea realizar?: ")
		log.Printf("[1] Ver monto acumulado en el pozo  [2] Seguir con la siguiente etapa")
		fmt.Scan(&opcioncita)
		if opcioncita == 1{
			PedirMonto()
		} else{
			continue
		}
	}



	return &pb.Jugada{Jugador: in.Jugador, Ronda: ronda, Muertos: in.Muertos}, nil
	
}

// funcion random rand.Intn(3)
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

func PedirMonto() (error){
	conn, err := grpc.Dial("10.6.40.172:9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewLidezPozoClient(conn)

	response, err := c.PedirMonto(context.Background(), &pb.Mensajito2{Id: 1})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Monto del premio actual: %s", response.Monto)

}


func main() {
	var opcion int32 
	log.Printf("** Bienvenido al Juego del Calamar **")
	ServerJugador()
	log.Printf("Inicio Etapa 1: Luz Verde Luz Roja")


}
