package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
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

	fmt.Println(width)
	fmt.Println(height)

	// Reset io reader
	imgFile.Seek(0, 0)

	// Load image
	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(img.At(100, 100).RGBA())
	var rgbArr [649][649]pixel
	for w := 0; w < width; w++ {
		for h := 0; h < height; h++ {
			r, g, b, _ := img.At(w, h).RGBA()
			rgbArr[w][h] = pixel{r, g, b}
		}
	}

	// Make an image with small circles
	const gs int = 11

	for wg := 0; wg < 649/gs; wg++ {
		for hg := 0; hg < 649/gs; hg++ {
			var rSum float32 = 0
			var gSum float32 = 0
			var bSum float32 = 0
			for i := 0; i < gs; i++ {
				for j := 0; j < gs; j++ {
					pixel := rgbArr[wg*gs+i][hg*gs+j]
					rSum += float32(pixel.r)
					gSum += float32(pixel.g)
					bSum += float32(pixel.b)
				}
			}
		}
	}

	/*
		canvas := gg.NewContext(width, height)
		canvas.SetHexColor("#0000ff")
		canvas.SetLineWidth(2)
		canvas.DrawRectangle(100, 210, float64(imageWidth), float64(imageHeight))
		canvas.Stroke()
		canvas.DrawImage(image, 100, 210)

		//canvas.DrawCircle(500, 500, 400)
		//canvas.SetRGBA(0, 0, 0, 0.5)
		//canvas.Fill()
		canvas.SavePNG("test.png")
	*/
}

func getAverage(rgbArr [][]pixel) {

}
