package main

import (
	"fmt"
)

type Header struct {
	Format    string
	LineWidth uint
	Encoding  string
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

func (img *Image) InspectPixel(x uint, y uint) (*Pixel, *ImageError) {
	index := y*img.header.LineWidth + x

	if index >= (uint)(len(img.data)) {
		return nil, newImageError("Index out of range")
	}

	return &img.data[index], nil
}

type ImageError string

func (e ImageError) Error() string {
	return string(e)
}

func newImageError(message string) *ImageError {
	return (*ImageError)(&message)
}

func isHeaderValid(header Header) (err *ImageError) {
	formatLen := len(header.Format)

	if formatLen < 3 || formatLen > 4 {
		err = newImageError("Header is too long or too short")
		return
	}

	formatInt := 0

	for _, char := range header.Format {
		switch char {
		case 'R':
			formatInt += 10
		case 'B':
			formatInt += 100
		case 'G':
			formatInt += 1000
		case 'A':
			formatInt += 1
		default:
			err = newImageError("Wrong letter in header format")
			return
		}
	}

	if formatInt != 1110 && formatInt != 1111 {
		err = newImageError("Header does not have red, green or blue component")
		return
	}

	if header.Encoding != "None" && header.Encoding != "RLE" {
		err = newImageError("Wrong header encoding. Should be None or RLE")
		return
	}

	return
}

func ParseImage(data []byte, header Header) (*Image, *ImageError) {

	image := new(Image)

	if headerError := isHeaderValid(header); headerError != nil {
		return nil, headerError
	}

	image.header = header
	var pixel *Pixel

	parsers := make(map[string]func() *ImageError)

	parsers["None"] = func() *ImageError {

		formatLen := len(header.Format)

		if len(data)%formatLen != 0 {
			return newImageError("Not enough data for whole pixel")
		}

		for index, colourIntesity := range data {

			formatIndex := index % formatLen
			if formatIndex == 0 {
				pixel = new(Pixel)
			}

			colour := header.Format[formatIndex]

			switch colour {
			case 'R':
				pixel.Red = colourIntesity
			case 'G':
				pixel.Green = colourIntesity
			case 'B':
				pixel.Blue = colourIntesity
			case 'A':
				pixel.Alpha = colourIntesity
				pixel.needsPremultiply = true
			}

			if formatIndex == formatLen-1 {
				pixel.premultiply()
				image.data = append(image.data, *pixel)
			}

		}

		return nil
	}

	parsers["RLE"] = func() *ImageError {

		var pixelsCount int
		var colourIntesity byte
		var colourIndex int

		for i := 0; i < len(data); i++ {
			pixelsCount = (int)(data[i])

			pixel = new(Pixel)

			for formatIndex := range header.Format {

				colourIndex = i + 1 + formatIndex

				if colourIndex >= len(data) {
					return newImageError("Not enough data for pixel")
				}

				colourIntesity = data[colourIndex]
				colour := header.Format[formatIndex]

				switch colour {
				case 'R':
					pixel.Red = colourIntesity
				case 'G':
					pixel.Green = colourIntesity
				case 'B':
					pixel.Blue = colourIntesity
				case 'A':
					pixel.Alpha = colourIntesity
					pixel.needsPremultiply = true
				}
			}

			pixel.premultiply()

			for i := 0; i < pixelsCount; i++ {
				image.data = append(image.data, *pixel)
			}

			i += len(header.Format)
		}

		return nil
	}

	if err := parsers[header.Encoding](); err != nil {
		return nil, err
	}

	pixelsCount := len(image.data)

	if pixelsCount > 0 && (uint)(pixelsCount)%header.LineWidth > 0 {
		return nil, newImageError("Not enough data for a whole row")
	}

	return image, nil
}

func alphaBlend(colour byte, alpha byte) byte {
	return (byte)(((int)(colour)*(int)(alpha) + 127) / 255)
}
