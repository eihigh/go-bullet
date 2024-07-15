package main

import (
	"fmt"
	"image/color"
	"iter"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	bullets []*bullet
)

type empty = struct{}

type bullet struct {
	next       func() (empty, bool) // fireした後でセットされる
	stop       func()
	x, y       float64
	dir, speed float64
	deleted    bool
}

type next func() bool

func (f next) skip(n int) {
	for range n {
		if !f() {
			break
		}
	}
}

func newCoro(f func(next)) (next func() (empty, bool), stop func()) {
	seq := func(yield func(empty) bool) {
		next := func() bool {
			return yield(empty{})
		}
		f(next)
	}
	return iter.Pull(seq)
}

func fire(b *bullet, action func(next)) {
	b.next, b.stop = newCoro(action)
	bullets = append(bullets, b)
}

func updateBullets() {
	j := 0
	for _, b := range bullets {
		// 画面外チェック
		margin := 100.
		if b.x < -margin || b.x > screenWidth+margin || b.y < -margin || b.y > screenHeight+margin {
			b.deleted = true
		}
		if b.deleted {
			b.stop() // 忘れずに！
			continue
		}

		// 生存していれば詰める
		bullets[j] = b
		j++

		// 更新
		_, running := b.next()
		if !running {
			b.deleted = true
		}
	}

	// 切り詰める
	bullets = bullets[:j]
}

func drawBullets(screen *ebiten.Image) {
	msg := fmt.Sprint("bullets: ", len(bullets))
	ebitenutil.DebugPrint(screen, msg)

	for _, b := range bullets {
		if b.deleted {
			continue
		}

		vector.DrawFilledCircle(screen, float32(b.x), float32(b.y), 5, color.White, false)
	}
}

func seq(n int) iter.Seq2[int, float64] {
	return func(yield func(int, float64) bool) {
		for i := range n {
			if !yield(i, float64(i)/float64(n)) {
				break
			}
		}
	}
}

func mix(a, b, t float64) float64 {
	return a + (b-a)*t
}
