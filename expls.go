package main

import (
	"fmt"
	"io"
	"os"
	"time"
	//	"math"
	"image"
	"image/png"
	"math/rand"
	//	"image/draw"

	//	"github.com/llgcode/draw2d/draw2dimg"
	//	"github.com/llgcode/draw2d/draw2dkit"

	_ "image/gif"
	_ "image/jpeg"
)

const (
	ITERSDIR = "exp_iters"
)

var (
	rando = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func init() {
	os.Mkdir(ITERSDIR, 0777)
}

func main() {
	fmt.Println("Experiment. Please ignore.")

	sourceImagePath := os.Args[1]
	readPNG(sourceImagePath)
	//loadAndCreateSourceImage(sourceImagePath)
}

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// Load source image from path arg
// Return PNG format image file
func loadAndCreateSourceImage(source string) {
	fmt.Println(source)

	var formattedSourceFile io.Writer

	sourceFile, err := os.Open(source)
    check(err)

	convertToPNG(formattedSourceFile, sourceFile)
}

// convertToPNG converts from any recognized format to PNG.
func convertToPNG(w io.Writer, r io.Reader) error {
	img, _, err := image.Decode(r)
    check(err)

    return png.Encode(w, img)
}

func readPNG(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return png.Decode(f)
}