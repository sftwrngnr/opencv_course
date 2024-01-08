package w1lib

import (
	"fmt"

	"gocv.io/x/gocv"
)

func Color_images() {
	// Path of the image to be loaded
	// Here we are supplying a relative path
	imageFile := DATA_PATH + "/images/musk.jpg"

	// Read the image
	img := gocv.IMRead(imageFile, gocv.IMReadColor)
	fmt.Printf("image Dimension ={%d x %d}\n", img.Rows(), img.Cols())
	window := gocv.NewWindow("Color Image")
	defer window.Close()
	window.IMShow(img)

	for {
		c := gocv.WaitKey(20)
		if c == 27 {
			break
		}
	}
}
