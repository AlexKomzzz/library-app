package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/AlexKomzzz/library-app/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatal("not enough arguments")
	}

	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connect client: ", err)
	}

	defer conn.Close()

	client := api.NewLibraryClient(conn)

	switch {
	case flag.Arg(0) == "authors":
		var nameBook string
		for i := 1; i < flag.NArg(); i++ {
			nameBook = fmt.Sprint(nameBook, "", flag.Arg(i))
		}

		book := &api.Book{
			Title: nameBook,
		}

		res, err := client.SearchAuthor(context.Background(), book)
		if err != nil {
			log.Fatal("error seachAuthor: ", err)
		}

		fmt.Println(res.GetAuthors())
	case flag.Arg(0) == "book":

		author := &api.Author{
			Name: flag.Arg(1),
		}
		res, err := client.SearchBook(context.Background(), author)
		if err != nil {
			log.Fatal("error searchBook: ", err)
		}

		fmt.Println(res.GetBooks())
	}
}
