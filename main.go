package main

// 事実上のメイン関数
func top(next next) {
	for range 120 {
		b := &simpleBullet{}
		fire(&b.bullet, b.action)
		next()
		next()
	}
}
