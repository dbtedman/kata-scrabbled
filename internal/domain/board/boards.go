package board

import (
	"github.com/dbtedman/kata-scrabbled/internal/domain/board_square"
	"strconv"
)

type Boards struct {
	boardSquares [][]board_square.BoardSquareValue
}

func (receiver *Boards) Seed() {
	for _, i := range []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O"} {
		var row []board_square.BoardSquareValue
		for j := range [15]int{} {
			row = append(row, board_square.BoardSquareValue{
				Id:    "board" + i + strconv.Itoa(j+1),
				Value: "",
			})
		}
		receiver.boardSquares = append(receiver.boardSquares, row)
	}
}

func (receiver *Boards) List() ([][]board_square.BoardSquareValue, error) {
	return receiver.boardSquares, nil
}

func (receiver *Boards) Update(value [][]board_square.BoardSquareValue) error {
	receiver.boardSquares = value
	return nil
}
