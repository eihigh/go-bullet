package main

type simpleBullet struct {
	bullet
}

// コルーチン内でインスタンスを作るとスタックに残り続けて効率がよくない気がする
// なのでfireXXX関数を作ってインスタンスを作る処理を分離した方がいい気がする
func fireSimpleBullet(x, y, dir, speed float64) {
	b := &simpleBullet{} // この関数内でインスタンスを作る
	b.x = x
	b.y = y
	b.dir = dir
	b.speed = speed
	fire(&b.bullet, b.action)
}

func (b *simpleBullet) action(yield yield0) {
}

type accelBullet struct {
	bullet
	fromSpeed, toSpeed float64
	duration           int
}

func fireAccelBullet(x, y, dir, fromSpeed, toSpeed float64, duration int) {
	b := &accelBullet{}
	b.x = x
	b.y = y
	b.dir = dir
	b.fromSpeed = fromSpeed
	b.toSpeed = toSpeed
	b.duration = duration

	fire(&b.bullet, b.action)
}

func (b *accelBullet) action(yield yield0) {
	term(yield, b.duration, &b.speed, b.fromSpeed, b.toSpeed)
	// equivalent:
	// for _, t := range seq(b.duration) {
	// 	b.speed = mix(b.fromSpeed, b.toSpeed, t)
	// 	yield()
	// }
	// b.speed = b.toSpeed
}
