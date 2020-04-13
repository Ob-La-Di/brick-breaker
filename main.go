package main

import (
	"brick-breaker/game"
	"github.com/hajimehoshi/ebiten"
)

const (
	WIDTH  = 1366
	HEIGHT = 768
)

var g *game.Game

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	g.Update()
	g.Draw(screen)
	return nil
}

func main() {
	g = game.New(HEIGHT, WIDTH)
	if err := ebiten.Run(update, WIDTH, HEIGHT, 1, "Hello world!"); err != nil {
		panic(err)
	}
}
