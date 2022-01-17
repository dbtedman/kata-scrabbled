package resource

import (
	"github.com/dbtedman/kata-scrabbled/internal/domain/board"
	"github.com/dbtedman/kata-scrabbled/internal/domain/board_square"
	"github.com/dbtedman/kata-scrabbled/internal/domain/square"
	"github.com/dbtedman/kata-scrabbled/internal/domain/tile"
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

type Index struct {
	BoardsRepository *board.Boards
	TilesRepository  *tile.Tiles
}

type IndexData struct {
	Squares     [][]board_square.BoardSquare
	YourSquares [][]board_square.BoardSquare
	Suggestions []board_square.Suggestion
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
	useCase := square.RenderSquares{
		BoardsRepository: ir.BoardsRepository,
	}
	squares, _ := useCase.Run()

	var yourSquares [][]board_square.BoardSquare
	var yourSquare []board_square.BoardSquare

	existingTiles, _ := ir.TilesRepository.List()

	for _, existingTile := range existingTiles {
		yourSquare = append(yourSquare, board_square.BoardSquare{
			Id:    existingTile.Id,
			Value: existingTile.Value,
		})
	}

	yourSquares = append(yourSquares, yourSquare)

	suggestions := []board_square.Suggestion{{
		Word:             strings.ToUpper("Word"),
		BoardSquareStart: "A1",
		BoardSquareEnd:   "A4",
		Score:            12,
	}}

	data := IndexData{Squares: squares, YourSquares: yourSquares, Suggestions: suggestions}

	fp := path.Join("web", "template", "index.html")
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
