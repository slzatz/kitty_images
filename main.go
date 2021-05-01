package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

var path = "/home/slzatz/Pictures/wood_ducks_smaller.jpg"

func getFile(fpath string) (*os.File, int64, error) {

	pF, err := os.Open(fpath)
	if err != nil {
		return nil, 0, err
	}

	fInf, err := pF.Stat()
	if err != nil {
		pF.Close()
		return nil, 0, err
	}

	return pF, fInf.Size(), nil
}
func loadImage(path string) (iImg image.Image, imgFmt string, E error) {

	pF, err := os.Open(path)
	if err != nil {
		return
	}
	defer pF.Close()

	return image.Decode(pF)
}

func main() {
	b := IsTermKitty()
	fmt.Printf("Terminal is kitty: %t\n", b)
	//fIn, nImgLen, err := getFile(fpath)
	iImg, _, err := loadImage(path)
	if err != nil {
		log.Fatal(err)
	}
	err = KittyWriteImage(os.Stdout, iImg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("")

}
