package main

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"image/png"
	"os"
)

func main() {
	// query number of active displays
	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		// get display boundaries
		bounds := screenshot.GetDisplayBounds(i)

		// capture screenshots into image data
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}

		// specify saved filename
		fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		file, _ := os.Create(fileName)
		defer file.Close()

		// encode image data to PNG format and write it to the output file
		png.Encode(file, img)

		fmt.Printf("Display #%d : %v \"%s\"\n", i, bounds, fileName)
	}
}
