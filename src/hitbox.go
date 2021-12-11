package pok

type HitBox struct {
	x, y, w, h int
}

func (h *HitBox) CollideCenter(h2 *HitBox) bool {
	x, y := h2.x+h2.w/2, h2.y+h2.h/2
	return h.x <= x && x <= h.x+h.w && h.y <= y && y <= h.y+h.h
}

func (h *HitBox) Collide(h2 *HitBox) bool {
	var t [4][2]int
	t[0][0], t[0][1] = h2.x, h2.y
	t[1][0], t[1][1] = h2.x+h2.w, h2.y
	t[2][0], t[2][1] = h2.x+h2.w, h2.y+h2.h
	t[3][0], t[3][1] = h2.x, h2.y+h2.h
	for _, i := range t {
		if h.x <= i[0] && i[0] <= h.x+h.w && h.y <= i[1] && i[1] <= h.y+h.h {
			return true
		}
	}
	return false
}

func (h *HitBox) Update(x, y float64) {
	h.x, h.y = int(x), int(y)
}
