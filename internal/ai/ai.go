package ai

import "wasm_tictactoe/internal/board"

// BestMove returns the index of the best move for `mark` on board `b` using minmax
func BestMove(b board.Board, mark string) int {
	_, move := minimax(b, mark, mark, 0)
	return move
}

// minimax implements the minimax algorithm. It returns the best score and the index of the best move for `toMove` on board `b`.
func minimax(b board.Board, maximizer, toMove string, depth int) (int, int) {
	if over, winner := b.Terminal(); over {
		switch winner {
		case maximizer:
			return 10 - depth, -1
		case "":
			return 0, -1
		default:
			return depth - 10, -1
		}
	}

	maximizing := toMove == maximizer
	bestMove := -1
	bestScore := 1000
	if maximizing {
		bestScore = -1000
	}

	for _, i := range b.AvailableMoves() {
		next := b.Played(i, toMove)
		score, _ := minimax(next, maximizer, board.Other(toMove), depth+1)

		if maximizing && score > bestScore {
			bestScore, bestMove = score, i
		} else if !maximizing && score < bestScore {
			bestScore, bestMove = score, i
		}
	}
	return bestScore, bestMove
}
