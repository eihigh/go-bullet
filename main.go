package main

// 事実上のメイン関数
func top(next next) {
	for range 120 {
		fireArrowBullet(vw/2, vh/13, tau/4)
		next.skip(30)
	}
}
