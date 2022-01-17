package domain

import (
	"strconv"
)

type Boards struct {
	boardSquares [][]BoardSquareValue
}

func (receiver *Boards) Seed() {
	for _, i := range []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O"} {
		var row []BoardSquareValue
		for j := range [15]int{} {
			row = append(row, BoardSquareValue{
				Id:    "board" + i + strconv.Itoa(j+1),
				Value: "",
			})
		}
		receiver.boardSquares = append(receiver.boardSquares, row)
	}
}

func (receiver *Boards) List() ([][]BoardSquareValue, error) {
	return receiver.boardSquares, nil
}

func (receiver *Boards) Update(value [][]BoardSquareValue) error {
	receiver.boardSquares = value
	return nil
}
