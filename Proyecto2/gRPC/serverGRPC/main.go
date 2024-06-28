package main

import (
	"context"
	"fmt"
	"net"
	"serverGRPC/kafka"
	"serverGRPC/model"
	pb "serverGRPC/server"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	tweet := model.Data{
		Texto: in.GetTexto(),
		Pais:  in.GetPais(),
	}

	fmt.Println(tweet)

	kafka.Produce(tweet)

	return &pb.ReplyInfo{Info: "Twitter's recibidos."}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":3001")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}

// package main

// import (
// 	"context"
// 	"fmt"
// 	"net"
// 	"serverGRPC/kafka"
// 	"serverGRPC/model"
// 	pb "serverGRPC/server"

// 	"google.golang.org/grpc"
// )

// type server struct {
// 	pb.UnimplementedGetInfoServer
// }

// // type Data struct {
// // 	Texto string
// // 	Pais  string
// // }

// func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
// 	tweet := model.Data{
// 		Texto: in.GetTexto(),
// 		Pais:  in.GetPais(),
// 	}

// 	fmt.Println(tweet)
// 	kafka.Produce(tweet)

// 	return &pb.ReplyInfo{Info: "Twitter's recibidos."}, nil
// }

// func main() {
// 	listen, err := net.Listen("tcp", ":3001")
// 	if err != nil {
// 		panic(err)
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterGetInfoServer(s, &server{})

// 	if err := s.Serve(listen); err != nil {
// 		panic(err)
// 	}
// }

// // package main

// // import (
// // 	"context"
// // 	"net"

// // 	"serverGRPC/kafka"
// // 	"serverGRPC/model"
// // 	pb "serverGRPC/server"

// // 	"google.golang.org/grpc"
// // )

// // type server struct {
// // 	pb.UnimplementedGetInfoServer
// // 	// rdb           *redis.Client
// // 	// totalMensajes int64
// // 	// mu            sync.Mutex
// // }

// // var ctx = context.Background()

// // func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
// // 	// Procesar la solicitud recibida
// // 	// s.mu.Lock()
// // 	// defer s.mu.Unlock()
// // 	// s.totalMensajes++

// // 	// err := s.rdb.HIncrBy(ctx, "paises", in.GetPais(), 1).Err()
// // 	// if err != nil {
// // 	// 	return nil, err
// // 	// }

// // 	// // Incrementar el contador global y enviar a Redis con clave "total_messages"
// // 	// err = s.rdb.Set(ctx, "cantMensajes", s.totalMensajes, 0).Err()
// // 	// if err != nil {
// // 	// 	return nil, err
// // 	// }
// // 	tweet := model.Data{
// // 		Texto: in.GetTexto(),
// // 		Pais:  in.GetPais(),
// // 	}

// // 	kafka.Produce(tweet)

// // 	// Devolver la respuesta con los datos procesados
// // 	return &pb.ReplyInfo{
// // 		Info: "Twitter recibido.",
// // 	}, nil
// // }

// // func main() {
// // 	// Inicializar el cliente Redis
// // 	// rdb := redis.NewClient(&redis.Options{
// // 	// 	Addr:     "redis:6379",
// // 	// 	Password: "",
// // 	// 	DB:       0,
// // 	// })

// // 	// // Probar la conexión con Redis
// // 	// _, err := rdb.Ping(ctx).Result()
// // 	// if err != nil {
// // 	// 	log.Fatalf("No se pudo conectar a Redis: %v", err)
// // 	// }

// // 	// Configurar el servidor gRPC para escuchar en el puerto 3001
// // 	listen, err := net.Listen("tcp", ":3001")
// // 	if err != nil {
// // 		panic(err)
// // 	}

// // 	// Crear una instancia del servidor gRPC
// // 	s := grpc.NewServer()

// // 	// Registrar el servicio gRPC generado y el servidor personalizado
// // 	//pb.RegisterGetInfoServer(s, &server{rdb: rdb})

// // 	// Iniciar el servidor gRPC
// // 	if err := s.Serve(listen); err != nil {
// // 		panic(err)
// // 	}
// // }
