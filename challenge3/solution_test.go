package main

import (
	"errors"
	"fmt"
	"testing"
)

func assertColor(pixel *Pixel, rgb ...byte) error {

	if pixel == nil {
		return errors.New("Pixel was nil")
	}

	if pixel.Color().Red != rgb[0] {
		errorStr := fmt.Sprintf("Wrong Red component: expected %d, got %d", rgb[0],
			pixel.Color().Red)
		return errors.New(errorStr)
	}

	if pixel.Color().Green != rgb[1] {
		errorStr := fmt.Sprintf("Wrong Green component: expected %d, got %d", rgb[1],
			pixel.Color().Green)
		return errors.New(errorStr)
	}

	if pixel.Color().Blue != rgb[2] {
		errorStr := fmt.Sprintf("Wrong Blue component: expected %d, got %d", rgb[2],
			pixel.Color().Blue)
		return errors.New(errorStr)
	}

	return nil
}

func TestBasicRGBCall(t *testing.T) {
	data := []byte{
		0, 12, 244, 13, 26, 52, 31, 33, 41,
	}
	header := Header{"RGB", 3, "None"}

	var pixel *Pixel

	picture, parseError := ParseImage(data, header)

	if parseError != nil {
		t.Fatalf(fmt.Sprintf("Parsing the image returned error: %s",
			parseError.Error()))
	}

	pixel, _ = picture.InspectPixel(0, 0)

	if err := assertColor(pixel, 0, 12, 244); err != nil {
		t.Error(err)
	}

	pixel, _ = picture.InspectPixel(1, 0)

	if err := assertColor(pixel, 13, 26, 52); err != nil {
		t.Error(err)
	}
}

func TestBasicRGBACall(t *testing.T) {
	data := []byte{
		0, 12, 244, 128, 14, 26, 52, 127, 31, 33, 41, 255, 36, 133, 241, 255,
	}
	header := Header{"RGBA", 4, "None"}

	picture, parseError := ParseImage(data, header)

	if parseError != nil {
		t.Fatalf(fmt.Sprintf("Parsing the image returned error: %s",
			parseError.Error()))
	}

	firstPixel, _ := picture.InspectPixel(0, 0)
	if err := assertColor(firstPixel, 0, 6, 122); err != nil {
		t.Error(err)
	}

	secondPixel, _ := picture.InspectPixel(3, 0)
	if err := assertColor(secondPixel, 36, 133, 241); err != nil {
		t.Error(err)
	}

	thirdPixel, _ := picture.InspectPixel(2, 0)
	if err := assertColor(thirdPixel, 31, 33, 41); err != nil {
		t.Error(err)
	}
}

func TestBasicRGBARowsCall(t *testing.T) {
	data := []byte{
		0, 12, 244, 127, 14, 26, 52, 127,
		31, 33, 41, 255, 36, 133, 241, 255,
	}

	var pixel *Pixel
	header := Header{"RGBA", 2, "None"}

	picture, parseError := ParseImage(data, header)

	if parseError != nil {
		t.Fatalf(fmt.Sprintf("Parsing the image returned error: %s",
			parseError.Error()))
	}

	pixel, _ = picture.InspectPixel(1, 1)
	if err := assertColor(pixel, 36, 133, 241); err != nil {
		t.Error(err)
	}

	pixel, _ = picture.InspectPixel(0, 1)
	if err := assertColor(pixel, 31, 33, 41); err != nil {
		t.Error(err)
	}
	pixel, _ = picture.InspectPixel(1, 0)
	if err := assertColor(pixel, 7, 13, 26); err != nil {
		t.Error(err)
	}
}

func TestBasicBGRARowsCall(t *testing.T) {
	data := []byte{
		0, 12, 244, 127, 14, 26, 52, 127,
		31, 33, 41, 255, 36, 133, 241, 255,
	}

	var pixel *Pixel
	header := Header{"BGRA", 2, "None"}

	picture, parseError := ParseImage(data, header)

	if parseError != nil {
		t.Fatalf(fmt.Sprintf("Parsing the image returned error: %s",
			parseError.Error()))
	}

	pixel, _ = picture.InspectPixel(1, 1)
	if err := assertColor(pixel, 241, 133, 36); err != nil {
		t.Error(err)
	}

	pixel, _ = picture.InspectPixel(1, 0)
	if err := assertColor(pixel, 26, 13, 7); err != nil {
		t.Error(err)
	}
}

func TestBasicAlphaNormalization(t *testing.T) {
	data := []byte{
		22, 12, 244, 127,
	}

	var pixel *Pixel
	header := Header{"RGBA", 1, "None"}

	picture, parseError := ParseImage(data, header)

	if parseError != nil {
		t.Fatalf(fmt.Sprintf("Parsing the image returned error: %s",
			parseError.Error()))
	}

	pixel, _ = picture.InspectPixel(0, 0)
	if err := assertColor(pixel, 11, 6, 122); err != nil {
		t.Error(err)
	}
}

func TestAlphaJustBelowHalf(t *testing.T) {
	data := []byte{
		1, 255, 0, 127, // alpha 127 is below 0.5
	}

	var pixel *Pixel
	header := Header{"RGBA", 1, "None"}

	picture, parseError := ParseImage(data, header)

	if parseError != nil {
		t.Fatalf(fmt.Sprintf("Parsing the image returned error: %s",
			parseError.Error()))
	}

	pixel, _ = picture.InspectPixel(0, 0)
	if err := assertColor(pixel, 0, 127, 0); err != nil {
		t.Error(err)
	}
}

func TestAlphaJustAboveHalf(t *testing.T) {
	data := []byte{
		1, 255, 0, 128, // alpha 128 is above 0.5
	}

	var pixel *Pixel
	header := Header{"RGBA", 1, "None"}

	picture, parseError := ParseImage(data, header)

	if parseError != nil {
		t.Fatalf(fmt.Sprintf("Parsing the image returned error: %s",
			parseError.Error()))
	}

	pixel, _ = picture.InspectPixel(0, 0)
	if err := assertColor(pixel, 1, 128, 0); err != nil {
		t.Error(err)
	}
}

func TestPixelColorStringRepresentation(t *testing.T) {
	data := []byte{
		22, 0, 255, 255,
	}

	header := Header{"RGBA", 1, "None"}

	picture, parseError := ParseImage(data, header)

	if parseError != nil {
		t.Fatalf(fmt.Sprintf("Parsing the image returned error: %s",
			parseError.Error()))
	}

	pixel, _ := picture.InspectPixel(0, 0)
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

	header := Header{"GRB", 3, "None"}

	picture, parseError := ParseImage(data, header)

	if parseError != nil {
		t.Fatalf(fmt.Sprintf("Parsing the image returned error: %s",
			parseError.Error()))
	}

	var pixel *Pixel

	pixel, _ = picture.InspectPixel(0, 0)

	if err := assertColor(pixel, 12, 0, 244); err != nil {
		t.Error(err)
	}

	pixel, _ = picture.InspectPixel(2, 0)

	if err := assertColor(pixel, 33, 31, 41); err != nil {
		t.Error(err)
	}
}

func TestZeroAlphaWithRGBA(t *testing.T) {
	data := []byte{
		0, 12, 244, 0, 14, 26, 52, 127,
		31, 33, 41, 255, 36, 133, 241, 255,
	}

	var pixel *Pixel
	header := Header{"RGBA", 2, "None"}

	picture, parseError := ParseImage(data, header)

	if parseError != nil {
		t.Fatalf(fmt.Sprintf("Parsing the image returned error: %s",
			parseError.Error()))
	}

	pixel, _ = picture.InspectPixel(0, 0)
	if err := assertColor(pixel, 0, 0, 0); err != nil {
		t.Error(err)
	}
}

func TestMaxAlphaWithRGBA(t *testing.T) {
	data := []byte{
		0, 12, 244, 255, 1, 127, 255, 255,
		31, 128, 41, 255, 36, 133, 241, 255,
	}

	var pixel *Pixel
	header := Header{"RGBA", 2, "None"}

	picture, parseError := ParseImage(data, header)

	if parseError != nil {
		t.Fatalf(fmt.Sprintf("Parsing the image returned error: %s",
			parseError.Error()))
	}

	pixel, _ = picture.InspectPixel(0, 0)
	if err := assertColor(pixel, 0, 12, 244); err != nil {
		t.Error(err)
	}

	pixel, _ = picture.InspectPixel(1, 0)
	if err := assertColor(pixel, 1, 127, 255); err != nil {
		t.Error(err)
	}

	pixel, _ = picture.InspectPixel(0, 1)
	if err := assertColor(pixel, 31, 128, 41); err != nil {
		t.Error(err)
	}

	pixel, _ = picture.InspectPixel(1, 1)
	if err := assertColor(pixel, 36, 133, 241); err != nil {
		t.Error(err)
	}
}

func TestParseImageWithWrongFormats(t *testing.T) {

	data := []byte{}

	header := Header{"RGBADD", 2, "None"}

	if _, err := ParseImage(data, header); err == nil {
		t.Error("Parsing the image did not return error for wrong format RGBADD")
	}

	nothingRemotelyClose := Header{"SCII", 2, "None"}

	if _, err := ParseImage(data, nothingRemotelyClose); err == nil {
		t.Error("Parsing the image did not return error for wrong format SCII")
	}

	emptyHeader := Header{"", 10, "None"}

	if _, err := ParseImage(data, emptyHeader); err == nil {
		t.Error("Parsing the image did not return error for empty format")
	}

	repetetiveColour := Header{"RRGB", 10, "None"}

	if _, err := ParseImage(data, repetetiveColour); err == nil {
		t.Error("Parsing the image did not return error for wrong format RRGB")
	}

	repetetiveAlpha := Header{"RGBAA", 10, "None"}

	if _, err := ParseImage(data, repetetiveAlpha); err == nil {
		t.Error("Parsing the image did not return error for wrong format RGBAA")
	}

	moreAlphaThanBlue := Header{"RGAA", 10, "None"}

	if _, err := ParseImage(data, moreAlphaThanBlue); err == nil {
		t.Error("Parsing the image did not return error for wrong format RGAA")
	}
}

func TestWrongEncoding(t *testing.T) {

	header := Header{"RGBA", 2, "FooBar"}

	data := []byte{}

	if _, err := ParseImage(data, header); err == nil {
		t.Error("No error when wrong encoding is used")
	}

}

func TestWrongAmountOfData(t *testing.T) {

	header := Header{"RGBA", 2, "None"}

	tooShortForPixel := []byte{
		0, 12, 244,
	}

	if _, err := ParseImage(tooShortForPixel, header); err == nil {
		t.Error("No error when there was not enough data for a whole pixel")
	}

	tooShortForARow := []byte{
		0, 12, 244, 233,
	}

	if _, err := ParseImage(tooShortForARow, header); err == nil {
		t.Error("No error when there was not enough data for a whole row")
	}

}

func TestRLEFormat(t *testing.T) {

	header := Header{"RGBA", 5, "RLE"}

	data := []byte{
		3, 0, 12, 244, 255,
		1, 33, 2, 3, 255,
		1, 12, 15, 16, 38,
	}

	var pixel *Pixel

	picture, parseError := ParseImage(data, header)

	if parseError != nil {
		t.Fatalf(fmt.Sprintf("Parsing the image returned error: %s",
			parseError.Error()))
	}

	pixel, _ = picture.InspectPixel(2, 0)

	if err := assertColor(pixel, 0, 12, 244); err != nil {
		t.Error(err)
	}

	pixel, _ = picture.InspectPixel(3, 0)

	if err := assertColor(pixel, 33, 2, 3); err != nil {
		t.Error(err)
	}

}
