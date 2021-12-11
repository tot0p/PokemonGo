package pok

import (
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var Bag = [2]int{10, 0} // 1 : Pokeball , 2: Soin

type Player struct {
	x     float64
	y     float64
	xIn   float64
	yIn   float64
	img   *ebiten.Image
	d     bool
	g     bool
	b     bool
	h     bool
	lastD int
	pok   [6]*Pokemon
	mp    *MapPok
	hb    *HitBox
}

func (p *Player) Draw(screen *ebiten.Image, count *int, In bool) {
	op := &ebiten.DrawImageOptions{}
	var x, y float64
	if In {
		x, y = p.GetPosIn()
		op.GeoM.Translate(x+16, y+16)
	} else {
		x, y = p.GetPox()
		op.GeoM.Translate(x-p.mp.xComp, y-p.mp.yComp)
	}
	i := (*count / 10) % 4
	sx, sy := 0+i*32, 0
	if p.d {
		sy = 2 * 48
		p.lastD = sy
	} else if p.g {
		sy = 1 * 48
		p.lastD = sy
	} else if p.h {
		sy = 3 * 48
		p.lastD = sy
	} else if p.b {
		p.lastD = sy
	}
	if !p.h && !p.g && !p.d && !p.b {
		sx = 0
		i = 0
		sy = p.lastD
	}
	screen.DrawImage(p.GetImg().SubImage(image.Rect(sx, sy, sx+32, sy+48)).(*ebiten.Image), op)
}

func (p *Player) AddPok(pok *Pokemon) {
	cursor := 0
	for p.pok[cursor] != nil {
		cursor++
	}
	p.pok[cursor] = pok
}

func (p *Player) Init(img *ebiten.Image, mp *MapPok) {
	p.img = img
	p.mp = mp
	p.x = 0
	p.y = 64
	p.xIn = 0
	p.yIn = 0
	p.hb = &HitBox{x: int(p.x), y: int(p.y), w: 32, h: 48}
	p.g, p.d, p.b, p.h = false, false, false, false
}

//in
func (p *Player) SetPosIn(x, y float64) {
	p.xIn, p.yIn = x, y
}

func (p *Player) GetPosIn() (float64, float64) {
	return p.xIn, p.yIn
}

func (p *Player) GetPok1() *Pokemon {
	return p.pok[0]
}

func (p *Player) GetPok() *[6]*Pokemon {
	return &p.pok
}

func (p *Player) SetPos(x, y float64) {
	p.x, p.y = x, y
}

func (p *Player) GetPox() (float64, float64) {
	return p.x, p.y
}

func (p *Player) GetImg() *ebiten.Image {
	return p.img
}

func (p *Player) GetHB() *HitBox {
	p.hb.Update(p.x, p.y)
	return p.hb
}
