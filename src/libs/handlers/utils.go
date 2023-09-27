package handlers

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"

	crud "github.com/Mecuate/crud_module"
	"github.com/gorilla/mux"
)

func StabilizeRouter(r *mux.Router) crud.MuxRouter {
	router := crud.MuxRouter{
		Router: r,
	}
	return router
}

func validateMediaFormattingURLString(current string) bool {
	eval, _ := regexp.MatchString(`\d{2,4}[xX]\d{2,4}`, current)
	if eval == true {
		return true
	} else {
		return false
	}
}

func ConvertJPEGToPNG(w io.Writer, r io.Reader) error {
	img, err := jpeg.Decode(r)

	if err != nil {
		return err
	}
	return png.Encode(w, img)
}

func ConvertPNGToJPEG(w io.Writer, r io.Reader) error {
	img, err := png.Decode(r)

	if err != nil {
		return err
	}
	return jpeg.Encode(w, img, &jpeg.Options{Quality: 80})
}

func getImageFromFilePath(fileLocation string) (*os.File, error) {
	mydir, err := os.Getwd()
	filePath := fmt.Sprintf("%s%s", mydir, fileLocation)
	f, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f, err
}

func ResImg(file_addr string, x int, y int) []byte {
	f, err := os.Open(file_addr)

	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(f)

	if err != nil {
		log.Fatal(err)
	}

	resImg := resize(img, x, y)
	imgBytes := imgToBytes(resImg)

	return imgBytes
}

func SaveImage(imgBytes []byte, filepath string) {
	err := ioutil.WriteFile(filepath, imgBytes, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func resize(img image.Image, length int, width int) image.Image {
	minX := img.Bounds().Min.X
	minY := img.Bounds().Min.Y
	maxX := img.Bounds().Max.X
	maxY := img.Bounds().Max.Y
	for (maxX-minX)%length != 0 {
		maxX--
	}
	for (maxY-minY)%width != 0 {
		maxY--
	}
	scaleX := (maxX - minX) / length
	scaleY := (maxY - minY) / width

	imgRect := image.Rect(0, 0, length, width)
	resImg := image.NewRGBA(imgRect)
	draw.Draw(resImg, resImg.Bounds(), &image.Uniform{C: color.White}, image.ZP, draw.Src)
	for y := 0; y < width; y += 1 {
		for x := 0; x < length; x += 1 {
			averageColor := getAverageColor(img, minX+x*scaleX, minX+(x+1)*scaleX, minY+y*scaleY, minY+(y+1)*scaleY)
			resImg.Set(x, y, averageColor)
		}
	}
	return resImg
}

func getAverageColor(img image.Image, minX int, maxX int, minY int, maxY int) color.Color {
	var averageRed float64
	var averageGreen float64
	var averageBlue float64
	var averageAlpha float64
	scale := 1.0 / float64((maxX-minX)*(maxY-minY))

	for i := minX; i < maxX; i++ {
		for k := minY; k < maxY; k++ {
			r, g, b, a := img.At(i, k).RGBA()
			averageRed += float64(r) * scale
			averageGreen += float64(g) * scale
			averageBlue += float64(b) * scale
			averageAlpha += float64(a) * scale
		}
	}

	averageRed = math.Sqrt(averageRed)
	averageGreen = math.Sqrt(averageGreen)
	averageBlue = math.Sqrt(averageBlue)
	averageAlpha = math.Sqrt(averageAlpha)

	averageColor := color.RGBA{
		R: uint8(averageRed),
		G: uint8(averageGreen),
		B: uint8(averageBlue),
		A: uint8(averageAlpha)}

	return averageColor
}

func imgToBytes(img image.Image) []byte {
	buff := bytes.NewBuffer(nil)
	err := jpeg.Encode(buff, img, &jpeg.Options{Quality: 100})

	if err != nil {
		log.Fatal(err)
	}

	return buff.Bytes()
}
