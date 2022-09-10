package libraryserver

import (
	"context"
	"database/sql"

	"github.com/AlexKomzzz/library-app/pkg/api"
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
	return &api.Authors{}, nil
}

func (s *GRPCserver) SearchBook(ctx context.Context, author *api.Author) (*api.Books, error) {
	return &api.Books{}, nil
}
