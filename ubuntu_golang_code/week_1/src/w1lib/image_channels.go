package w1lib

import (
	"gocv.io/x/gocv"
)

func Image_Channels() {
	// Path of the image to be loaded
	// Here we are supplying a relative path
	imageFile := DATA_PATH + "/images/musk.jpg"

	// Read the image
	img := gocv.IMRead(imageFile, gocv.IMReadAnyColor)
	imgRgb := img.Clone()
	gocv.CvtColor(img, &imgRgb, gocv.ColorBGRToRGB)
	origwin := gocv.NewWindow("Original read")
	origwin.IMShow(img)
	window := gocv.NewWindow("Color Image with RGB from BGR")
	defer window.Close()
	window.IMShow(imgRgb)
	// Pull in individual channels
	bgr := gocv.Split(img)
	rwin := gocv.NewWindow("Red channel")
	defer rwin.Close()
	rwin.IMShow(bgr[0])
	gwin := gocv.NewWindow("Green channel")
	defer gwin.Close()
	gwin.IMShow(bgr[1])
	bwin := gocv.NewWindow("Blue channel")
	defer bwin.Close()
	bwin.IMShow(bgr[2])

	for {
		c := gocv.WaitKey(20)
		if c == 27 {
			break
		}
	}
}
