package main

import (
	pb "clientGRPC/client"
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ctx = context.Background()

type Data struct {
	Texto string
	Pais  string
}

func sendData(c *fiber.Ctx) error {
	var data []Data // Change to slice of Data
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	// Initialize gRPC connection
	conn, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("could not close connection: %v", err)
		}
	}(conn)

	// Create a client for the gRPC service
	cl := pb.NewGetInfoClient(conn)

	// Iterate over each object in the JSON array and send it to the gRPC server
	for _, tweet := range data {
		ret, err := cl.ReturnInfo(context.Background(), &pb.RequestId{
			Texto: tweet.Texto,
			Pais:  tweet.Pais,
		})
		if err != nil {
			return err
		}

		fmt.Println("Respuesta del servidor ", ret)
	}

	return nil
}

func main() {
	app := fiber.New()

	// Endpoint to handle POST requests
	app.Post("/insert", sendData)

	// Start the Fiber server
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

// func sendData(c *fiber.Ctx) error {
// 	var data map[string]string
// 	e := c.BodyParser(&data)
// 	if e != nil {
// 		return e
// 	}

// 	tweet := Data{
// 		Texto: data["texto"],
// 		Pais:  data["pais"],
// 	}

// 	conn, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()),
// 		grpc.WithBlock())
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}

// 	cl := pb.NewGetInfoClient(conn)
// 	defer func(conn *grpc.ClientConn) {
// 		err := conn.Close()
// 		if err != nil {
// 			log.Fatalf("could not close connection: %v", err)
// 		}
// 	}(conn)

// 	ret, err := cl.ReturnInfo(ctx, &pb.RequestId{
// 		Texto: tweet.Texto,
// 		Pais:  tweet.Pais,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("Respuesta del servidor ", ret)

// 	return nil
// }

// func main() {
// 	app := fiber.New()

// 	app.Post("/insert", sendData)

// 	err := app.Listen(":3000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
