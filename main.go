package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	steganography "github.com/frukas/SteganographyChallenge/Steganography"
)

type message struct {
	Path    string `json:"Path"`
	Message string `json:"Message"`
}

type responseJSON struct {
	FilePath string `json:"FilePath"`
}

//TODO break the function in smalls ones. Saves a copy of the chosen file on the c:\imagens
func uploadGatinhos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Fprintf(w, "Salvando arquivo...\n")

		//creates the folder c:\imagens if it does not exist. TODO let the user set the folder name.
		if _, err := os.Stat("c:\\imagens"); os.IsNotExist(err) {
			if err := os.Mkdir("c:\\imagens", 0755); err != nil {
				log.Fatal("It was not possible create the folder:", err)
			}
		}

		r.ParseMultipartForm(10 << 20)

		file, handler, err := r.FormFile("newFile")
		if err != nil {
			fmt.Println("Erro ao fazer upload", err)
			return
		}
		defer file.Close()

		//TODO check the file is BMP
		fmt.Printf("Arquivo salvo: %+v\n", handler.Filename)

		tempFile, err := ioutil.TempFile("c:\\imagens", "upload-*.bmp")
		if err != nil {
			fmt.Println("Erro ao criar arquivo temporario: ", err)
		}

		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println("Erro ao converter arquivo em bytes", err)
			return
		}

		tempFile.Write(fileBytes)

		fmt.Fprintf(w, "Nome do arquivo: %+v\n", tempFile.Name())

		fmt.Fprintf(w, "Arquivo Salvo com Sucesso\n")
	}
}

//Create a copy of the image with the alter image and returns the file name
func writeMessageOnImage(w http.ResponseWriter, r *http.Request) {
	var msg message
	if r.Method == "POST" {

		messages, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error receiving the JSON", err)
		}

		err = json.Unmarshal(messages, &msg)
		if err != nil {
			fmt.Println("Error converting the JSON file", err)
		}

		fmt.Println("Path:", msg.Path, "Message:", msg.Message)
	}

	filePath := fmt.Sprint("c:\\imagens\\", msg.Path)

	modifiedFile := steganography.WriteHiddenMessage(filePath, msg.Message)

	res := responseJSON{modifiedFile}

	resJSON, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(resJSON)
}

//To download the request image in bmp format.
func getImage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		fileName := r.URL.Query()["FileName"]

		filePath := fmt.Sprint("c:\\imagens\\", fileName[0])
		copyFileName := fmt.Sprint("attachment; filename=", fileName[0])

		w.Header().Set("Content-Disposition", copyFileName)
		w.Header().Set("Content-Type", "image/bmp")
		http.ServeFile(w, r, filePath)
	}
}

//Return the secret message of the chosen file.
func getHiddenMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fileName := r.URL.Query()["FileName"]
		filePath := fmt.Sprint("c:\\imagens\\", fileName[0])

		msg := steganography.GetHiddenMessage(filePath)
		fmt.Fprintf(w, "Hidden Message: %s", msg)
	}
}

//The default port is 8080.
func setupRoutes() {
	http.HandleFunc("/post", uploadGatinhos)
	http.HandleFunc("/write-message-on-image", writeMessageOnImage)
	http.HandleFunc("/get-image", getImage)
	http.HandleFunc("/decode-message-from-image", getHiddenMessage)

	fmt.Println("Server started at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("It was not possible start the server: ", err)
	}

}

func main() {

	setupRoutes()
}
