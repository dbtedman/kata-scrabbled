package entity

type BoardSquare struct {
	Label         string
	Id            string
	Value         string
	ModifierClass string
}

type BoardSquareValue struct {
	Id    string
	Value string
}

type Suggestion struct {
	Word             string
	BoardSquareStart string
	BoardSquareEnd   string
	Score            int
}
