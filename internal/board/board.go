package board

// possible winning lines
var lines = [8][3]int{
	{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
	{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
	{0, 4, 8}, {2, 4, 6},
}

type Board [9]string

// Empty checks whether the board is empty
func (b Board) Empty() bool {
	return b == Board{}
}

// Full checks whether the board is full
func (b Board) Full() bool {
	for _, v := range b {
		if v == "" {
			return false
		}
	}
	return true
}

// Winner returns "X" or "O" if that player has three in a row, otherwise "".
func (b Board) Winner() string {
	for _, l := range lines {
		a, c, d := b[l[0]], b[l[1]], b[l[2]]
		if a != "" && a == c && c == d {
			return a
		}
	}
	return ""
}

// Terminal checks whether the game is over (win or full board)
func (b Board) Terminal() (over bool, winner string) {
	if w := b.Winner(); w != "" {
		return true, w
	}
	return b.Full(), ""
}

// AvailableMoves returns the indices of every empty cell
func (b Board) AvailableMoves() []int {
	moves := make([]int, 0, 9)
	for i, v := range b {
		if v == "" {
			moves = append(moves, i)
		}
	}
	return moves
}

// Played returns a copy of b with mark placed at i.
func (b Board) Played(i int, mark string) Board {
	b[i] = mark
	return b
}

// Other returns the opposing mark
func Other(mark string) string {
	if mark == "X" {
		return "O"
	}
	return "X"
}
