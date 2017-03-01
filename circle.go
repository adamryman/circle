package circle

import (
	"image"
	"image/color"
	"math"
	//. "github.com/y0ssar1an/q"
)

func Simple(x, y, radius int) *Circle {
	return &Circle{
		Point:  image.Pt(x, y),
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
	return color.RGBAModel
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
	Circle
	Θ1 float64
	Θ2 float64
}

func (s *Sector) At(x, y int) color.Color {
	// Center around origin
	xx, yy, rr := float64(x-s.Point.X), -float64(y-s.Point.Y), float64(s.Radius)
	if crr := xx*xx + yy*yy; crr < rr*rr {
		theta := math.Atan2(yy, xx)
		if theta < 0 {
			theta = theta + 2*math.Pi
		}
		if theta >= s.Θ1 && theta < s.Θ2 {
			return Red{}

		}
		return color.Black
	}
	return color.White
}

type Red struct{}

func (_ Red) RGBA() (r, g, b, a uint32) {
	return 0xffff, 0, 0, 0xffff
}
