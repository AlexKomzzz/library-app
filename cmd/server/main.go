package main

import (
	"log"
	"net"

	"github.com/AlexKomzzz/library-app/pkg/api"
	"github.com/AlexKomzzz/library-app/pkg/libraryserver"
	"github.com/AlexKomzzz/library-app/pkg/repository"
	"google.golang.org/grpc"
)

func main() {

	// подключение в БД MySQL
	db, err := repository.NewMysqlDB()
	if err != nil {
		log.Fatal("failed to initialize db: ", err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	srv := libraryserver.NewGRPCServer(db)
	s := grpc.NewServer()
	api.RegisterLibraryServer(s, srv)

	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: ", err)
	}
}
