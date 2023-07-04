package main

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
)

const pnt_path = "/Users/rotaercw/Downloads/test-image.jpeg"

func pngreader() (io.Reader, error) {
	data, e := os.ReadFile(pnt_path)
	if e != nil {
		return nil, e
	}

	reader := bytes.NewReader(data)
	//config, format, err := image.DecodeConfig(reader)
	//if err != nil {
	//}
	//fmt.Println("Width:", config.Width, "Height:", config.Height, "Format:", format)
	//

	img, str, _ := image.Decode(reader)
	fmt.Println(str)

	//blue := color.RGBA{0, 0, 255, 255}
	m := image.NewRGBA(img.Bounds())
	zp := m.Bounds().Min
	//rect := image.Rectangle{image.Point {10, 10}, image.Point{100,100}}
	draw.Draw(m, img.Bounds(), img, zp, draw.Src)
	imaging.Save(m,"1.png")

	return nil, nil

}


func main() {
	pngreader()
}
