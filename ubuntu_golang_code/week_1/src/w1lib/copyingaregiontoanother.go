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
	copywin.IMShow(copiedImage)
	gocv.WaitKey(0)
	copywin.Close()
	copyroi := copiedImage.Region(image.Rectangle{image.Point{180, 40}, image.Point{320, 180}})
	fmt.Printf("copyroi: Rows: %d Cols: %d\n", copyroi.Rows(), copyroi.Cols())

	myTmp := copiedImage.Region(image.Rectangle{image.Point{40, 10}, image.Point{40 + copyroi.Cols(), 10 + copyroi.Rows()}})
	fmt.Printf("myTmp: Rows: %d Cols: %d\n", myTmp.Rows(), myTmp.Cols())

	copyroi.CopyTo(&myTmp)
	myTmp = copiedImage.Region(image.Rectangle{image.Point{330, 10}, image.Point{330 + copyroi.Cols(), 10 + copyroi.Rows()}})
	copyroi.CopyTo(&myTmp)

	copywin = gocv.NewWindow("New Copied Image")
	copywin.IMShow(copiedImage)
	gocv.WaitKey(0)
	copywin.Close()
	gocv.IMWrite("results/outputImage.png", copiedImage)
}
