package ball

import (
	"brick-breaker/object"
	"golang.org/x/mobile/geom"
)

type Ball struct {
	Center   object.Point
	Radius   int
	Velocity object.Velocity
}

func New(height, width int) *Ball {
	return &Ball{Velocity: object.Velocity{X: 5, Y: 10}, Radius: 12, Center: object.Point{X: geom.Pt(width / 2), Y: geom.Pt(height / 2)}}
}

func (b Ball) Draw(pixels []byte, screenWidth int) {
	x, y, dx, dy := b.Radius-1, 0, 1, 1
	err := dx - (b.Radius * 2)

	for x > y {
		object.DrawPixel(pixels, object.Point{X: b.Center.X + geom.Pt(x), Y: b.Center.Y + geom.Pt(y)}, screenWidth)
		object.DrawPixel(pixels, object.Point{X: b.Center.X + geom.Pt(x), Y: b.Center.Y + geom.Pt(y)}, screenWidth)
		object.DrawPixel(pixels, object.Point{X: b.Center.X + geom.Pt(y), Y: b.Center.Y + geom.Pt(x)}, screenWidth)
		object.DrawPixel(pixels, object.Point{X: b.Center.X - geom.Pt(y), Y: b.Center.Y + geom.Pt(x)}, screenWidth)
		object.DrawPixel(pixels, object.Point{X: b.Center.X - geom.Pt(x), Y: b.Center.Y + geom.Pt(y)}, screenWidth)
		object.DrawPixel(pixels, object.Point{X: b.Center.X - geom.Pt(x), Y: b.Center.Y - geom.Pt(y)}, screenWidth)
		object.DrawPixel(pixels, object.Point{X: b.Center.X - geom.Pt(y), Y: b.Center.Y - geom.Pt(x)}, screenWidth)
		object.DrawPixel(pixels, object.Point{X: b.Center.X + geom.Pt(y), Y: b.Center.Y - geom.Pt(x)}, screenWidth)
		object.DrawPixel(pixels, object.Point{X: b.Center.X + geom.Pt(x), Y: b.Center.Y - geom.Pt(y)}, screenWidth)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}
		if err > 0 {
			x--
			dx += 2
			err += dx - (b.Radius * 2)
		}
	}
}

func (b *Ball) Update() {
	b.Center.X += b.Velocity.X
	b.Center.Y += b.Velocity.Y
}
func (b *Ball) HandleCollision() {}
