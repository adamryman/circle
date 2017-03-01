package main

import (
	"fmt"
	"image/png"
	"math"
	"os"

	"github.com/adamryman/circle"
	flag "github.com/spf13/pflag"
)

var (
	x = flag.IntP("x", "y", 0, "x point")
	y = flag.IntP("y", "x", 0, "y point")
	r = flag.IntP("radius", "r", 5, "radius")
	f = flag.StringP("file", "f", "output.png", "file output")
)

func main() {
	flag.Parse()

	f, err := os.Create(*f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	c := circle.Simple(*x, *y, *r)
	s := circle.Sector{
		Circle: *c,
		Θ1:     0,
		Θ2:     math.Pi + math.Pi/3,
	}

	err = png.Encode(f, &s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
