// dead code
package cmd

import (
	"crypto/rand"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

// integer to string array
func ItoSA(t []interface{}) []string {
	s := make([]string, len(t))
	for i, v := range t {
		s[i] = fmt.Sprint(v)
	}
	return s
}

func CreateImage(uuid_string string) string {
	rect := image.Rect(0, 0, 100, 100)
	pix := make([]uint8, rect.Dx()*rect.Dy()*4)
	rand.Read(pix)
	created_image := &image.NRGBA{
		Pix:    pix,
		Stride: rect.Dx() * 4,
		Rect:   rect,
	}

	file_name := fmt.Sprint(uuid_string + ".png")
	SaveImage(file_name, created_image)
	return file_name
}

func SaveImage(fileName string, img *image.NRGBA) {
	basePath := os.Getenv("RUN_PATH") + "/assets/"
	imgFile, err := os.Create(basePath + filepath.Base(fileName))
	defer imgFile.Close()
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	png.Encode(imgFile, img.SubImage(img.Rect))
}
