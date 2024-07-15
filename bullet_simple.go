package main

import "math/rand"

type simpleBullet struct {
	bullet
}

func (b *simpleBullet) action(next next) {
	b.x = 200
	b.dir = tau/4 + (rand.Float64()-0.5)*tau/8
	b.speed = 30
	next.skip(60)
}
