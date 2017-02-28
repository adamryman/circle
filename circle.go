package circle

import (
	"image"
	"image/color"
)

func Simple(x, y, radius int) *Circle {
	return &Circle{
		Point: image.Point{
			X: x,
			Y: y,
		},
		Radius: radius,
	}
}

// Circle implements image.Image
// godoc image Image
type Circle struct {
	Point  image.Point
	Radius int
}

func (c *Circle) ColorModel() color.Model {
	return color.GrayModel
}

func (c *Circle) Bounds() image.Rectangle {
	return image.Rect(
		c.Point.X-c.Radius, // x0
		c.Point.Y-c.Radius, // y0
		c.Point.X+c.Radius, // x1
		c.Point.Y+c.Radius, // y1
	)
}

func (c *Circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.Point.X)+0.5, float64(y-c.Point.Y)+0.5, float64(c.Radius)
	if xx*xx+yy*yy < rr*rr {
		return color.Black
	}
	return color.White
}

// Sector coming soon
type Sector struct {
	Θ1 float64
	Θ2 float64
}
