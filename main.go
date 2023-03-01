package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/achmadrizkin/database"
	pb "github.com/achmadrizkin/proto"
	"github.com/achmadrizkin/server"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// jika kode mengalami crash, nomor line akan ditampilkan
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Book service started")

	// membuat koneksi ke DB
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	// membuat gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	// melakukan register BookServiceServer
	pb.RegisterMovieServiceServer(s, &server.Server{})
	// mengaktifkan reflection
	// agar bisa digunakan untuk pengujian dengan evans
	reflection.Register(s)

	go func() {
		fmt.Println("Starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Menunggu hingga dihentikan dengan Ctrl + C
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Lakukan block hingga sinyal sudah didapatkan
	<-ch
	fmt.Println("Stopping the server..")
	s.Stop()
	fmt.Println("Stopping listener...")
	lis.Close()
	fmt.Println("End of Program")
}
