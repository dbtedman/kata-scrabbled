package domain

type Tiles struct {
	tiles []BoardSquareValue
}

func (receiver *Tiles) Seed() {
	receiver.tiles = []BoardSquareValue{
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

func (receiver *Tiles) List() ([]BoardSquareValue, error) {
	return receiver.tiles, nil
}

func (receiver *Tiles) Update(value []BoardSquareValue) error {
	receiver.tiles = value
	return nil
}
