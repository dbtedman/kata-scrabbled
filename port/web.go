package port

import (
	"github.com/dbtedman/kata-scrabbled/port/web"
	"github.com/dbtedman/kata-scrabbled/repository"
	"log"
	"net/http"
)

type Web struct {
	Addr string
}

func (w Web) Listen() error {
	boardsRepository := repository.Boards{}
	tilesRepository := repository.Tiles{}

	boardsRepository.Seed()
	tilesRepository.Seed()

	index := web.Index{
		BoardsRepository: &boardsRepository,
		TilesRepository:  &tilesRepository,
	}
	board := web.Board{
		BoardsRepository: &boardsRepository,
	}
	tiles := web.Tiles{
		TilesRepository: &tilesRepository,
	}
	recommended := web.Recommended{}

	// Support serving of static assets
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", index.Handle)
	http.HandleFunc("/board", board.Handle)
	http.HandleFunc("/tiles", tiles.Handle)
	http.HandleFunc("/recommended", recommended.Handle)

	log.Printf("Listening to requests on %s...", w.Addr)

	return http.ListenAndServe(w.Addr, nil)
}
