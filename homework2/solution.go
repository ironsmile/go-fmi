package main

import (
	"fmt"
)

type Header struct {
	Format    string
	LineWidth uint
}

type Pixel struct {
	Red              byte
	Green            byte
	Blue             byte
	Alpha            byte
	needsPremultiply bool
}

// In case *someone* calls us with a wrong method name
// I will penalise him/her with one extra copy
func (pixel *Pixel) Color() Pixel {
	return pixel.Colour()
}

// quack quack!
func (pixel *Pixel) Colour() Pixel {
	return *pixel
}

func (pixel *Pixel) premultiply() {
	if !pixel.needsPremultiply {
		return
	}

	pixel.needsPremultiply = false

	pixel.Red = alphaBlend(pixel.Red, pixel.Alpha)
	pixel.Green = alphaBlend(pixel.Green, pixel.Alpha)
	pixel.Blue = alphaBlend(pixel.Blue, pixel.Alpha)

}

func (pixel Pixel) String() string {
	return fmt.Sprintf("Red: %d, Green: %d, Blue: %d", pixel.Red, pixel.Green,
		pixel.Blue)
}

type Image struct {
	header Header
	data   []Pixel
}

func (img *Image) InspectPixel(x uint, y uint) Pixel {
	index := y*img.header.LineWidth + x
	return img.data[index]
}

func ParseImage(data []byte, header Header) *Image {
	image := new(Image)
	image.header = header

	var pixel *Pixel
	format_len := len(header.Format)

	for index, colour_intensity := range data {

		format_index := index % format_len
		if format_index == 0 {
			pixel = new(Pixel)
		}

		colour := header.Format[format_index]

		switch (string)(colour) {
		case "R":
			pixel.Red = colour_intensity
		case "G":
			pixel.Green = colour_intensity
		case "B":
			pixel.Blue = colour_intensity
		case "A":
			pixel.Alpha = colour_intensity
			pixel.needsPremultiply = true
		}

		if format_index == format_len-1 {
			pixel.premultiply()
			image.data = append(image.data, *pixel)
		}

	}

	return image
}

func alphaBlend(colour byte, alpha byte) byte {
	// WRONG!
	// should be (byte)(((int)(colour)*(int)(alpha) + 127) / 255)
	// but due to explicit requirements it is as it is
	return (byte)((((float64)(colour) / 255.0) * ((float64)(alpha) / 255.0)) * 255.0)
}
