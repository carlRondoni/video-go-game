package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game estructura base del juego
type Game struct{}

// Update se llama cada frame, pero aquí no hacemos nada aún
func (g *Game) Update() error {
	return nil
}

// Draw dibuja en la pantalla
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 50, G: 50, B: 50, A: 255}) // fondo gris
	ebitenutil.DebugPrint(screen, "Hello World!")
}

// Layout define el tamaño lógico de la pantalla
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240 // resolución lógica
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello World")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
