package main

import (
	"strconv"
	"syscall/js"

	"wasm_tictactoe/internal/global"
	"wasm_tictactoe/internal/state"
)

var game = state.New() // init the game state

// func to update the HTML page to reflect the current game state. Called after every move.
func render(this js.Value, args []js.Value) any {
	document := js.Global().Get("document")

	for i := 0; i < 9; i++ {
		cell := document.Call("getElementById", strconv.Itoa(i))
		cell.Set("innerText", game.Board[i])
	}

	status := document.Call("getElementById", "status")
	switch game.Winner {
	case "":
		status.Set("innerText", "Turn: "+game.Turn)
	case "Draw":
		status.Set("innerText", "Draw!")
	default:
		status.Set("innerText", "Winner: "+game.Winner)
	}

	difficulty := document.Call("getElementById", "difficulty")
	difficulty.Set("value", string(game.AIDifficulty))

	aiToggle := document.Call("getElementById", "aiToggle")
	aiToggle.Set("checked", game.AIEnabled)
	return nil
}

// move is called from the js when a cell is clicked
func move(this js.Value, args []js.Value) any {
	i := args[0].Int()
	game.Move(i)
	js.Global().Call("render")
	return nil
}

// toggleAI is called from the js when the AI toggle is clicked
func toggleAI(this js.Value, args []js.Value) any {
	game.ToggleAI()
	js.Global().Call("render")
	return nil
}

// reset is called from the js when the reset button is clicked
func reset(this js.Value, args []js.Value) any {
	game.Reset()
	js.Global().Call("render")
	return nil
}

func main() {
	c := make(chan struct{})

	js.Global().Set("move", js.FuncOf(move))
	js.Global().Set("render", js.FuncOf(render))
	js.Global().Set("reset", js.FuncOf(reset))
	js.Global().Set("toggleAI", js.FuncOf(toggleAI))
	js.Global().Set("changeDifficulty", js.FuncOf(changeDifficulty))

	js.Global().Call("render")
	<-c
}

func changeDifficulty(this js.Value, args []js.Value) any {
	if len(args) == 0 {
		return nil
	}

	switch args[0].String() {
	case "easy":
		game.ChangeDifficulty(global.EasyDifficulty)
	case "medium":
		game.ChangeDifficulty(global.MediumDifficulty)
	case "hard":
		game.ChangeDifficulty(global.HardDifficulty)
	}

	return nil
}
