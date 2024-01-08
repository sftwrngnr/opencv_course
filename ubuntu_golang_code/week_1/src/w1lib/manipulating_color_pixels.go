package w1lib

import (
	"image"

	"gocv.io/x/gocv"
)

func ManipulatingColorPixels() {
	imagePath := DATA_PATH + "/images/number_zero.jpg"
	testImage := gocv.IMRead(imagePath, gocv.IMReadColor)

	PrintImageChannelValues(&testImage) //In Utils
	//s := gocv.Scalar{0.0, 255.0, 255.0, 0.0}
	//myRegion := testImage.Region(image.Rectangle{Min: image.Point{0,0}, Max: image.Point{0,0}})
	bgr := gocv.Split(testImage)
	bgr[0].SetUCharAt(0, 0, 0)
	bgr[1].SetUCharAt(0, 0, 255)
	bgr[2].SetUCharAt(0, 0, 255)

	gocv.Merge(bgr, &testImage)
	gocv.IMWrite("results/zero1.png", testImage)
	bgr[0].SetUCharAt(1, 1, 255)
	bgr[1].SetUCharAt(1, 1, 255)
	bgr[2].SetUCharAt(1, 1, 0)
	gocv.Merge(bgr, &testImage)
	gocv.IMWrite("results/zero2.png", testImage)
	bgr[0].SetUCharAt(2, 2, 255)
	bgr[1].SetUCharAt(2, 2, 0)
	bgr[2].SetUCharAt(2, 2, 255)
	gocv.Merge(bgr, &testImage)
	gocv.IMWrite("results/zero3.png", testImage)
	myRegion := image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{3, 3}}
	tr := bgr[0].Region(myRegion)
	tr.SetTo(gocv.Scalar{Val1: 255, Val2: 0, Val3: 0, Val4: 0})
	tr = bgr[1].Region(myRegion)
	tr.SetTo(gocv.Scalar{Val1: 0, Val2: 0, Val3: 0, Val4: 0})
	tr = bgr[2].Region(myRegion)
	tr.SetTo(gocv.Scalar{Val1: 0, Val2: 0, Val3: 0, Val4: 0})
	gocv.Merge(bgr, &testImage)
	myRegion = image.Rectangle{Min: image.Point{0, 3}, Max: image.Point{3, 6}}
	tr = bgr[0].Region(myRegion)
	tr.SetTo(gocv.Scalar{Val1: 0, Val2: 0, Val3: 0, Val4: 0})
	tr = bgr[1].Region(myRegion)
	tr.SetTo(gocv.Scalar{Val1: 255, Val2: 0, Val3: 0, Val4: 0})
	tr = bgr[2].Region(myRegion)
	tr.SetTo(gocv.Scalar{Val1: 0, Val2: 0, Val3: 0, Val4: 0})
	gocv.Merge(bgr, &testImage)
	myRegion = image.Rectangle{Min: image.Point{0, 6}, Max: image.Point{3, 9}}
	tr = bgr[0].Region(myRegion)
	tr.SetTo(gocv.Scalar{Val1: 0, Val2: 0, Val3: 0, Val4: 0})
	tr = bgr[1].Region(myRegion)
	tr.SetTo(gocv.Scalar{Val1: 0, Val2: 0, Val3: 0, Val4: 0})
	tr = bgr[2].Region(myRegion)
	tr.SetTo(gocv.Scalar{Val1: 255, Val2: 0, Val3: 0, Val4: 0})
	gocv.Merge(bgr, &testImage)
	gocv.IMWrite("results/zero.png", testImage)
	PrintImageChannelValues(&testImage)
	window := gocv.NewWindow("Current Image")
	defer window.Close()
	window.IMShow(testImage)
	for {
		c := gocv.WaitKey(20)
		if c == 27 {
			break
		}
	}

}
