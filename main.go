package main

// 事実上のメイン関数
func top(yield yield0) {
	// 渦弾幕
	d := false
	spawn(&d, func(yield yield0) {
		dir := 0.0
		for range 180 {
			// 画面中央から、7方向に渦を巻くように弾幕を発射
			for _, t := range seq(7) {
				fireSimpleBullet(vw/2, vh/2, dir+tau*t, sx/2)
			}
			// 発射方向をずらして次のフレームへ
			dir -= tau / 110
			yield()
		}
	})

	// 矢印弾幕
	for range 120 {
		fireArrowBullet(vw/2, vh/13, tau/4+tau/30) // 画面中央上辺りから、左下に向けて矢印弾幕を発射
		yield.skip(30)                             // 30フレーム待つ
		fireArrowBullet(vw/2, vh/13, tau/4-tau/30) // 画面中央上辺りから、右下に向けて矢印弾幕を発射
		yield.skip(30)                             // 30フレーム待つ
	}
}
