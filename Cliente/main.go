package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "lab2/Pb"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.40.169:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewClienteSvClient(conn)

	response, err := c.DimeHola(context.Background(), &pb.Mensaje{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}
