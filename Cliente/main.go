package main

import (
	"time"
	pb "Lab2_SD/Pb"
	"context"
	"google.golang.org/grpc"
	"log"
	"fmt"
	"math/rand"
)


func main() {
	rand.Seed(time.Now().UnixNano())
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
	var jugadas2 = []int32{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0} //mantener el registro

	var ronda int32 = 1
	cont = 0 //para recorrer a los jugadores

	log.Printf("\n----Etapa 1: Luz Verde, Luz Roja ----")
	var opcion int32

	//FOR NUEVO AQUI ESTA SI QUE SI VAMOS DIOSITO
	for i:=0; i<=4; i++{
		if i < 4 {
			if muertos[0] == 0{ //si jugador por pantalla sigue vivo, sigue jugando las rondas
				if jugadas2[0] >= 21 {
					jugadas[0] = 0
				}else{
					log.Printf("\nEscoja un número del 1 al 10: ")
					fmt.Scan(&opcion)
					jugadas[cont] = opcion	
				}
			}else{ //esta MUERTO :c
				jugadas[0] = 0
			}
			for cont:=1;cont<16;cont++{
				//si tiene mas de 21, sigue vivo, pero jugada = 0
				//si tiene menos de 21, sigue jugando y tira jugada
				//si esta muerto se le asigna jugada = 0
				if jugadas2[cont] < 21 {
					if muertos[cont] == 1{ //tay dead
						jugadas[cont] = 0
					}else{
						jugadas[cont] = rand.Int31n(10) +1
					}					
				} else {
					jugadas[cont] = 0
				}
			}


			response3, err := c.MandarJugada(context.Background(), &pb.Jugada{Jugador: jugadas, Ronda: ronda, Muertos: muertos})
			if err != nil {
				log.Fatalf("Error %s", err)
			}

			log.Printf("Numero de ronda actual: %d", response3.Ronda)
			log.Printf("\nEl lider escogió su número. Los jugadores muertos son: ")
			
			
			for cont:=0;cont<16;cont++{
				if response3.Muertos[cont] == 1{
					log.Printf(" %d ", cont)
				}
			
			}

			for cont:=0;cont<16;cont++{
				jugadas2[cont] = jugadas2[cont] + jugadas[cont]
			}

			for cont:=0;cont<16;cont++{
				muertos[cont] = response3.Muertos[cont]
			}
			log.Printf("\nFin de la ronda %d", ronda)
			ronda = ronda + 1
		
		} else{ 
			
			response3, err := c.MandarJugada(context.Background(), &pb.Jugada{Jugador: jugadas2, Ronda: ronda, Muertos: muertos})
			if err != nil {
				log.Fatalf("Error %s", err)
			}
			log.Printf("Numero de ronda actual: %d", response3.Ronda)
			log.Printf("Fin de la Etapa 1")

		}
	}

}
