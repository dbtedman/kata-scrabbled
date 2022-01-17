package resource

import (
	"encoding/json"
	"fmt"
	"github.com/dbtedman/kata-scrabbled/internal/domain"
	"net/http"
)

type Tiles struct {
	TilesRepository *domain.Tiles
}

func (ti Tiles) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		ti.handlePost(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (ti Tiles) handlePost(w http.ResponseWriter, r *http.Request) {
	var body tilesHandlePostRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Request body did not match expected format.", http.StatusBadRequest)
		return
	}

	converted, _ := convertTilesToBoardSquareValue(body.Tiles)
	_ = ti.TilesRepository.Update(converted)

	data := tilesHandlePostResponse{}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Something went wrong on our end.", http.StatusInternalServerError)
		return
	}
}

func convertTilesToBoardSquareValue(tiles []tilesValue) ([]domain.BoardSquareValue, error) {
	var resultRow []domain.BoardSquareValue
	for _, tile := range tiles {
		resultRow = append(resultRow, domain.BoardSquareValue{
			Value: tile.Value,
			Id:    tile.Id,
		})
	}
	return resultRow, nil
}

type tilesHandlePostRequestBody struct {
	Tiles []tilesValue `json:"tiles"`
}

type tilesValue struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

type tilesHandlePostResponse struct {
}
