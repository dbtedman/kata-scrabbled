package repository

import (
	"github.com/dbtedman/kata-scrabbled/entity"
	"strconv"
)

type Boards struct {
	boardSquares [][]entity.BoardSquareValue
}

func (receiver *Boards) Seed() {
	for _, i := range []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O"} {
		var row []entity.BoardSquareValue
		for j := range [15]int{} {
			row = append(row, entity.BoardSquareValue{
				Id:    "board" + i + strconv.Itoa(j+1),
				Value: "",
			})
		}
		receiver.boardSquares = append(receiver.boardSquares, row)
	}
}

func (receiver *Boards) List() ([][]entity.BoardSquareValue, error) {
	return receiver.boardSquares, nil
}

func (receiver *Boards) Update(value [][]entity.BoardSquareValue) error {
	receiver.boardSquares = value
	return nil
}
