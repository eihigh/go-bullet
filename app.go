package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	screenWidth, screenHeight = 600., 800.
)

type app struct {
	top  func() (empty, bool)
	stop func()
}

func main() {
	a := &app{}
	a.top, a.stop = newCoro(top)

	ebiten.SetWindowSize(600, 800)
	if err := ebiten.RunGame(a); err != nil {
		panic(err)
	}
}

func (a *app) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}
	if _, running := a.top(); !running {
		return ebiten.Termination
	}
	updateBullets()
	return nil
}

func (app) Draw(screen *ebiten.Image) {
	drawBullets(screen)
}

func (app) Layout(ww, wh int) (sw, sh int) {
	return int(screenWidth), int(screenHeight)
}
