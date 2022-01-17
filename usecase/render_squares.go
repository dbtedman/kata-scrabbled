package usecase

import (
	"github.com/dbtedman/kata-scrabbled/entity"
	"github.com/dbtedman/kata-scrabbled/repository"
	"strings"
)

type RenderSquares struct {
	BoardsRepository *repository.Boards
}

func (rs RenderSquares) Run() ([][]entity.BoardSquare, error) {
	var squares [][]entity.BoardSquare

	squareValues, _ := rs.BoardsRepository.List()

	for _, squareValueRow := range squareValues {
		var row []entity.BoardSquare

		for squareValueCount, squareValue := range squareValueRow {
			row = append(row, entity.BoardSquare{
				Label:         strings.Replace(squareValue.Id, "board", "", 1),
				ModifierClass: rs.tileClass(string([]rune(strings.Replace(squareValue.Id, "board", "", 1))[0]), squareValueCount+1),
				Id:            squareValue.Id,
				Value:         squareValue.Value,
			})
		}

		squares = append(squares, row)
	}

	return squares, nil
}

func (rs RenderSquares) tileClass(row string, column int) string {
	if row == "H" && column == 8 {
		return "tile-middle"
	}

	if (row == "A" || row == "H" || row == "O") && (column == 1 || column == 8 || column == 15) {
		return "tile-3w"
	}

	if ((row == "B" || row == "N") && (column == 2 || column == 14)) || (row == "C" || row == "M") && (column == 3 || column == 13) || (row == "D" || row == "L") && (column == 4 || column == 12) || (row == "E" || row == "K") && (column == 5 || column == 11) {
		return "tile-2w"
	}

	if (row == "A" || row == "O") && (column == 4 || column == 12) || (row == "D" || row == "L") && (column == 1 || column == 8 || column == 15) || ((row == "C" || row == "M") && (column == 7 || column == 9)) || ((row == "G" || row == "I") && (column == 3 || column == 7 || column == 9 || column == 13)) || (row == "H" && (column == 4 || column == 12)) {
		return "tile-2l"
	}

	if ((row == "B" || row == "N") && (column == 6 || column == 10)) || ((row == "F" || row == "J") && (column == 2 || column == 6 || column == 10 || column == 14)) {
		return "tile-3l"
	}

	return ""
}
