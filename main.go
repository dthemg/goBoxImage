package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"math/rand"
	"os"

	"github.com/fogleman/gg"
)

// Pixel struct
type pixel struct {
	r uint32
	g uint32
	b uint32
}

func main() {
	// Set format
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	imgFile, err := os.Open("resources/kingfisher.jpg")

	if err != nil {
		fmt.Println("img.jpg not found")
	}
	defer imgFile.Close()

	imgCfg, _, err := image.DecodeConfig(imgFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	width := imgCfg.Width
	height := imgCfg.Height

	// Reset io reader
	imgFile.Seek(0, 0)

	// Load image
	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read pixel values into array
	var rgbArr [][]pixel
	for w := 0; w < width; w++ {
		var row []pixel
		for h := 0; h < height; h++ {
			r, g, b, _ := img.At(w, h).RGBA()
			row = append(row, pixel{r, g, b})
		}
		rgbArr = append(rgbArr, row)
	}

	const maxPxval float64 = 65535

	canvas := gg.NewContext(width, height)
	canvas.SetHexColor("#0000ff")

	nSquares := 10000
	for n := 0; n < nSquares; n++ {
		wIdx := rand.Intn(width)
		hIdx := rand.Intn(height)

		wSz := 5 + rand.Intn(100)
		hSz := 5 + rand.Intn(100)

		wSzCap := wSz
		if wIdx+wSz >= width {
			wSzCap = width - wIdx
		}
		hSzCap := hSz
		if hIdx+hSz >= height {
			hSzCap = height - hIdx
		}

		var rSum float64 = 0
		var gSum float64 = 0
		var bSum float64 = 0

		for i := wIdx; i < wIdx+wSzCap; i++ {
			for j := hIdx; j < hIdx+hSzCap; j++ {
				pixel := rgbArr[i][j]
				rSum += float64(pixel.r)
				gSum += float64(pixel.g)
				bSum += float64(pixel.b)
			}
		}
		wSzCapFloat := float64(wSzCap)
		hSzCapFloat := float64(hSzCap)
		var rpx = rSum / wSzCapFloat / hSzCapFloat / maxPxval
		var gpx = gSum / wSzCapFloat / hSzCapFloat / maxPxval
		var bpx = bSum / wSzCapFloat / hSzCapFloat / maxPxval

		canvas.DrawRectangle(
			float64(wIdx-wSz/2),
			float64(hIdx-hSz/2),
			float64(wSz),
			float64(hSz),
		)
		canvas.SetRGBA(rpx, gpx, bpx, 0.5)
		canvas.Fill()
	}

	canvas.SavePNG("out.png")
}
