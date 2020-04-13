package paddle

import (
	"brick-breaker/object"
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/mobile/geom"
)

type Paddle struct {
	object.Rectangle
}

func NewPaddle(screenHeight, screenWidth int) *Paddle {
	return &Paddle{*object.NewRectangle(20, 100, screenWidth/2, screenHeight-40)}
}

func (p Paddle) Draw(pixels []byte, screenWidth int) {
	p.Rectangle.Draw(pixels, screenWidth)
}

func (p *Paddle) Update(screenWidth int) {
	cursorX, _ := ebiten.CursorPosition()
	x := cursorX
	if cursorX < 0 {
		x = 0
	} else if x+p.Rectangle.Width > screenWidth {
		x = screenWidth - p.Rectangle.Width - 1
	}

	p.Rectangle.TopLeft.X = geom.Pt(x)

}
func (p *Paddle) HandleCollision() {}
