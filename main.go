package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	processLink()
}
func viewlink() string {
	if len(os.Args) < 2 {
		return "No url provided"
	}
	url := os.Args[1]
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	button := doc.Find(".plugin-download.button.download-button.button-large")
	href, exists := button.Attr("href")
	if exists {
		//fmt.Println("Enlace encontrado:", href)
	} else {
		//fmt.Println("No se encontró el enlace")
	}
	return href

}
func processLink() {
	url := viewlink()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	filename := path.Base(url)
	out, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	r, err := zip.OpenReader(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	for _, f := range r.File {
		fmt.Println(f.Name)
	}

	errx := os.Remove(filename)
	if errx != nil {
		fmt.Println(err)
		return
	}
}
