package main

// 事実上のメイン関数
func top(next next) {
	for range 120 {
		fireArrowBullet(vw/2, vh/13, tau/4) // 画面中央上辺りから、下に向けて矢印弾幕を発射
		next.skip(30)                       // 30フレーム待つ
	}
}
