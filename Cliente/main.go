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
	//Comienza conexión con el Lider, le pide jugar
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
	// Termina de pedirle al Lider jugar, recibe su respuesta

	//var id int32 = 0

	//-------------------- Comienzo a manejar la etapa 1 del juego -----------------------------------------
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

	//for que maneja rondas, jugadas del jugador por pantalla y de los jugadores bot
	for i:=0; i<=4; i++{
		if i < 4 {
			if muertos[0] == 0{ //si jugador por pantalla sigue vivo
				if jugadas2[0] >= 21 {  //ya tiene 21 o más, no sigue jugando (se registra jugada como 0)
					jugadas[0] = 0
				}else{  //está vivo, no tiene 21 o más, así que sigue jugando
					log.Printf("\nEscoja un número del 1 al 10: ")
					fmt.Scan(&opcion)
					jugadas[cont] = opcion	
				}
			}else{ //jugador por pantalla está vivo
				jugadas[0] = 0
			}
			for cont:=1;cont<16;cont++{  // for para manejar las jugadas de los bots
				if jugadas2[cont] < 21 {
					if muertos[cont] == 1{ //jugador bot esta muerto, se registra su jugada como 0
						jugadas[cont] = 0
					}else{  //jugador bot está vivo, se escoge un numero al azar entre 1 y 10.
						jugadas[cont] = rand.Int31n(10) +1
					}					
				} else { //ya tiene 21 o más, siguiente jugada se registra como 0
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
	// ------------------- FIN ETAPA 1 ---------------------------------------

}
