package object

import (
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/mobile/geom"
)

type Point geom.Point
type Velocity Point

type Object interface {
	Draw(*ebiten.Image)
	Update()
	HandleCollision()
}

type Rectangle struct {
	TopLeft Point
	Height  int
	Width   int
}

func NewRectangle(height, width, x, y int) *Rectangle {
	return &Rectangle{Width: width, Height: height, TopLeft: Point{X: geom.Pt(x), Y: geom.Pt(y)}}
}

func (r Rectangle) Draw(pixels []byte, screenWidth int) {
	for i := int(r.TopLeft.X); i < int(r.TopLeft.X)+r.Width; i++ {
		DrawPixel(pixels, Point{X: geom.Pt(i), Y: r.TopLeft.Y}, screenWidth)
		DrawPixel(pixels, Point{X: geom.Pt(i), Y: geom.Pt(int(r.TopLeft.Y) + r.Height)}, screenWidth)
	}

	for i := int(r.TopLeft.Y); i < int(r.TopLeft.Y)+r.Height; i++ {
		DrawPixel(pixels, Point{X: r.TopLeft.X, Y: geom.Pt(i)}, screenWidth)
		DrawPixel(pixels, Point{X: geom.Pt(int(r.TopLeft.X) + r.Width), Y: geom.Pt(i)}, screenWidth)
	}
}

func DrawPixel(pixels []byte, point Point, width int) {
	pixels[(int(point.Y)*width+int(point.X))*4] = 0xff
	pixels[((int(point.Y)*width+int(point.X))*4)+1] = 0xff
	pixels[((int(point.Y)*width+int(point.X))*4)+2] = 0xff
	pixels[((int(point.Y)*width+int(point.X))*4)+3] = 0xff
}
