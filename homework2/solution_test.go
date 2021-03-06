// Disclaimer
// Some of the tests here are not fit to represent real premultiplied alpha composition.
// They test a set of arbitrary rules for something similar.

package main

import (
	"errors"
	"fmt"
	"testing"
)

func assertColor(pixel Pixel, rgb ...byte) error {
	if pixel.Color().Red != rgb[0] {
		error_str := fmt.Sprintf("Wrong Red component: expected %d, got %d", rgb[0],
			pixel.Color().Red)
		return errors.New(error_str)
	}

	if pixel.Color().Green != rgb[1] {
		error_str := fmt.Sprintf("Wrong Green component: expected %d, got %d", rgb[1],
			pixel.Color().Green)
		return errors.New(error_str)
	}

	if pixel.Color().Blue != rgb[2] {
		error_str := fmt.Sprintf("Wrong Blue component: expected %d, got %d", rgb[2],
			pixel.Color().Blue)
		return errors.New(error_str)
	}

	return nil
}

func TestBasicRGBCall(t *testing.T) {
	data := []byte{
		0, 12, 244, 13, 26, 52, 31, 33, 41,
	}
	header := Header{"RGB", 3}
	picture := ParseImage(data, header)

	if err := assertColor(picture.InspectPixel(0, 0), 0, 12, 244); err != nil {
		t.Error(err)
	}

	if err := assertColor(picture.InspectPixel(1, 0), 13, 26, 52); err != nil {
		t.Error(err)
	}
}

func TestBasicRGBACall(t *testing.T) {
	data := []byte{
		0, 12, 244, 128, 14, 26, 52, 127, 31, 33, 41, 255, 36, 133, 241, 255,
	}
	header := Header{"RGBA", 4}
	picture := ParseImage(data, header)

	first_pixel := picture.InspectPixel(0, 0)
	if err := assertColor(first_pixel, 0, 6, 122); err != nil {
		t.Error(err)
	}

	second_pixel := picture.InspectPixel(3, 0)
	if err := assertColor(second_pixel, 36, 133, 241); err != nil {
		t.Error(err)
	}

	third_pixel := picture.InspectPixel(2, 0)
	if err := assertColor(third_pixel, 31, 33, 41); err != nil {
		t.Error(err)
	}
}

func TestBasicRGBARowsCall(t *testing.T) {
	data := []byte{
		0, 12, 244, 127, 14, 26, 52, 127,
		31, 33, 41, 255, 36, 133, 241, 255,
	}

	var pixel Pixel
	header := Header{"RGBA", 2}
	picture := ParseImage(data, header)

	pixel = picture.InspectPixel(1, 1)
	if err := assertColor(pixel, 36, 133, 241); err != nil {
		t.Error(err)
	}

	pixel = picture.InspectPixel(0, 1)
	if err := assertColor(pixel, 31, 33, 41); err != nil {
		t.Error(err)
	}
	pixel = picture.InspectPixel(1, 0)
	if err := assertColor(pixel, 6, 12, 25); err != nil {
		t.Error(err)
	}
}

func TestBasicBGRARowsCall(t *testing.T) {
	data := []byte{
		0, 12, 244, 127, 14, 26, 52, 127,
		31, 33, 41, 255, 36, 133, 241, 255,
	}

	var pixel Pixel
	header := Header{"BGRA", 2}
	picture := ParseImage(data, header)

	pixel = picture.InspectPixel(1, 1)
	if err := assertColor(pixel, 241, 133, 36); err != nil {
		t.Error(err)
	}

	pixel = picture.InspectPixel(1, 0)
	if err := assertColor(pixel, 25, 12, 6); err != nil {
		t.Error(err)
	}
}

func TestBasicAlphaNormalization(t *testing.T) {
	data := []byte{
		22, 12, 244, 127,
	}

	var pixel Pixel
	header := Header{"RGBA", 1}
	picture := ParseImage(data, header)

	pixel = picture.InspectPixel(0, 0)
	if err := assertColor(pixel, 10, 5, 121); err != nil {
		t.Error(err)
	}
}

func TestAlphaJustBelowHalf(t *testing.T) {
	data := []byte{
		1, 255, 0, 127, // alpha 127 is below 0.5
	}

	var pixel Pixel
	header := Header{"RGBA", 1}
	picture := ParseImage(data, header)

	pixel = picture.InspectPixel(0, 0)
	if err := assertColor(pixel, 0, 127, 0); err != nil {
		t.Error(err)
	}
}

func TestAlphaJustAboveHalf(t *testing.T) {
	data := []byte{
		1, 255, 0, 128, // alpha 128 is above 0.5
	}

	var pixel Pixel
	header := Header{"RGBA", 1}
	picture := ParseImage(data, header)

	pixel = picture.InspectPixel(0, 0)
	if err := assertColor(pixel, 0, 128, 0); err != nil {
		t.Error(err)
	}
}

func TestPixelColorStringRepresentation(t *testing.T) {
	data := []byte{
		22, 0, 255, 255,
	}

	header := Header{"RGBA", 1}
	picture := ParseImage(data, header)

	pixel := picture.InspectPixel(0, 0)
	colour := pixel.Color()

	found := fmt.Sprint(colour)
	expected := "Red: 22, Green: 0, Blue: 255"

	if found != expected {
		t.Errorf("String representation was wrong. Expected %s but got %s", expected,
			found)
	}
}

func TestBasicGRBCall(t *testing.T) {
	data := []byte{
		0, 12, 244, 13, 26, 52, 31, 33, 41,
	}

	header := Header{"GRB", 3}
	picture := ParseImage(data, header)

	if err := assertColor(picture.InspectPixel(0, 0), 12, 0, 244); err != nil {
		t.Error(err)
	}

	if err := assertColor(picture.InspectPixel(2, 0), 33, 31, 41); err != nil {
		t.Error(err)
	}
}

func TestZeroAlphaWithRGBA(t *testing.T) {
	data := []byte{
		0, 12, 244, 0, 14, 26, 52, 127,
		31, 33, 41, 255, 36, 133, 241, 255,
	}

	var pixel Pixel
	header := Header{"RGBA", 2}
	picture := ParseImage(data, header)

	pixel = picture.InspectPixel(0, 0)
	if err := assertColor(pixel, 0, 0, 0); err != nil {
		t.Error(err)
	}
}

func TestMaxAlphaWithRGBA(t *testing.T) {
	data := []byte{
		0, 12, 244, 255, 1, 127, 255, 255,
		31, 128, 41, 255, 36, 133, 241, 255,
	}

	var pixel Pixel
	header := Header{"RGBA", 2}
	picture := ParseImage(data, header)

	pixel = picture.InspectPixel(0, 0)
	if err := assertColor(pixel, 0, 12, 244); err != nil {
		t.Error(err)
	}

	pixel = picture.InspectPixel(1, 0)
	if err := assertColor(pixel, 1, 127, 255); err != nil {
		t.Error(err)
	}

	pixel = picture.InspectPixel(0, 1)
	if err := assertColor(pixel, 31, 128, 41); err != nil {
		t.Error(err)
	}

	pixel = picture.InspectPixel(1, 1)
	if err := assertColor(pixel, 36, 133, 241); err != nil {
		t.Error(err)
	}
}
