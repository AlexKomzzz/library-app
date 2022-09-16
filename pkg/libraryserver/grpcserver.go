package libraryserver

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AlexKomzzz/library-app/pkg/api"
)

const (
	books = "books"
	auth  = "authors"
	a_b   = "authors_books"
)

type GRPCserver struct {
	api.UnimplementedLibraryServer
	db *sql.DB
}

func NewGRPCServer(db *sql.DB) *GRPCserver {
	return &GRPCserver{
		db: db,
	}
}

func (s *GRPCserver) SearchAuthor(ctx context.Context, book *api.Book) (*api.Authors, error) {

	query := fmt.Sprintf("SELECT %s.surname FROM %s JOIN %s ON books.id = %s.id_book JOIN %s ON %s.id_author = %s.id WHERE books.title = ?", auth, books, a_b, a_b, auth, a_b, auth)
	row, err := s.db.QueryContext(ctx, query, book.Title)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	var result string
	for row.Next() {
		var surname string
		err := row.Scan(&surname)
		if err != nil {
			return nil, err
		}

		if result == "" {
			result = surname
		} else {
			result = fmt.Sprintf("%s, %s", result, surname)
		}
	}

	if result == "" {
		result = "У этой книги нет автора. Проверьте правильность названия книги."
	}

	err = row.Close()
	if err != nil {
		return nil, err
	}

	return &api.Authors{Authors: result}, nil
}

func (s *GRPCserver) SearchBook(ctx context.Context, author *api.Author) (*api.Books, error) {
	query := fmt.Sprintf("SELECT books.title FROM %s JOIN %s ON %s.id = %s.id_author JOIN books ON %s.id_book = books.id WHERE %s.surname = ?", auth, a_b, auth, a_b, a_b, auth)
	row, err := s.db.QueryContext(ctx, query, author.Name)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var result string

	for row.Next() {
		var title string
		err := row.Scan(&title)
		if err != nil {
			return nil, err
		}

		if result == "" {
			result = title
		} else {
			result = fmt.Sprintf("%s, %s", result, title)
		}
	}

	if result == "" {
		result = "У этого автора нет книг. Проверьте правильность фамилии автора."
	}

	err = row.Close()
	if err != nil {
		return nil, err
	}

	return &api.Books{Books: result}, nil
}
