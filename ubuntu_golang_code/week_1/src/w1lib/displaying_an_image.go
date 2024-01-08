package w1lib

import (
	"gocv.io/x/gocv"
)

func Displaying_an_Image() {
	// Path of the image to be loaded
	// Here we are supplying a relative path
	imageFile := DATA_PATH + "/images/number_zero.jpg"

	// Read the image
	img := gocv.IMRead(imageFile, 0)
	gocv.IMWrite("results/testImage.png", img)
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
