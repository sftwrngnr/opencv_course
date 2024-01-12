package w1lib

import (
	"gocv.io/x/gocv"
)

func CreateNewImages() {
	// Path of the image to be loaded
	// Here we are supplying a relative path
	imageFile := DATA_PATH + "/images/boy.jpg"
	emptyImg := gocv.Zeros(100, 200, gocv.MatChannels3)
	gocv.IMWrite("results/emptyMatrix.png", emptyImg)
	// Read the image
	img := gocv.IMRead(imageFile, gocv.IMReadColor)
	gocv.IMWrite("results/testImage.png", img)
	emptyImg = gocv.Zeros(100, 200, gocv.MatChannels3)
	emptyImg.MultiplyFloat(255)
	gocv.IMWrite("results/emptyMatrix1.png", emptyImg)
	emptyOriginal := gocv.
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
