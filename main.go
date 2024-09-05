package main

import (
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

func takeScreenShot(elapse int) {
	displayBound := 0

	bounds := screenshot.GetDisplayBounds(displayBound)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	fileName := fmt.Sprintf("%d_%dx%d_%d.png", displayBound, bounds.Dx(), bounds.Dy(), elapse)
	file, _ := os.Create(fileName)
	defer file.Close()
	png.Encode(file, img)

	fmt.Printf("#%d : %v \"%s\"\n", displayBound, bounds, fileName)
}

func main() {
	elapse := 0

	for {
		takeScreenShot(elapse)
		elapse += 1
		time.Sleep(1 * time.Second)
	}

}
