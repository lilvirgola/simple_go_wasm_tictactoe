// Package state owns the mutable game session: whose turn it is, whether
// the AI is enabled, and the win/draw outcome. It orchestrates the pure
// board package and the ai package but holds all the mutable state itself.
package state

import (
	"wasm_tictactoe/internal/ai"
	"wasm_tictactoe/internal/board"
)

// Game is one tic-tac-toe session.
type Game struct {
	Board       board.Board
	Turn        string
	Winner      string // "", "X", "O", or "Draw"
	AIEnabled   bool
	HumanPlayer string
	AIPlayer    string
}

// New returns a fresh game with the human playing X, AI playing O, and AI
// assistance enabled by default.
func New() *Game {
	g := &Game{}
	g.reset()
	return g
}

func (g *Game) reset() {
	g.Board = board.Board{}
	g.Turn = "X"
	g.Winner = ""
	g.HumanPlayer = "X"
	g.AIPlayer = "O"
}

// Reset starts a new round, preserving the AIEnabled setting.
func (g *Game) Reset() {
	g.reset()
}

// ToggleAI flips whether the AI plays for AIPlayer.
func (g *Game) ToggleAI() {
	g.AIEnabled = !g.AIEnabled
}

// Over reports whether the round has ended.
func (g *Game) Over() bool {
	return g.Winner != ""
}

// Move plays `mark` at cell i for whoever's turn it currently is. It
// refuses the move (returning false) if the cell is taken, the round is
// over, or it is not a human's turn to act (i.e. it's the AI's turn while
// AI assistance is enabled). After a valid human move it automatically
// plays the AI's reply, if applicable.
func (g *Game) Move(i int) bool {
	if !g.canHumanMove(i) {
		return false
	}

	g.play(i, g.Turn)

	if !g.Over() && g.AIEnabled && g.Turn == g.AIPlayer {
		g.playAIMove()
	}
	return true
}

func (g *Game) canHumanMove(i int) bool {
	if i < 0 || i > 8 {
		return false
	}
	if g.Board[i] != "" || g.Over() {
		return false
	}
	if g.AIEnabled && g.Turn == g.AIPlayer {
		return false
	}
	return true
}

// play places mark at i, updates the board, and settles win/draw/turn
// bookkeeping. Caller is responsible for validating the move first.
func (g *Game) play(i int, mark string) {
	g.Board = g.Board.Played(i, mark)

	if over, winner := g.Board.Terminal(); over {
		if winner == "" {
			g.Winner = "Draw"
		} else {
			g.Winner = winner
		}
		return
	}
	g.Turn = board.Other(g.Turn)
}

func (g *Game) playAIMove() {
	i := ai.BestMove(g.Board, g.AIPlayer)
	if i == -1 {
		return
	}
	g.play(i, g.AIPlayer)
}
