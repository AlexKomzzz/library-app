package main

import (
	"context"
	"flag"
	"fmt"
	"log"
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

	//for {
	//in := bufio.NewReader(os.Stdin)

	//req, _, err := in.ReadLine()
	//

	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatal("not enough arguments")
	}

	switch {
	case flag.Arg(0) == "book": // если первое слово book значит вызываем метод поиска автора по книге
		var nameBook string
		// считываем название книги
		for i := 1; i < flag.NArg(); i++ {
			if i == 1 {
				nameBook = flag.Arg(i)
			} else {
				nameBook = fmt.Sprint(nameBook, " ", strings.Replace(flag.Arg(i), " ", "", -1))
			}
		}

		book := &api.Book{
			Title: nameBook,
		}

		res, err := client.SearchAuthor(context.Background(), book)
		if err != nil {
			log.Fatal("error seachAuthor: ", err)
		}

		fmt.Println(res.GetAuthors())

	case flag.Arg(0) == "author":

		author := &api.Author{
			Name: flag.Arg(1),
		}
		res, err := client.SearchBook(context.Background(), author)
		if err != nil {
			log.Fatal("error searchBook: ", err)
		}

		fmt.Println(res.GetBooks())
	}
	//}
}
