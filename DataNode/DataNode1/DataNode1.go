package main

import (
	pb "Lab2_SD/Pb"
	"context"
	"google.golang.org/grpc"
	"log"
	"fmt"
  	"os"
	"net"
)
type Server struct {
	pb.UnimplementedClienteSvServer
}

var path = "/home/alumno/Lab2_SD/DataNode/DataNode1/"

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

func (s *Server) EnviarADataNode(ctx context.Context, in *pb.Jugada) (*pb.Mensajito2, error) {
	for i:=0;i<16;i++ {
		path = path+"jugador_"+fmt.Sprint(i)+"_ronda_"+fmt.Sprint(in.Ronda)
		crearArchivo()
		var linea string = fmt.Sprint(in.Jugador[i])
		escribeArchivo(linea)
	}
	return &pb.Mensajito2{Id: 1}, nil
}


func ServerJugador(){
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9100))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterLiderNNServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}


func main(){
	ServerJugador()
}