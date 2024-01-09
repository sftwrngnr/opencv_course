package w1lib

import "gocv.io/x/gocv"

func SplittingAndMergingChannels() {
	imagePath := DATA_PATH + "/images/musk.jpg"

	// Read image in Grayscale format
	testImage := gocv.IMRead(imagePath, gocv.IMReadAnyColor)
	bgr := gocv.Split(testImage)
	gocv.IMWrite("results/blueChannel.png", bgr[0])
	gocv.IMWrite("results/greenChannel.png", bgr[1])
	gocv.IMWrite("results/redChannel.png", bgr[2])

	mergedImage := gocv.NewMat()
	/*
			rws := testImage.Rows()
			cols := testImage.Cols()
			ons := gocv.Ones(rws, cols, gocv.MatTypeCV8U)

		myarr := make([]gocv.Mat, 0)
		myarr = append(myarr, bgr[0])
		myarr = append(myarr, bgr[1])
		myarr = append(myarr, bgr[2])
	*/

	gocv.Merge(bgr, &mergedImage)
	gocv.IMWrite("results/mergedOutput.png", mergedImage)

	bgrwin := gocv.NewWindow("Image BGR")
	defer bgrwin.Close()
	bgrwin.IMShow(mergedImage)
	bwin := gocv.NewWindow("Blue Channel")
	defer bwin.Close()
	bwin.IMShow(bgr[0])
	gwin := gocv.NewWindow("Green Channel")
	defer gwin.Close()
	gwin.IMShow(bgr[1])
	rwin := gocv.NewWindow("Red Channel")
	defer rwin.Close()
	rwin.IMShow(bgr[2])

	for {
		c := gocv.WaitKey(20)
		if c == 27 {
			break
		}
	}

}
