package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/fogleman/gg"
)

// Pixel struct
type pixel struct {
	r uint32
	g uint32
	b uint32
}

const width int = 1000
const height int = 1000

func main() {
	// Set format
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	imgFile, err := os.Open("./me.jpg")

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

	var rgbArr [649][649]pixel
	for w := 0; w < width; w++ {
		for h := 0; h < height; h++ {
			r, g, b, _ := img.At(w, h).RGBA()
			rgbArr[w][h] = pixel{r, g, b}
		}
	}

	const gs int = 20
	const gsSq = float64(gs * gs)
	const maxPxval float64 = 65535

	canvas := gg.NewContext(width, height)
	canvas.SetHexColor("#0000ff")
	canvas.SetLineWidth(2)

	for wg := 0; wg < 649/gs; wg++ {
		for hg := 0; hg < 649/gs; hg++ {
			var rSum float64 = 0
			var gSum float64 = 0
			var bSum float64 = 0
			for i := 0; i < gs; i++ {
				for j := 0; j < gs; j++ {
					pixel := rgbArr[wg*gs+i][hg*gs+j]
					rSum += float64(pixel.r)
					gSum += float64(pixel.g)
					bSum += float64(pixel.b)
				}
			}
			var rpx = rSum / gsSq / maxPxval
			var gpx = gSum / gsSq / maxPxval
			var bpx = bSum / gsSq / maxPxval
			canvas.DrawCircle(float64(wg*gs), float64(hg*gs), float64(gs)/2.)
			canvas.SetRGB(rpx, gpx, bpx)
			canvas.Fill()
		}
	}
	canvas.SavePNG("out.png")
}
