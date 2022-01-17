package resource

import (
	"encoding/json"
	"fmt"
	"github.com/dbtedman/kata-scrabbled/internal/domain"
	"net/http"
)

type Board struct {
	BoardsRepository *domain.Boards
}

func (ir Board) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		ir.handlePost(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Record changes to the current board, updates will occur as dumps of the entire 15x15
// board.
func (ir Board) handlePost(w http.ResponseWriter, r *http.Request) {
	var body handlePostRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Request body did not match expected format.", http.StatusBadRequest)
		return
	}

	converted, _ := convert(body.Board)
	_ = ir.BoardsRepository.Update(converted)

	data := handlePostResponse{}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Something went wrong on our end.", http.StatusInternalServerError)
		return
	}
}

func convert(board [][]boardSquareValue) ([][]domain.BoardSquareValue, error) {
	var result [][]domain.BoardSquareValue

	for _, boardRow := range board {
		var resultRow []domain.BoardSquareValue

		for _, boardColumn := range boardRow {
			resultRow = append(resultRow, domain.BoardSquareValue{
				Value: boardColumn.Value,
				Id:    boardColumn.Id,
			})
		}

		result = append(result, resultRow)
	}

	return result, nil
}

type handlePostRequestBody struct {
	Board [][]boardSquareValue `json:"board"`
}

type boardSquareValue struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

type handlePostResponse struct {
}
