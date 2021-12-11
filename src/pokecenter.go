package pok

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bank struct {
	x, y int
	img  *ebiten.Image
}

type NPC struct {
	x, y int
	img  *ebiten.Image
}

type Pokecenter struct {
	pc   *Bank
	NPC  *NPC
	Exit *ExitP
	img  *ebiten.Image
	mp   []tile
	x, y float64
}

func (p *Pokecenter) DrawIn(screen *ebiten.Image) {
	for _, i := range p.mp {
		i.Draw(screen, -16, -16)
	}
	p.Exit.Draw(screen, -16, -16)
}

func (p *Pokecenter) EventExit(h *HitBox) bool {
	return p.Exit.Ex(h)
}

func (p *Pokecenter) Init(x, y float64) {
	p.x, p.y = x, y
	p.pc = &Bank{x: 0, y: 0, img: LoadImg("data/Tilesets/Poke_Centre_interior.PNG").SubImage(image.Rect(1280, 200, 256, 1376)).(*ebiten.Image)}
	p.NPC = &NPC{x: 0, y: 0}
	p.Exit = &ExitP{}
	p.Exit.Init(192, 320)
	p.img = LoadImg("data/Tilesets/Outside.png").SubImage(image.Rect(0, 10432, 160, 10592)).(*ebiten.Image)
	p.loadPokeCenter()
}

func (p *Pokecenter) loadPokeCenter() {
	temp := LoadFileMap("data/Map/PokeCenter.txt")
	t2 := func(i int) bool { return i < 32 }
	var img *ebiten.Image
	tileset := LoadImg("data/Tilesets/Poke_Centre_interior.PNG")
	for y, k := range temp {
		for x, i := range k {
			t := tile{}
			sx, sy := 32*(i%8), 32*(i/8)
			img = tileset.SubImage(image.Rect(sx, sy, sx+32, sy+32)).(*ebiten.Image)
			t.Init(float64(x), float64(y), i, img)
			t.SetSolidCond(t2)
			p.mp = append(p.mp, t)
		}
	}
}

func (p *Pokecenter) CheckCollideOut(h *HitBox) bool {
	h2 := HitBox{x: int(p.x), y: int(p.y + 25), w: 160, h: 95}
	return h2.Collide(h)
}

func (p *Pokecenter) Enter(h *HitBox) bool {
	h2 := HitBox{x: int(p.x + 64), y: int(p.y + 96), w: 32, h: 32}
	return h2.Collide(h)
}

func (p *Pokecenter) DrawOutside(screen *ebiten.Image, xComp, yComp float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.x-xComp, p.y-yComp)
	screen.DrawImage(p.img, op)
}
