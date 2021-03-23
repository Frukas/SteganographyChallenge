package steganography

import (
	"fmt"
	"image"
	"log"
	"os"

	"golang.org/x/image/bmp"
)

//Converts a array of byte to a string so the messagem can be read
func byteToString(conv []byte) string {
	var result string
	for _, c := range conv {
		result = fmt.Sprint(result, string(c))
	}
	return result
}

//Breaks the bit in to a list of int
func writeListOnByte(letter *byte, list []uint8) {

	for index, l := range list {
		if l == 1 {
			*letter = flipLastByte(*letter)
		}
		if index < 7 {
			*letter = *letter << 1
		}
	}
}

//GetHiddenMessage is the exported function responsible to start the process
func GetHiddenMessage(pathfile string) string {
	imgFile, err := os.Open(pathfile)
	if err != nil {
		log.Fatal(err)
	}

	img, err := bmp.Decode(imgFile)
	if err != nil {
		log.Fatal("Erro no Decode")
	}

	return getHiddenMessageFromImage(img)
}

//Reads the last bit of the alpha of the image, putting together the message and returning .
func getHiddenMessageFromImage(img image.Image) string {

	var msg []byte
	indeX := img.Bounds().Min.X
	indeY := img.Bounds().Min.Y

	for {
		letter := byte(0)
		var list []uint8
		for j := 0; j < 8; j++ {

			tempColor := img.At(indeX, indeY)
			_, _, _, a := tempColor.RGBA()
			list = append(list, getLastByte(byte(a/257)))
			indeX++
		}

		writeListOnByte(&letter, list)
		list = nil
		msg = append(msg, letter)

		if letter == byte(46) {
			return byteToString(msg)
		}
		if indeX > img.Bounds().Max.X {
			indeY++
		}
		if indeY > img.Bounds().Max.Y {
			fmt.Println("Out of bound - not found")
			return "Not found"
		}
	}
}
