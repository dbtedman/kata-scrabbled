package web

import (
	"github.com/dbtedman/kata-scrabbled/internal/domain"
	resource2 "github.com/dbtedman/kata-scrabbled/web/resource"
	"log"
	"net/http"
)

type Web struct {
	Addr string
}

func (w Web) Listen() error {
	boardsRepository := domain.Boards{}
	tilesRepository := domain.Tiles{}

	boardsRepository.Seed()
	tilesRepository.Seed()

	index := resource2.Index{
		BoardsRepository: &boardsRepository,
		TilesRepository:  &tilesRepository,
	}
	board := resource2.Board{
		BoardsRepository: &boardsRepository,
	}
	tiles := resource2.Tiles{
		TilesRepository: &tilesRepository,
	}
	recommended := resource2.Recommended{}

	// Support serving of static assets
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", index.Handle)
	http.HandleFunc("/board", board.Handle)
	http.HandleFunc("/tiles", tiles.Handle)
	http.HandleFunc("/recommended", recommended.Handle)

	log.Printf("Listening to requests on %s...", w.Addr)

	return http.ListenAndServe(w.Addr, nil)
}
