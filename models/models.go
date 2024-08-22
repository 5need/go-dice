package models

type RollStats struct {
	Rolls         []int
	RollAmounts   []int
	PreviousInput string
	CurrentRoll   string
	Selection     []DiceIdentifier
}

type DiceIdentifier struct {
	RollValue   int
	FromTheLeft int
}

func NewRollStats() RollStats {
	return RollStats{
		Rolls:         []int{},
		RollAmounts:   []int{0, 0, 0, 0, 0, 0},
		PreviousInput: "",
		CurrentRoll:   "",
		Selection:     []DiceIdentifier{},
	}
}
