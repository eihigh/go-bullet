package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	screenWidth, screenHeight = 600., 800.

	vw, vh = 1.0, screenHeight / screenWidth
	sx, sy = vw / 120, vh / 120 // なんとなくいい感じに速度の基準になりそうな値
)

type app struct {
	top   func() (empty, bool)
	stop  func()
	pause bool
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
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		return ebiten.Termination
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		a.pause = !a.pause
	}
	if a.pause {
		return nil
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
