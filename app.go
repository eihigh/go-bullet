package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	screenWidth, screenHeight = 600., 800.

	vw, vh = 1.0, screenHeight / screenWidth
	sx, sy = vw / 120, vh / 120
)

type app struct {
	top  func() (empty, bool)
	stop func()
}

func main() {
	a := &app{}

	ebiten.SetWindowSize(600, 800)
	if err := ebiten.RunGame(a); err != nil {
		panic(err)
	}
}

func (a *app) Update() error {
	if a.top == nil {
		a.top, a.stop = newCoro(top)
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}
	if _, running := a.top(); !running {
		// return ebiten.Termination
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
