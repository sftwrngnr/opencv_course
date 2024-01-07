package w1lib

import (
	"gocv.io/x/gocv"
)

func Intro() {
	// Path of the image to be loaded
	// Here we are supplying a relative path
	imageFile := DATA_PATH + "/images/boy.jpg"

	// Read the image
	img := gocv.IMRead(imageFile, gocv.IMReadColor)
	window := gocv.NewWindow("Boy")
	window.IMShow(img)

	for {
		c := gocv.WaitKey(20)
		if c == 27 {
			break
		}
	}
	window.Close()
}
