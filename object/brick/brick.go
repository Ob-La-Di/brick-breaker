package brick

import (
	"brick-breaker/object"
)

type Brick struct {
	*object.Rectangle
	broken bool
}

func NewBrick(screenHeight, screenWidth int) *Brick {
	return &Brick{object.NewRectangle(50, 100, screenWidth/2-50, screenHeight/2-100), false}
}

func (b Brick) Draw(pixels []byte, screenWidth int) {
	if !b.broken {
		b.Rectangle.Draw(pixels, screenWidth)
	}
}

func (b *Brick) Update()          {}
func (b *Brick) HandleCollision() {}
