package main

import (
	"image/color"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	screenWidth, screenHeight = 600., 800.

	vw, vh = 1.0, screenHeight / screenWidth
	sx, sy = vw / 120, vh / 120 // なんとなくいい感じに速度の基準になりそうな値

	circle *ebiten.Image
)

type app struct {
	top   func() (empty, bool)
	stop  func()
	pause bool
}

func main() {
	cp, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	defer cp.Close()
	if err := pprof.StartCPUProfile(cp); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	a := &app{}

	ebiten.SetWindowSize(600, 800)
	if err := ebiten.RunGame(a); err != nil {
		panic(err)
	}

	mp, err := os.Create("mem.pprof")
	if err != nil {
		panic(err)
	}
	defer mp.Close()
	runtime.GC()
	if err := pprof.WriteHeapProfile(mp); err != nil {
		panic(err)
	}
}

func (a *app) Update() error {
	if circle == nil {
		circle = ebiten.NewImage(10, 10)
		vector.DrawFilledCircle(circle, 5, 5, 5, color.White, true)
	}
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
