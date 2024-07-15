package main

func fireArrowBullet(x, y, dir float64) {
	fireAccelBullet(x, y, dir, sy/3, sy/1.02, 180)
	fire2way := func(dd, speed float64) {
		dd /= 1.5
		fireAccelBullet(x, y, dir+dd, sy/3, speed, 180)
		fireAccelBullet(x, y, dir-dd, sy/3, speed, 180)
	}
	fire2way(tau/120, sy/1.1)
	fire2way(tau/60, sy/1.2)
	fire2way(tau/40, sy/1.3)
	fire2way(tau/30, sy/1.4)
	fire2way(tau/24, sy/1.5)
}
