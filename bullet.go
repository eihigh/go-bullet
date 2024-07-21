package main

import (
	"fmt"
	"iter"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	tau = math.Pi * 2
)

var (
	coros   []*coro
	bullets []*bullet
)

type empty = struct{}

type coro struct {
	next func() (empty, bool)
	stop func()

	deleted *bool // weak reference だとか context.Context の代用のイメージ

	// - next が false を返した -> それ以上呼ぶ必要はないので stop & delete from coros
	// - deleted が true になった -> bullet が削除されたので stop & delete from coros
}

type bullet struct {
	x, y       float64
	dir, speed float64
	deleted    bool
}

type yield0 func() bool

func (f yield0) skip(n int) {
	for range n {
		if !f() {
			break
		}
	}
}

func pull0(seq0 func(yield0)) (yield1 func() (empty, bool), stop func()) {
	seq1 := func(yield1 func(empty) bool) {
		yield0 := func() bool {
			return yield1(empty{})
		}
		seq0(yield0)
	}
	return iter.Pull(seq1)
}

// go statament のコルーチン版
// deleted は context.Context の代用のイメージ
func spawn(deleted *bool, seq0 func(yield0)) {
	next, stop := pull0(seq0)
	coro := &coro{next, stop, deleted}
	coros = append(coros, coro)
}

func fire(b *bullet, seq0 func(yield0)) {
	spawn(&b.deleted, seq0)
	bullets = append(bullets, b)
}

func updateCoros() {
	j := 0
	// ループ中に要素数が増えるタイプのループなので注意
	for i := 0; i < len(coros); i++ {
		c := coros[i]
		// deleted チェック
		if *c.deleted {
			c.stop() // 忘れずに！
			continue
		}

		// 更新
		if _, running := c.next(); !running {
			c.stop() // 忘れずに！
			continue
		}

		// 生存していれば詰める
		coros[j] = c
		j++
	}

	// 切り詰める
	coros = coros[:j]
}

func updateBullets() {
	j := 0
	for _, b := range bullets {
		// 画面外チェック
		margin := vw / 10
		if b.x < -margin || b.x > vw+margin || b.y < -margin || b.y > vh+margin {
			b.deleted = true
		}
		if b.deleted {
			continue
		}

		// 生存していれば詰める
		bullets[j] = b
		j++

		// 更新
		y, x := math.Sincos(b.dir)
		b.x += x * b.speed
		b.y += y * b.speed
	}

	// 切り詰める
	bullets = bullets[:j]
}

func drawBullets(screen *ebiten.Image) {
	msg := fmt.Sprintf("bullets: %d\ncoros: %d", len(bullets), len(coros))
	ebitenutil.DebugPrint(screen, msg)

	opt := &ebiten.DrawImageOptions{}
	for _, b := range bullets {
		if b.deleted {
			continue
		}

		opt.GeoM.Reset()
		opt.GeoM.Translate(b.x*screenWidth, b.y*screenWidth)
		screen.DrawImage(circle, opt)
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
