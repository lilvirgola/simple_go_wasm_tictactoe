package ai

import (
	"math/rand"
	"wasm_tictactoe/internal/board"
	"wasm_tictactoe/internal/global"
)

var (
	coefficents = map[global.Difficulty]float64{
		global.EasyDifficulty:   0.5,
		global.MediumDifficulty: 0.75,
		global.HardDifficulty:   1.0,
	}
)

// BestMove returns the index of the best move for `mark` on board `b` using minmax or a random move based on the difficulty level
func BestMove(b board.Board, mark string, difficulty global.Difficulty) int {
	if rand.Float64() > coefficents[difficulty] {
		moves := b.AvailableMoves()
		return moves[rand.Intn(len(moves))]
	}

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
