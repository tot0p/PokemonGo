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
	Fname   font.Face
	Flevel  font.Face
	PokCapt *Pokemon
)

type Battle struct {
	pok1        *Pokemon
	pPok        *[6]*Pokemon
	pok2        *Pokemon
	imgB        *ebiten.Image
	imgBasePok1 *ebiten.Image
	imgBasePok2 *ebiten.Image
	imgGuiHp    *ebiten.Image
	imgGuiXp    *ebiten.Image
	imgLv       *ebiten.Image
	InteB       *InteBattle
	playerTour  bool
	cursor      int
}

func (b *Battle) Init(pok *[6]*Pokemon) {
	b.pok2 = &Pokemon{}
	t := []string{"Hericendre", "Carapuce", "Bulbizarre", "Salameche"}
	g := RandListString(t)
	b.pok2.Init(g)
	b.playerTour = true
	b.pPok = pok
	b.cursor = 0
	for pok[b.cursor].pv <= 0 && b.cursor < 6 {
		b.cursor++
	}
	b.pok1 = pok[b.cursor]
	b.InteB = &InteBattle{
		img:  LoadImg("data/GUI/battle/overlay_fight.png"),
		x:    0,
		y:    288,
		att1: b.pok1.att1,
		att2: b.pok1.att2,
	}
	b.imgB = LoadImg("data/battlebacks/battlebgForest.png")
	b.imgBasePok1 = LoadImg("data/battlebacks/playerbaseForestGrass.png")
	b.imgBasePok2 = LoadImg("data/battlebacks/enemybaseForestGrass.png")
	b.imgGuiHp = LoadImg("data/GUI/battle/overlay_hp.png")
	b.imgGuiXp = LoadImg("data/GUI/battle/overlay_exp.png")
	b.imgLv = LoadImg("data/GUI/battle/overlay_lv.png")
	tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)
	Fname, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    22,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	Flevel, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    12,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}

func (b *Battle) Draw(screen *ebiten.Image) {
	//Back
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(b.imgB, op)
	b.InteB.Draw(screen)
	//Pok En
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(256, 100)
	screen.DrawImage(b.imgBasePok2, op)
	op.GeoM.Translate(48, -85) //bon pour les petits mais pas les gros #dracofeu
	screen.DrawImage(LoadImg("data/pokemon/img/"+b.pok2.imgName+".png"), op)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 50)
	img := LoadImg("data/GUI/battle/databox_normal_foe.png")
	b.drawDataBox(img, true, b.pok2.lv, b.pok2.pv, b.pok2.pvmax, b.pok2.name, 0)
	screen.DrawImage(img, op)
	//Pok Player
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-100, 224)
	screen.DrawImage(b.imgBasePok1, op)
	op.GeoM.Translate(176, -96)
	screen.DrawImage(LoadImg("data/pokemon/img/"+b.pok1.imgName+"b.png"), op)
	//hp Player
	img = LoadImg("data/GUI/battle/databox_normal.png")
	b.drawDataBox(img, false, b.pok1.lv, b.pok1.pv, b.pok1.pvmax, b.pok1.name, b.pok1.exp)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(252, 195)
	screen.DrawImage(img, op)
}

func (b *Battle) drawDataBox(screen *ebiten.Image, en bool, lv int, hp int, hpmax int, name string, xp int) {
	var strlv string
	op := &ebiten.DrawImageOptions{}
	if en {
		text.Draw(screen, name, Fname, 5, 25, color.Black)
		op.GeoM.Translate(118, 40)
	} else {
		text.Draw(screen, name, Fname, 40, 25, color.Black)
		op.GeoM.Translate(40, 76)
		screen.DrawImage(b.imgGuiXp.SubImage(image.Rect(0, 0, int(192*xp/100), 4)).(*ebiten.Image), op)
		text.Draw(screen, strconv.Itoa(hp)+"/"+strconv.Itoa(hpmax), Flevel, 175, 65, color.Black)
		op.GeoM.Reset()
		op.GeoM.Translate(136, 40)
	}
	t := (float64(hp) / float64(hpmax))
	var g int
	if hp < 0 {
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
	screen.DrawImage(b.imgGuiHp.SubImage(image.Rect(0, 6*(g-1), int(96*t), 6*g)).(*ebiten.Image), op)
	op.GeoM.Reset()
	op.GeoM.Translate(180, 10)
	screen.DrawImage(b.imgLv, op)
	if lv < 10 {
		strlv = "0" + strconv.Itoa(lv)
	} else {
		strlv = strconv.Itoa(lv)
	}
	text.Draw(screen, strlv, Flevel, 202, 22, color.Black)
}

func (b *Battle) Event(count int) {
	b.InteB.Event(count)
	if AttackValue != 0 {
		b.pok2.SetPv(b.pok2.GetPv() - AttackValue)
		AttackValue = 0
		b.pok1.SetPv(b.pok1.GetPv() - int(float64(b.pok2.att1)*RandForAtt1())/6)
	} else if Soin != 0 {
		t := b.pok1.GetPv()
		if b.pok1.pvmax-t < Soin {
			b.pok1.SetPv(b.pok1.pvmax)
		} else {
			b.pok1.SetPv(t + Soin)
		}
		Soin = 0
	} else if Capt {
		if RandPourInt(b.pok2.tc) {
			Capt = false
			EscapeBattle = true
			PokCapt = b.pok2
		}
	} else if b.pok2.GetPv() <= 0 {
		EscapeBattle = true
		b.pok1.AddExp(1000)
		fmt.Println("Win")
	} else if b.pok1.GetPv() <= 0 {
		if b.cursor < 5 {
			b.cursor++
			if t := b.pPok[b.cursor]; t != nil {
				if t.GetPv() >= 0 {
					b.pok1 = b.pPok[b.cursor]
					b.InteB.SetPokAtt(b.pok1.att1, b.pok1.att2)
				}
			} else {
				EscapeBattle = true
				fmt.Println("Lose")
			}
		} else {
			EscapeBattle = true
			fmt.Println("Lose")
		}
	}
}
