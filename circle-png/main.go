package main

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
	"strconv"

	"github.com/adamryman/circle"
	flag "github.com/spf13/pflag"

	. "github.com/y0ssar1an/q"
)

var (
	x = flag.IntP("x", "y", 0, "x point")
	y = flag.IntP("y", "x", 0, "y point")
	r = flag.IntP("radius", "r", 5, "radius")
	f = flag.StringP("file", "f", "output.png", "file output")
)

func main() {
	flag.Parse()
	Q(flag.Args())
	colors, err := parseColors(flag.Args()...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	Q(colors)

	f, err := os.Create(*f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	s := circle.ColorCircle(*r, colors...)

	err = png.Encode(f, s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parseColors(colors ...string) ([]color.Color, error) {
	c := make([]color.Color, len(colors))
	for i, s := range colors {
		color, err := parseColor(s)
		if err != nil {
			return nil, err
		}
		c[i] = color
	}
	return c, nil

}

func parseColor(s string) (color.Color, error) {
	if len(s) != 6 {
		return nil, fmt.Errorf("color should be 6 digits")
	}
	n, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return nil, fmt.Errorf("not hexadecimal: %v", err)
	}
	rgba := &color.RGBA{}
	rgba.B, n = uint8(n%256), n/256
	rgba.G, n = uint8(n%256), n/256
	rgba.R = uint8(n % 256)
	rgba.A = 255
	return rgba, nil

}
