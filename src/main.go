package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/schollz/progressbar/v3"
)

func main() {
	fmt.Print("Hello World")
}

func downloadFile(url string, fileName string, path string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	file, _ := os.OpenFile(path+fileName, os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)
	io.Copy(io.MultiWriter(file, bar), resp.Body)
	return nil
}

func exitProgram(inputReboot string) {
	if inputReboot != "y" && inputReboot != "Y" && inputReboot != "" {
		os.Exit(1)
	}
}
