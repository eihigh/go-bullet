package main

type simpleBullet struct {
	bullet
}

func (b *simpleBullet) action(next next) {
	next.skip(20)
	fire(&b.bullet, b.action)
}
