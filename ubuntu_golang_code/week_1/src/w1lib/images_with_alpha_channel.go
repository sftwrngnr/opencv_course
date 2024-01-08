package w1lib

import (
	"fmt"
	"image"
	"log"
	"os"

	_ "image/png"

	"gocv.io/x/gocv"
)

func ImagesWithAlphaChannel() {
	// Path of the image to be loaded
	// Here we are supplying a relative path
	imageFile := DATA_PATH + "/images/panther.png"

	// Read the image using raw image read
	inFile, err := os.Open(imageFile)
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()
	img, _, err := image.Decode(inFile)
	if err != nil {
		log.Fatal(err)
	}
	matData, err := gocv.ImageToMatRGBA(img)
	fmt.Printf("Number of channels is %d\n", matData.Channels())
	chandata := gocv.Split(matData)
	Combo := gocv.NewMat()
	gocv.Merge([]gocv.Mat{chandata[0], chandata[1], chandata[2]}, &Combo)
	combowin := gocv.NewWindow("RGB Image (no Alpha)")
	defer combowin.Close()
	combowin.IMShow(Combo)
	maskwin := gocv.NewWindow("Alpha Channel")
	defer maskwin.Close()
	maskwin.IMShow(chandata[3])
	origwin := gocv.NewWindow("Original Pic")
	defer origwin.Close()
	origwin.IMShow(matData)
	// Now write data
	gocv.IMWrite("results/colorChannels.png", Combo)
	gocv.IMWrite("results/alphaChannel.png", chandata[3])

	for {
		c := gocv.WaitKey(20)
		if c == 27 {
			break
		}
	}
}
