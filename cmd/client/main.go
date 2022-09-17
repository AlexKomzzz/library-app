package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/AlexKomzzz/library-app/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connect client: ", err)
	}

	defer conn.Close()

	client := api.NewLibraryClient(conn)

	for {
		in := bufio.NewReader(os.Stdin)

		req, _, err := in.ReadLine()
		if err != nil {
			log.Fatal("the request failed: ", err)
		}

		request := strings.Split(string(req), " ")

		if len(request) < 2 {
			log.Fatal("not enough arguments")
		}

		switch {
		case request[0] == "book": // если первое слово book значит вызываем метод поиска автора по книге
			nameBook := request[1]
			// считываем название книги
			for i := 2; i < len(request); i++ {
				nameBook = fmt.Sprint(nameBook, " ", request[i])
			}

			book := &api.Book{
				Title: nameBook,
			}

			res, err := client.SearchAuthor(context.Background(), book)
			if err != nil {
				log.Fatal("error seachAuthor: ", err)
			}

			fmt.Println(res.GetAuthors())

		case request[0] == "author":

			author := &api.Author{
				Name: request[1],
			}

			res, err := client.SearchBook(context.Background(), author)
			if err != nil {
				log.Fatal("error searchBook: ", err)
			}

			fmt.Println(res.GetBooks())
		}
	}
}
