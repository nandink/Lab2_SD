package main

import (
	pb "Lab2_SD/Pb"
	"context"
	"google.golang.org/grpc"
	"log"
	"fmt"
  	"os"
	"net"
	"math/rand"
	"time"
)

//var path = "/home/sofia/Escritorio/Tareas/SD_Final/Lab2_SD/NameNode/archivo.txt"
var path = "/home/alumno/Lab2_SD/NameNode/Registro.txt"

type Server struct {
	pb.UnimplementedLiderNNServer
}

type Jugada struct {
	jugador [16]string
    ronda int32
    muertos [16]string
}

func conexionNameLider(){
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9050)) //bajar el firewall para 9050
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil{
		log.Fatalf("No se puede crear el servidor: %v", err)
	}

	pb.RegisterLiderNNServer(s, &Server{})
}

func crearArchivo() {
	//Verifica que el archivo existe
	var _, err = os.Stat(path)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
	  var file, err = os.Create(path)
	  if existeError(err) {
		return
	  }
	  defer file.Close()
	}
	fmt.Println("File Created Successfully", path)
  }

  func escribeArchivo(linea string) {
	// Abre archivo usando permisos READ & WRITE
	var file, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if existeError(err) {
	  return
	}
	defer file.Close()

	// Escribe algo de texto linea por linea
	_, err = file.WriteString(linea)
	if existeError(err) {
	  return
	}

	// Salva los cambios
	err = file.Sync()
	if existeError(err) {
		return
	  }
	  fmt.Println("Archivo actualizado existosamente.")
	}

	func existeError(err error) bool {
	  if err != nil {
		fmt.Println(err.Error())
	  }
	  return (err != nil)
	}

	func (s *Server) EnviarJugadas(ctx context.Context, in *pb.Jugada) (*pb.Mensajito2, error) {
		crearArchivo()
		for i:=0;i<16;i++{
			rand.Seed(time.Now().UnixNano())
			var n_azar int32 = rand.Int31n(3)

			if n_azar == 0{
				//ip: 10.6.40.169
				var linea string = "Jugador_"+fmt.Sprint(in.Jugador[i])+" Ronda_"+fmt.Sprint(in.Ronda)+" 10.6.40.169"
				escribeArchivo(linea)
				//ir al datanode -> 9120
				var conn *grpc.ClientConn
				conn, err := grpc.Dial("10.6.40.169:9120", grpc.WithInsecure())
				if err != nil {
					log.Fatalf("did not connect: %s", err)
				}
				defer conn.Close()
				c := pb.NewLiderNNClient(conn)


				response, err := c.EnviarADataNode(context.Background(), &pb.Jugada{Jugador: in.Jugador, Ronda: in.Ronda, Muertos: in.Muertos})
				if err != nil {
					log.Fatalf("Error when calling SayHello: %s", err)
				}
				log.Printf("Respuesta: %d", response.Id)


			} else if n_azar == 1{
				//ip: 10.6.40.170
				var linea string = "Jugador_"+fmt.Sprint(in.Jugador[i])+" Ronda_"+fmt.Sprint(in.Ronda)+" 10.6.40.170"
				escribeArchivo(linea)
				//ir al datanode -> 9100
				var conn *grpc.ClientConn
				conn, err := grpc.Dial("10.6.40.169:9100", grpc.WithInsecure())
				if err != nil {
					log.Fatalf("did not connect: %s", err)
				}
				defer conn.Close()
				c := pb.NewLiderNNClient(conn)

				response, err := c.EnviarADataNode(context.Background(), &pb.Jugada{Jugador: in.Jugador, Ronda: in.Ronda, Muertos: in.Muertos})
				if err != nil {
					log.Fatalf("Error when calling SayHello: %s", err)
				}
				log.Printf("Respuesta: %d", response.Id)


			} else {
				// ip: 10.6.40.172
				var linea string = "Jugador_"+fmt.Sprint(in.Jugador[i])+" Ronda_"+fmt.Sprint(in.Ronda)+" 10.6.40.172"
				escribeArchivo(linea)
				//ir al datanode -> 9110
				var conn *grpc.ClientConn
				conn, err := grpc.Dial("10.6.40.169:9110", grpc.WithInsecure())
				if err != nil {
					log.Fatalf("did not connect: %s", err)
				}
				defer conn.Close()
				c := pb.NewLiderNNClient(conn)

				response, err := c.EnviarADataNode(context.Background(), &pb.Jugada{Jugador: in.Jugador, Ronda: in.Ronda, Muertos: in.Muertos})
				if err != nil {
					log.Fatalf("Error when calling SayHello: %s", err)
				}
				log.Printf("Respuesta: %d", response.Id)

			}

		}

		return &pb.Mensajito2{Id: 1}, nil
	}

func main(){
	conexionNameLider()
}
