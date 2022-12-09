package main

import (
	"golang.org/x/image/draw"
	"image"
	"image/png"
	"os"
	"path/filepath"
)

func main() {
	filePath := "io/img.png"

	// ファイルの拡張値を取得
	ext := filepath.Ext(filePath)

	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	i, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	newImage := image.NewRGBA(image.Rect(0, 0, 300, 300))

	draw.CatmullRom.Scale(newImage, newImage.Bounds(), i, i.Bounds(), draw.Over, nil)
	newFile, err := os.Create("io/output" + ext)
	if err != nil {
		panic(err)
	}
	defer func(newFile *os.File) {
		err := newFile.Close()
		if err != nil {
			panic(err)
		}
	}(newFile)

	if err := png.Encode(newFile, newImage); err != nil {
		panic(err)
	}
}
