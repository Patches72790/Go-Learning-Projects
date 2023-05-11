package main

import (
	"image"
	"image/png"
	"os"
)

func main() {

	rect := image.NewRGBA(image.Rect(0, 0, 512, 512))

	out, err := os.Create("out.png")

	if err != nil {
		panic("Error opening out file")
	}

	defer out.Close()
	err = png.Encode(out, rect)

	if err != nil {
		panic(err)
	}
}
