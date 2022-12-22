package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	// for _, alias := range []string{"4"} {
	for _, alias := range []string{"1", "2", "3", "4"} {
		cfgPath := fmt.Sprintf("config/boxes-%s.json", alias)
		imagePath := fmt.Sprintf("input/%s.jpeg", alias)

		st := time.Now()
		fmt.Printf("start to read and recognize image -> %s\n\n", imagePath)
		conf, err := readConfig(cfgPath)
		if err != nil {
			panic(err)
		}

		recognizeImage(imagePath, conf)

		fmt.Printf("\nend image <- %s %v\n\n", imagePath, time.Now().Sub(st))
	}
}

func readConfig(filename string) (BoxesConfig, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var payload BoxesConfig
	err = json.Unmarshal(content, &payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

type BoxesConfig []BoxConfig

type BoxConfig struct {
	Property  string      `json:"property"`
	Lang      []string    `json:"lang"`
	Psm       int         `json:"psm"`
	Gray      bool        `json:"gray"`
	Threshold int         `json:"threshold"`
	Top       Coordinates `json:"top"`
	Bottom    Coordinates `json:"bottom"`
}

type Coordinates struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func recognizeImage(imagePath string, conf BoxesConfig) {
	// fmt.Println("conf:", conf)
	if len(conf) == 0 {
		return
	}

	src, width, height, ext, err := readImage(imagePath)
	if err != nil {
		panic(err)
	}

	client := gosseract.NewClient()
	defer client.Close()
	// client.SetImage(imagePath)

	for _, box := range conf {
		if len(box.Lang) > 0 {
			client.SetLanguage(box.Lang...)
		}
		client.SetPageSegMode(gosseract.PageSegMode(box.Psm))

		x := float64(width) * box.Top.X
		y := float64(height) * box.Top.Y
		w := float64(width)*box.Bottom.X - x
		h := float64(height)*box.Bottom.Y - y

		img, err := cropImage(src, int(x), int(y), int(w), int(h))
		if err != nil {
			panic(err)
		}

		if box.Gray {
			img = grayImage(img, box.Threshold)
			// saveImage("out.jpeg", img, ext)
			// break
		}

		b := &bytes.Buffer{}
		if err := copyImage(bufio.NewWriter(b), img, ext); err != nil {
			panic(err)
		}

		client.SetImageFromBytes(b.Bytes())

		text, err := client.Text()
		if err != nil {
			panic(err)
		}

		fmt.Println(box.Property, "=", text)
	}
}

func readImage(path string) (img image.Image, width, height int, ext string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, ext, err = image.Decode(file)
	if err != nil {
		return
	}

	b := img.Bounds()
	width = b.Max.X
	height = b.Max.Y
	return
}

func cropImage(src image.Image, x, y, w, h int) (image.Image, error) {
	var subImg image.Image
	if rgbImg, ok := src.(*image.YCbCr); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.YCbCr)
	} else if rgbImg, ok := src.(*image.RGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.RGBA)
	} else if rgbImg, ok := src.(*image.NRGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.NRGBA)
	} else {
		return subImg, fmt.Errorf("not support image format")
	}

	return subImg, nil
}

func copyImage(w io.Writer, src image.Image, ext string) (err error) {
	if strings.EqualFold(ext, "jpg") || strings.EqualFold(ext, "jpeg") {
		err = jpeg.Encode(w, src, &jpeg.Options{Quality: 80})
	} else if strings.EqualFold(ext, "png") {
		err = png.Encode(w, src)
	} else if strings.EqualFold(ext, "gif") {
		err = gif.Encode(w, src, &gif.Options{NumColors: 256})
	}
	return err
}

func saveImage(p string, src image.Image, ext string) error {
	f, err := os.OpenFile(p, os.O_SYNC|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	return copyImage(f, src, ext)
}

func grayImage(m image.Image, threshold int) image.Image {
	grayImage := image.NewGray(m.Bounds())
	draw.Draw(grayImage, m.Bounds(), m, m.Bounds().Min, draw.Src)

	if threshold > 0 && threshold < 256 {
		grayImage = binarisationImage(grayImage, uint8(threshold))
	}

	return grayImage
}

// image binarisation and convert to black/white image
func binarisationImage(o *image.Gray, threshold uint8) *image.Gray {
	b := new(image.Gray)
	b.Stride = o.Stride
	b.Rect = o.Rect
	b.Pix = make([]uint8, len(o.Pix))
	copy(b.Pix, o.Pix)

	for i := 0; i < len(b.Pix); i++ {
		if b.Pix[i] > threshold {
			b.Pix[i] = 255
		} else {
			b.Pix[i] = 0
		}
	}

	return b
}
