package main

import (
	//pb "Lab2_SD/Pb"
	//"context"
	//"google.golang.org/grpc"
	"log"
	"fmt"
  	"os"
)

var path = "/home/sofia/Escritorio/Tareas/SD_Final/Lab2_SD/NameNode/archivo.txt"

type Server struct {
	pb.UnimplementedClienteSvServer
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

func main(){
	conexionNameLider()
}