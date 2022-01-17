package tile

import (
	"github.com/dbtedman/kata-scrabbled/internal/domain/board_square"
)

type Tiles struct {
	tiles []board_square.BoardSquareValue
}

func (receiver *Tiles) Seed() {
	receiver.tiles = []board_square.BoardSquareValue{
		{
			Id:    "yours0",
			Value: "",
		},
		{
			Id:    "yours1",
			Value: "",
		},
		{
			Id:    "yours2",
			Value: "",
		},
		{
			Id:    "yours3",
			Value: "",
		},
		{
			Id:    "yours4",
			Value: "",
		},
		{
			Id:    "yours5",
			Value: "",
		},
		{
			Id:    "yours6",
			Value: "",
		},
	}
}

func (receiver *Tiles) List() ([]board_square.BoardSquareValue, error) {
	return receiver.tiles, nil
}

func (receiver *Tiles) Update(value []board_square.BoardSquareValue) error {
	receiver.tiles = value
	return nil
}
