package web

import (
	"github.com/dbtedman/kata-scrabbled/entity"
	"github.com/dbtedman/kata-scrabbled/repository"
	"github.com/dbtedman/kata-scrabbled/usecase"
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

type Index struct {
	BoardsRepository *repository.Boards
	TilesRepository  *repository.Tiles
}

type IndexData struct {
	Squares     [][]entity.BoardSquare
	YourSquares [][]entity.BoardSquare
	Suggestions []entity.Suggestion
}

func (ir Index) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ir.handleGet(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Render html to display the current board, your tiles, available actions, and
// suggestions for next moves.
func (ir Index) handleGet(w http.ResponseWriter, r *http.Request) {
	useCase := usecase.RenderSquares{
		BoardsRepository: ir.BoardsRepository,
	}
	squares, _ := useCase.Run()

	var yourSquares [][]entity.BoardSquare
	var yourSquare []entity.BoardSquare

	existingTiles, _ := ir.TilesRepository.List()

	for _, existingTile := range existingTiles {
		yourSquare = append(yourSquare, entity.BoardSquare{
			Id:    existingTile.Id,
			Value: existingTile.Value,
		})
	}

	yourSquares = append(yourSquares, yourSquare)

	suggestions := []entity.Suggestion{{
		Word:             strings.ToUpper("Word"),
		BoardSquareStart: "A1",
		BoardSquareEnd:   "A4",
		Score:            12,
	}}

	data := IndexData{Squares: squares, YourSquares: yourSquares, Suggestions: suggestions}

	fp := path.Join("template", "index.html")
	tmpl, err := template.ParseFiles(fp)

	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	//_, _ = fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
