package resource

import (
	"github.com/dbtedman/kata-scrabbled/internal/domain"
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

type Index struct {
	BoardsRepository *domain.Boards
	TilesRepository  *domain.Tiles
}

type IndexData struct {
	Squares     [][]domain.BoardSquare
	YourSquares [][]domain.BoardSquare
	Suggestions []domain.Suggestion
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
	useCase := domain.RenderSquares{
		BoardsRepository: ir.BoardsRepository,
	}
	squares, _ := useCase.Run()

	var yourSquares [][]domain.BoardSquare
	var yourSquare []domain.BoardSquare

	existingTiles, _ := ir.TilesRepository.List()

	for _, existingTile := range existingTiles {
		yourSquare = append(yourSquare, domain.BoardSquare{
			Id:    existingTile.Id,
			Value: existingTile.Value,
		})
	}

	yourSquares = append(yourSquares, yourSquare)

	suggestions := []domain.Suggestion{{
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
