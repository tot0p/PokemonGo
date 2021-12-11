package main

import (
	"fmt"
	"image"
	"log"

	"pok"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	m        = pok.MapPok{}
	p        = pok.Player{}
	b        = pok.Battle{}
	i        = pok.Intro{}
	inbat    = false
	start    = true
	inPok    = false
	settings = pok.LoadJson("data/settings.json")
)

const (
	screenWidth  = 512
	screenHeight = 384
)

type Game struct {
	count   int
	lasPosX float64
	lasPosY float64
}

func (g *Game) Update() error {
	g.count++
	pok.Close()
	if !pok.PokManager {
		if start {
			start = !i.Choice(p.GetPok())
		} else if !inbat {
			inPok = pok.InPokCenter
			if inPok {
				inbat = true
				p.SetPosIn(226, 271)
			}
			p.Control()
			x, y := p.GetPox()
			//xComp, yComp := m.GetComp()
			if g.count%5 == 0 {
				//p.Debug()
				if x != g.lasPosX && y != g.lasPosY {
					if m.OnTallGrass(x+16, y+24) {
						g.lasPosX, g.lasPosY = x, y
						if pok.RandPourFloat(pok.ChanceAp(10)) {
							inbat = true
							b.Init(p.GetPok())
							p.ResetDep()
						}
					}
				}
			}
			pok.EventMainEnterPokMan(p.GetPok())
		} else if inPok {
			p.ControlIn()
			pok.EventMainEnterPokMan(p.GetPok())
			if pok.ExitIn {
				inbat = false
				pok.ExitIn = false
				pok.InPokCenter = false
			}
		} else {
			b.Event(g.count)
			if pok.EscapeBattle {
				pok.EscapeBattle = false
				inbat = false
				if pok.PokCapt != nil {
					p.AddPok(pok.PokCapt)
					pok.PokCapt = nil
					fmt.Println(*p.GetPok()) //Montre Pok Capturé
				}
				b = pok.Battle{}
			}
		}
	} else {
		p.ResetDep()
		pok.EventMainExitPokMan()
		pok.PokManagerView.Event()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if pok.PokManager {
		pok.PokManagerView.Draw(screen, g.count)
	} else if start {
		i.Draw(screen)
	} else if !inbat {
		m.Draw(screen)
		p.Draw(screen, &g.count, false)
	} else if inPok {
		m.DrawInPokCenter(screen)
		p.Draw(screen, &g.count, true)
	} else {
		b.Draw(screen)
	}
	pok.DebugMod(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 512, 384
}

func main() {
	p.Init(pok.LoadImg("data/Characters/boy_run.png"), &m)
	m.Init()
	i.Init()
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Pokémon Go")
	ebiten.SetMaxTPS(60)
	ebiten.SetWindowResizable(true)
	ebiten.SetFullscreen(settings["fs"].(bool))
	ebiten.SetWindowIcon([]image.Image{pok.LoadImgImage("data/logo.png")})
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
