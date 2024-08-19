package models

type RollStats struct {
	Rolls       []int
	RollAmounts []int
}

func NewRollStats() RollStats {
	return RollStats{
		Rolls:       []int{},
		RollAmounts: []int{0, 0, 0, 0, 0, 0},
	}

}

// type PlaceHolderUser struct {
// 	Email    string `json:"email"`
// 	Username string `json:"username"`
// }
