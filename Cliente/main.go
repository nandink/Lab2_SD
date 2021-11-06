package main

import (
	pb "Lab2_SD/Pb"
	"context"
	"google.golang.org/grpc"
	"log"
	"fmt"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.40.169:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewClienteSvClient(conn)

	response, err := c.DimeHola(context.Background(), &pb.Mensaje{Body: "Solicitando inscripcion!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Respuesta del Lider: %s", response.Body)

	var id := 0
	var lista_jug := [15]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	var cont := 0
	for {
		if cont < 16 {
			response2, err := c.MandarJugadores(context.Background(), &pb.Mensajito2{Id: lista_jug[cont]})
			if err != nil {
				log.Fatalf("Error when calling SayHello: %s", err)
			}
			log.Printf("Respuesta del Lider: %s", response2.Id)
		} else{
			return
		}
		cont = cont + 1
	}
}
