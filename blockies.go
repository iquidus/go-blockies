package blockies

import (
	"crypto/rand"
	"fmt"
	"io"
	"math"
	"strings"

	svg "github.com/ajstarks/svgo"
)

// Options is a struct that contains the options for the blockie
type Options struct {
	Size      int
	Scale     int
	Color     string
	BgColor   string
	SpotColor string
}

// Blockie is a struct that contains the blockie data
type Blockie struct {
	Options   *Options
	Randseed  []int32
	ImageData []int
}

// reverseInts reverses the order of the elements in a slice of ints
func reverseInts(input []int) {
	first := 0
	last := len(input) - 1
	for first < last {
		input[first], input[last] = input[last], input[first]
		first++
		last--
	}
}

// randomAddress generates a random address
func randomAddress() string {
	bytes := make([]byte, 20)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("0x%x", bytes)
}

// seedrand seeds the pseudorandom number generator
func (b *Blockie) seedrand(seed []int32) {
	b.Randseed = []int32{0, 0, 0, 0}
	for i := 0; i < len(seed); i++ {
		b.Randseed[i%4] = (b.Randseed[i%4] << 5) - b.Randseed[i%4] + seed[i]
	}
}

// rand generates a pseudorandom number
func (b *Blockie) rand() float64 {
	t := b.Randseed[0] ^ (b.Randseed[0] << 11)
	b.Randseed[0] = b.Randseed[1]
	b.Randseed[1] = b.Randseed[2]
	b.Randseed[2] = b.Randseed[3]
	b.Randseed[3] = b.Randseed[3] ^ (b.Randseed[3] >> 19) ^ t ^ (t >> 8)
	return float64((b.Randseed[3])>>0) / float64((1<<31)>>0)
}

// createColor generates a pseudorandom color
func (b *Blockie) createColor() string {
	h := int(b.rand() * 360)
	s := b.rand()*60 + 40
	l := (b.rand() + b.rand() + b.rand() + b.rand()) * 25

	color := fmt.Sprintf("hsl(%d,%.2f%%,%.2f%%)", h, s, l)
	return color
}

// createImageData generates the blockie image data
// as an int slice in which each element is 0, 1 or 2
// 0: background color
// 1: main color
// 2: spot color
func (b *Blockie) createImageData() []int {
	width := b.Options.Size
	height := b.Options.Size

	dataWidth := int(math.Ceil(float64(width) / 2))
	mirrorWidth := width - dataWidth

	var data []int
	for y := 0; y < height; y++ {
		row := make([]int, dataWidth)
		for x := 0; x < dataWidth; x++ {
			row[x] = int(math.Floor(b.rand() * 2.3))
		}
		r := append([]int{}, row[:mirrorWidth]...)
		reverseInts(r)
		row = append(row, r...)

		for i := 0; i < len(row); i++ {
			data = append(data, row[i])
		}
	}

	return data
}

// New creates a new blockie
func New(address string, options *Options) *Blockie {
	b := &Blockie{}
	if options != nil {
		b.Options = options
	} else {
		b.Options = &Options{}
	}

	var a string
	if address == "" {
		a = randomAddress()
	} else {
		a = strings.ToLower(address)
	}

	b.seedrand([]int32(a))

	if b.Options.Size == 0 {
		b.Options.Size = 8
	}
	if b.Options.Scale == 0 {
		b.Options.Scale = 4
	}
	if b.Options.Color == "" {
		b.Options.Color = b.createColor()
	}
	if b.Options.BgColor == "" {
		b.Options.BgColor = b.createColor()
	}
	if b.Options.SpotColor == "" {
		b.Options.SpotColor = b.createColor()
	}

	b.ImageData = b.createImageData()

	return b
}

// Write writes the blockie to an io.Writer
func (b *Blockie) Write(w io.Writer) {
	d := b.Options.Size * b.Options.Scale

	canvas := svg.New(w)
	canvas.Start(d, d)
	canvas.Rect(0, 0, d, d, "fill:"+b.Options.BgColor)

	for i, data := range b.ImageData {
		if data != 0 {
			row := int(math.Floor(float64(i / b.Options.Size)))
			col := i % b.Options.Size

			color := b.Options.Color
			if data == 2 {
				color = b.Options.SpotColor
			}
			canvas.Rect(col*b.Options.Scale, row*b.Options.Scale, b.Options.Scale, b.Options.Scale, "fill:"+color)
		}
	}

	canvas.End()
}
