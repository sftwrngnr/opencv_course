package w1lib

import (
	"gocv.io/x/gocv"
)

func ImagesWithAlphaChannel() {
	// Path of the image to be loaded
	// Here we are supplying a relative path
	imageFile := DATA_PATH + "/images/musk.jpg"

	// Read the image
	img := gocv.IMRead(imageFile, gocv.IMReadColor)
	imgRgb := img.Clone()
	gocv.CvtColor(img, &imgRgb, gocv.ColorBGRToRGB)

	window := gocv.NewWindow("Color Image with RGB from BGR")
	defer window.Close()
	window.IMShow(imgRgb)

	for {
		c := gocv.WaitKey(20)
		if c == 27 {
			break
		}
	}
}
