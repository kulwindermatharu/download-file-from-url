package main

import (
	"io"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	TempFileName := randSeq(20)
	URL := "File URL Path"
	//FileExt is Download file Extention
	FileExt := ""
	err := DownloadFile(TempFileName+FileExt, URL)
	if err != nil {
		panic(err)
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//DownloadFile ...
func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
