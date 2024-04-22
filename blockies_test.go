package blockies

import (
	"fmt"
	"testing"
)

func TestRandomAddress(t *testing.T) {
	address := randomAddress()
	if len(address) != 42 {
		t.Errorf("Address length is not 42")
	}
}

func TestReverseInts(t *testing.T) {
	// with no elements
	input := []int{}
	reverseInts(input)
	if len(input) != 0 {
		t.Errorf("Length is not 0")
	}
	// with one element
	input = []int{1}
	reverseInts(input)
	if input[0] != 1 {
		t.Errorf("First element is not 1")
	}
	// with many elements
	input = []int{1, 2, 3, 4, 5}
	reverseInts(input)
	if input[0] != 5 {
		t.Errorf("First element is not 5")
	}
	if input[1] != 4 {
		t.Errorf("Second element is not 4")
	}
	if input[2] != 3 {
		t.Errorf("Third element is not 3")
	}
	if input[3] != 2 {
		t.Errorf("Fourth element is not 2")
	}
	if input[4] != 1 {
		t.Errorf("Last element is not 1")
	}
}

func TestNewRandom(t *testing.T) {
	b := New("", nil)
	if b.Options.Size != 8 {
		t.Errorf("Size is not 8")
	}
	if b.Options.Scale != 4 {
		t.Errorf("Scale is not 4")
	}
	if b.Options.Color == "" {
		t.Errorf("Color is empty")
	}
	if b.Options.BgColor == "" {
		t.Errorf("BgColor is empty")
	}
}

func TestNewDefault(t *testing.T) {
	b := New("0x1234567890abcdef1234567890abcdef12345678", nil)
	if b.Options.Size != 8 {
		t.Errorf("Size is not 8")
	}
	if b.Options.Scale != 4 {
		t.Errorf("Scale is not 4")
	}
	if b.Options.Color == "" {
		t.Errorf("Color is empty")
	}
	if b.Options.BgColor == "" {
		t.Errorf("BgColor is empty")
	}
	if len(b.ImageData) != b.Options.Size*b.Options.Size {
		t.Errorf("ImageData length is incorrect")
	}
	have := fmt.Sprint(b.ImageData)
	want := "[0 0 0 0 0 0 0 0 0 0 1 0 0 1 0 0 1 1 1 1 1 1 1 1 0 0 0 0 0 0 0 0 1 1 0 0 0 0 1 1 1 0 1 0 0 1 0 1 1 1 1 0 0 1 1 1 0 1 1 2 2 1 1 0]"
	if have != want {
		t.Errorf("ImageData is incorrect: have: %v, want: %v", have, want)
	}
}

func TestNewCustom(t *testing.T) {
	b := New("0x1234567890abcdef1234567890abcdef12345678", &Options{
		Size:      10,
		Scale:     5,
		Color:     "#ff0000",
		BgColor:   "#00ff00",
		SpotColor: "#0000ff",
	})
	if b.Options.Size != 10 {
		t.Errorf("Size is not 8")
	}
	if b.Options.Scale != 5 {
		t.Errorf("Scale is not 4")
	}
	if b.Options.Color != "#ff0000" {
		t.Errorf("Color is incorrect")
	}
	if b.Options.BgColor != "#00ff00" {
		t.Errorf("BgColor is incorrect")
	}
	if b.Options.SpotColor != "#0000ff" {
		t.Errorf("SpotColor is incorrect")
	}
	if len(b.ImageData) != b.Options.Size*b.Options.Size {
		t.Errorf("ImageData length is incorrect")
	}
	have := fmt.Sprint(b.ImageData)
	want := "[1 2 1 1 1 1 1 1 2 1 1 0 1 0 0 0 0 1 0 1 1 1 1 0 1 1 0 1 1 1 2 0 1 0 0 0 0 1 0 2 0 0 0 0 1 1 0 0 0 0 0 1 1 1 1 1 1 1 1 0 0 0 0 0 1 1 0 0 0 0 1 0 0 1 0 0 1 0 0 1 1 0 1 1 1 1 1 1 0 1 0 0 1 1 2 2 1 1 0 0]"
	if have != want {
		t.Errorf("ImageData is incorrect: have: %v, want: %v", have, want)
	}
}
