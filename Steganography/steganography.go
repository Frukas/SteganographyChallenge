package steganography

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/image/bmp"
)

//Changes the last bit from 1 to 0 or 0 to 1.
func flipLastByte(changeByte byte) byte {
	return changeByte ^ byte(1)
}

//Returns 0 or 1 acording to the last bit.
func getLastByte(lastByte byte) byte {
	return lastByte & byte(1)
}

//Print the the byte as binary. For testing.
func PrintByte(prt byte) {
	f := fmt.Sprintf("% 08b", prt)
	fmt.Println("Forma Binaria:", f)
}

//Returns 0 or 1 according the index position of the byte.
func getIndexByteAsLastByte(indexByte byte, index int) byte {
	newIndex := 7 - index
	newByte := indexByte >> newIndex

	return getLastByte(newByte)
}

//convert a letter(rune) in a list of unit8 equivalent to his binary
func getBinaryChangeList(letter rune) []uint8 {
	var list []uint8
	for i := 0; i < 8; i++ {
		list = append(list, getIndexByteAsLastByte(byte(letter), i))
	}
	return list
}

//Creats a copy of the imagem changing to image.RGBA. The image.Image objet does not have any method.
func imageToRGBA(img image.Image) *image.RGBA {
	rgba := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, img.Bounds().Min, draw.Src)
	return rgba
}

//Changes the last bit of the alpha color.
func hideMessageInAlpha(rgba *image.RGBA, bitChange uint8, color uint8) byte {
	Avalue := byte(color)
	if getLastByte(Avalue) != bitChange {
		return flipLastByte(Avalue)
	}
	return byte(Avalue)
}

//Move pixel to pixel changing the last bit of alpha color writting the text according to msg.
func writeMessageToRGBA(rgba *image.RGBA, msg string) {

	indeX := rgba.Bounds().Min.X
	indeY := rgba.Bounds().Min.Y

	for _, m := range msg {

		list := getBinaryChangeList(m)
		for i := 0; i < 8; i++ {
			newColor := color.RGBA{
				R: rgba.RGBAAt(indeX, indeY).R,
				G: rgba.RGBAAt(indeX, indeY).G,
				B: rgba.RGBAAt(indeX, indeY).B,
				A: hideMessageInAlpha(rgba, list[i], rgba.RGBAAt(indeX, indeY).A),
			}

			if indeY > rgba.Rect.Max.Y {
				fmt.Println("Out of bound -not found")
				return
			}
			if indeX > rgba.Rect.Max.X {
				indeY++
			}
			rgba.SetRGBA(indeX, indeY, newColor)

			indeX++
		}
	}
}

//The exported function that opens the files, decode the imagem for manipulation, and once it is done enconde back and return the file path of the saved image.
func WriteHiddenMessage(pathfile string, msg string) string {
	var buf bytes.Buffer
	imgFile, err := os.Open(pathfile)
	if err != nil {
		log.Fatal(err)
	}

	img, err := bmp.Decode(imgFile)
	if err != nil {
		log.Fatal("Erro no Decode")
	}

	imgRBA := imageToRGBA(img)

	writeMessageToRGBA(imgRBA, msg)

	if err = bmp.Encode(&buf, imgRBA); err != nil {
		log.Fatal("Error enconding BMP:", err)
	}

	tempFile, err := ioutil.TempFile("c:\\imagens", "Modified-*.bmp")
	if err != nil {
		fmt.Println("Error creating temp file: ", err)
	}

	tempFile.Write(buf.Bytes())

	defer tempFile.Close()

	return tempFile.Name()

}
