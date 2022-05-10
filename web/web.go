package web

import (
	"github.com/dbtedman/kata-scrabbled/internal/domain/board"
	"github.com/dbtedman/kata-scrabbled/internal/domain/tile"
	"github.com/dbtedman/kata-scrabbled/web/resource"
	"log"
	"net/http"
)

type Web struct {
	Addr string
}

func (w Web) Listen() error {
	boardsRepository := board.Boards{}
	tilesRepository := tile.Tiles{}

	boardsRepository.Seed()
	tilesRepository.Seed()

	anIndex := resource.Index{
		BoardsRepository: &boardsRepository,
		TilesRepository:  &tilesRepository,
	}
	aBoard := resource.Board{
		BoardsRepository: &boardsRepository,
	}
	someTiles := resource.Tiles{
		TilesRepository: &tilesRepository,
	}
	aRecommended := resource.Recommended{}

	// Support serving of static assets
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", anIndex.Handle)
	http.HandleFunc("/board", aBoard.Handle)
	http.HandleFunc("/tiles", someTiles.Handle)
	http.HandleFunc("/recommended", aRecommended.Handle)

	log.Printf("Listening to requests on %s...", w.Addr)

	return http.ListenAndServe(w.Addr, nil)
}
