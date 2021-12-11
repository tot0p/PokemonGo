package pok

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type ExitP struct {
	img  *ebiten.Image
	x, y float64
	hb   *HitBox
}

func (e *ExitP) Init(x, y float64) {
	e.x, e.y = x, y
	e.img = LoadImg("data/Tilesets/Poke_Centre_interior.PNG").SubImage(image.Rect(0, 672, 96, 736)).(*ebiten.Image)
	e.hb = &HitBox{x: int(x), y: int(y), w: 96, h: 64}
}

func (e *ExitP) Draw(screen *ebiten.Image, xComp, yComp float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(e.x-xComp, e.y-yComp)
	screen.DrawImage(e.img, op)
}

func (e *ExitP) Ex(h *HitBox) bool {
	return e.hb.CollideCenter(h)
}
