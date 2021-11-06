package main

import (
	pb "Lab2_SD/Pb"
	"context"
	"google.golang.org/grpc"
	"log"
	"fmt"
	"math/rand"
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

	//var id int32 = 0
	var lista_jug  = [16]int32{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	var cont int32 = 1
	for {
		if cont < 16 {
			response2, err := c.MandarJugadores(context.Background(), &pb.Mensajito2{Id: lista_jug[cont]})
			if err != nil {
				log.Fatalf("Error when calling SayHello: %s", err)
			}
			log.Printf("Respuesta del Lider: %s", response2.Id)
		} else{
			fmt.Printf("Inicio Etapa 1")
			break
		}
		cont = cont + 1
	}

	var muertos = []int32{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	var jugadas = []int32{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	cont = 0 //para recorrer a los jugadores
	//random rand.Intn(3)
	log.Printf("\n----Etapa 1: Luz Verde, Luz Roja ----")
	log.Printf("\nEscoja un número del 1 al 10: ")
	var opcion int32
	for {
		if cont == 0{
			fmt.Scan(&opcion)
			jugadas[cont] = opcion
		} else if cont < 16{
			jugadas[cont] = rand.Int31n(10) +1
		} else{
			log.Printf("\nTodos los jugadores escogieron su número.")
			break
		}
		cont = cont + 1
	}
	//recibir jugada lider
	response3, err := c.MandarJugada(context.Background(), &pb.Jugada{Jugador: jugadas, Ronda: 1, Muertos: muertos})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Respuesta del Lider: %d", response3.Ronda)
	log.Printf("\nEl lider escogió su número. Los jugadores muertos son: ")

}
