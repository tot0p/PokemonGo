package pok

import (
	"fmt"
	"image"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	tempPC, _ = opentype.Parse(fonts.MPlus1pRegular_ttf)
	Flv, _    = opentype.NewFace(tempPC, &opentype.FaceOptions{
		Size:    24,
		DPI:     45,
		Hinting: font.HintingFull,
	})
	FnamePM, _ = opentype.NewFace(tempPC, &opentype.FaceOptions{
		Size:    12,
		DPI:     72,
		Hinting: font.HintingFull,
	})
)

type PokControl struct {
	icons                                    []*ebiten.Image
	back, box, blank, firstBox, hp1, hp2, lv *ebiten.Image
	pok                                      *[6]*Pokemon
	nbpok, Pan, choose, mode                 int
}

func (p *PokControl) Init(pok *[6]*Pokemon, mode ...int) {
	p.pok = pok
	if len(mode) > 0 {
		p.mode = mode[0]
	} else {
		p.mode = 0
	}
	for _, i := range pok {
		if i != nil {
			p.nbpok += 1
		}
	}
	for i := 0; i < p.nbpok; i++ {
		t := LoadImg("data/pokemon/icon/icon" + pok[i].GetId() + ".png")
		p.icons = append(p.icons, t)
	}
	p.back = LoadImg("data/GUI/PokemonControl/bg.PNG")
	p.box = LoadImg("data/GUI/PokemonControl/panel_rect.png")
	p.blank = LoadImg("data/GUI/PokemonControl/panel_blank.png")
	p.firstBox = LoadImg("data/GUI/PokemonControl/panel_round.png")
	p.hp1 = LoadImg("data/GUI/PokemonControl/overlay_hp_back.png")
	p.hp2 = LoadImg("data/GUI/battle/overlay_hp.png")
	p.lv = LoadImg("data/GUI/PokemonControl/overlay_lv.png")
}

func (p *PokControl) Draw(screen *ebiten.Image, count int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(p.back, op)
	screen.DrawImage(p.firstBox, op)
	op.GeoM.Translate(20, 0)
	if count%64 == 0 {
		if p.Pan > 0 {
			p.Pan -= 64
		} else {
			p.Pan += 64
		}
	}
	screen.DrawImage(p.icons[0].SubImage(image.Rect(0+p.Pan, 0, 64+p.Pan, 64)).(*ebiten.Image), op)
	op.GeoM.Translate(64, 50)
	screen.DrawImage(p.hp1, op)
	text.Draw(screen, p.pok[0].name, FnamePM, 85, 35, color.Black)
	op.GeoM.Translate(32, 4)
	p.drawHp(screen, p.pok[0].GetPv(), p.pok[0].pvmax, op)
	op.GeoM.Translate(-80, 10)
	screen.DrawImage(p.lv, op)
	p.drawlv(screen, 65, 75, 0)
	op.GeoM.Reset()
	op.GeoM.Translate(258, 16)
	x, y := 258, 16
	t := true
	for k, i := range p.pok[1:] {
		if i == nil {
			screen.DrawImage(p.blank, op)
		} else {
			screen.DrawImage(p.box, op)
			op.GeoM.Translate(20, 0)
			screen.DrawImage(p.icons[k+1].SubImage(image.Rect(0+p.Pan, 0, 64+p.Pan, 64)).(*ebiten.Image), op)
			op.GeoM.Translate(64, 50)
			screen.DrawImage(p.hp1, op)
			op.GeoM.Translate(32, 4)
			p.drawHp(screen, p.pok[k+1].GetPv(), p.pok[k+1].pvmax, op)
			op.GeoM.Translate(-80, 10)
			screen.DrawImage(p.lv, op)
			p.drawlv(screen, 65+x, 75+y, 0)
			op.GeoM.Translate(-36, -64)
			text.Draw(screen, i.name, FnamePM, 85+x, 35+y, color.Black)
		}
		if t {
			op.GeoM.Translate(-256, 80)
			t = false
			x -= 256
			y += 80
		} else {
			op.GeoM.Translate(256, 16)
			t = true
			x += 256
			y += 16
		}
	}
}

func (p *PokControl) drawlv(screen *ebiten.Image, x, y, i int) {
	t := strconv.Itoa(p.pok[i].lv)
	if p.pok[i].lv < 10 {
		t = "0" + t
	}
	text.Draw(screen, t, Flv, x, y, color.Black)
}

func (p *PokControl) drawHp(screen *ebiten.Image, pv, pvmax int, op *ebiten.DrawImageOptions) {
	t := (float64(pv) / float64(pvmax))
	var g int
	if pv < 0 {
		t = 0
		g = 3
	} else {
		if t <= 0.25 {
			g = 3
		} else if t <= 0.5 {
			g = 2
		} else {
			g = 1
		}
	}
	screen.DrawImage(p.hp2.SubImage(image.Rect(0, 2+6*(g-1), 2+int(96*t), 6*g)).(*ebiten.Image), op)
}

func (p *PokControl) Reset() {
	p.icons = nil
	p.nbpok = 0
}

func (p *PokControl) Event(pok ...*Pokemon) *Pokemon {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if p.pok[0] != nil && 2 <= x && x <= 253 && 8 <= y && y <= 95 {
			fmt.Println(p.pok[0].name, " ", 0)
			if p.choose == -1 {
				p.choose = 0
			} else if p.choose != 0 {
				p.swap(0)
				p.choose = -1
			}
		} else if p.pok[1] != nil && 258 <= x && x <= 520 && 24 <= y && y <= 110 {
			fmt.Println(p.pok[1].name, " ", 1)
			if p.choose == -1 {
				p.choose = 1
			} else if p.choose != 1 {
				p.swap(1)
				p.choose = -1
			}
		} else if p.pok[2] != nil && 2 <= x && x <= 253 && 104 <= y && y <= 190 {
			fmt.Println(p.pok[2].name, " ", 2)
			if p.choose == -1 {
				p.choose = 2
			} else if p.choose != 2 {
				p.swap(2)
				p.choose = -1
			}
		} else if p.pok[3] != nil && 258 <= x && x <= 520 && 120 <= y && y <= 207 {
			fmt.Println(p.pok[3].name, " ", 3)
			if p.choose == -1 {
				p.choose = 3
			} else if p.choose != 3 {
				p.swap(3)
				p.choose = -1
			}
		} else if p.pok[4] != nil && 2 <= x && x <= 253 && 200 <= y && y <= 286 {
			fmt.Println(p.pok[4].name, " ", 4)
			if p.choose == -1 {
				p.choose = 4
			} else if p.choose != 4 {
				p.swap(4)
				p.choose = -1
			}
		} else if p.pok[5] != nil && 258 <= x && x <= 520 && 216 <= y && y <= 302 {
			fmt.Println(p.pok[5].name, " ", 5)
			if p.choose == -1 {
				p.choose = 5
			} else if p.choose != 5 {
				p.swap(5)
				p.choose = -1
			}
		}
	}
	return nil
}

func (p *PokControl) swap(x int) {
	*p.pok[p.choose], *p.pok[x] = *p.pok[x], *p.pok[p.choose]
	p.icons[p.choose], p.icons[x] = p.icons[x], p.icons[p.choose]
}
