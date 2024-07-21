package main

import (
	"log"

	"github.com/arthurasanaliev/conways-game-of-life-go/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()

	ebiten.SetWindowSize(game.WINDOW_WIDTH, game.WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Conway's Game of Life")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
