package game

import (
	"brick-breaker/object"
	"brick-breaker/object/ball"
	"brick-breaker/object/brick"
	"brick-breaker/object/paddle"
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/mobile/geom"
	"math"
)

const (
	TOP_WALL int = iota
	RIGHT_WALL
	BOTTOM_WALL
	LEFT_WALL
)

type Game struct {
	ball          *ball.Ball
	paddle        *paddle.Paddle
	bricks        []*brick.Brick
	pixels        []byte
	width, height int
}

func New(screenHeight, screenWidth int) *Game {
	return &Game{
		ball:   ball.New(screenHeight, screenWidth),
		paddle: paddle.NewPaddle(screenHeight, screenWidth),
		bricks: []*brick.Brick{brick.NewBrick(screenHeight, screenWidth)},
		pixels: make([]byte, screenHeight*screenWidth*4),
		width:  screenWidth,
		height: screenHeight,
	}
}

func (g *Game) Update() {
	g.handleCollision()
	g.ball.Update()
	g.paddle.Update(g.width)
}

func clampBallAngle(angle geom.Pt) geom.Pt {
	const (
		angle150 = 2 * math.Pi * 150 / 360
		angle30  = 2 * math.Pi * 30 / 360
		angle90  = 2 * math.Pi * 90 / 360
		angle75  = 2 * math.Pi * 75 / 360
		angle105 = 2 * math.Pi * 105 / 360
	)
	if angle > angle150 {
		return angle150
	}
	if angle < angle30 {
		return angle30
	}
	if angle >= angle90 && angle < angle75 {
		return angle75
	}
	if angle >= angle90 && angle < angle105 {
		return angle105
	}

	return angle
}

func (g Game) isBallCollidingWithHorizontalWall() bool {
	return int(g.ball.Center.Y)+g.ball.Radius+int(g.ball.Velocity.Y) >= g.height ||
		int(g.ball.Center.Y)+g.ball.Radius+int(g.ball.Velocity.Y) <= 50
}

func (g Game) isBallCollidingWithVerticalWall() bool {
	return int(g.ball.Center.X)+g.ball.Radius+int(g.ball.Velocity.X) >= g.width ||
		int(g.ball.Center.X)+g.ball.Radius+int(g.ball.Velocity.X) <= 50
}

func (g *Game) handleCollision() {
	if circleRectangleCollision(&g.paddle.Rectangle, g.ball) {
		hitFactor := (g.ball.Center.X - g.paddle.TopLeft.X + geom.Pt(g.paddle.Width)/2) / geom.Pt(g.paddle.Width)
		vectorLength := math.Sqrt(float64(1 + hitFactor*hitFactor))
		g.ball.Velocity.X *= hitFactor / geom.Pt(vectorLength)
		g.ball.Velocity.Y *= 1/geom.Pt(vectorLength)
	} else if g.isBallCollidingWithHorizontalWall() {
		g.ball.Velocity.Y *= -1
	} else if g.isBallCollidingWithVerticalWall() {
		g.ball.Velocity.X *= -1
	}

	for _, b := range g.bricks {
		if circleRectangleCollision(b.Rectangle, g.ball) {
			g.ball.Velocity.X *= -1
			g.ball.Velocity.Y *= -1
		}
	}
}

func circleRectangleCollision(rectangle *object.Rectangle, ball *ball.Ball) bool {
	nextBallX, nextBallY := ball.Center.X+ball.Velocity.X, ball.Center.Y+ball.Velocity.Y

	testX, testY := nextBallX, nextBallY

	if nextBallX < rectangle.TopLeft.X {
		testX = rectangle.TopLeft.X
	} else if nextBallX > rectangle.TopLeft.X+geom.Pt(rectangle.Width) {
		testX = rectangle.TopLeft.X + geom.Pt(rectangle.Width)
	}

	if nextBallY < rectangle.TopLeft.Y {
		testY = rectangle.TopLeft.Y
	} else if nextBallY > rectangle.TopLeft.Y+geom.Pt(rectangle.Height) {
		testY = rectangle.TopLeft.Y + geom.Pt(rectangle.Height)
	}

	distX, distY := nextBallX-testX, nextBallY-testY

	distance := int(math.Sqrt(float64((distX * distX) + (distY * distY))))

	if distance <= ball.Radius {
		return true
	}

	return false
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < g.height*g.width*4; i++ {
		g.pixels[i] = 0
	}
	for _, brick := range g.bricks {
		brick.Draw(g.pixels, g.width)
	}

	g.ball.Draw(g.pixels, g.width)

	g.paddle.Draw(g.pixels, g.width)

	screen.ReplacePixels(g.pixels)
}
