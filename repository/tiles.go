package repository

import "github.com/dbtedman/kata-scrabbled/entity"

type Tiles struct {
	tiles []entity.BoardSquareValue
}

func (receiver *Tiles) Seed() {
	receiver.tiles = []entity.BoardSquareValue{
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

func (receiver *Tiles) List() ([]entity.BoardSquareValue, error) {
	return receiver.tiles, nil
}

func (receiver *Tiles) Update(value []entity.BoardSquareValue) error {
	receiver.tiles = value
	return nil
}
