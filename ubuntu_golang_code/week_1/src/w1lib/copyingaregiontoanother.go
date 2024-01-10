package w1lib

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

func CopyingARegionToAnother() {
	// Path of the image to be loaded
	// Here we are supplying a relative path
	imageFile := DATA_PATH + "/images/boy.jpg"

	// Read the image
	img := gocv.IMRead(imageFile, gocv.IMReadColor)
	window := gocv.NewWindow("Image")
	defer window.Close()
	window.IMShow(img)
	gocv.WaitKey(0)
	copiedImage := img.Clone()
	copywin := gocv.NewWindow("Copied Image")
	defer copywin.Close()
	copywin.IMShow(copiedImage)
	gocv.WaitKey(0)
	copyroi := copiedImage.Region(image.Rectangle{image.Point{40, 200}, image.Point{180, 320}})
	dims := copyroi.Size()
	fmt.Printf("%v\n", dims)

	for {
		c := gocv.WaitKey(20)
		if c == 27 {
			break
		}
	}
}
