package main

import (
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func pdf2txt(pdfDir, txtDir string) {
	_, err := os.Stat(txtDir)
	if err != nil {
		err = os.Mkdir(txtDir, fs.ModePerm)
		if err != nil {
			log.Fatalf("cannot create txt dir: %v\n", err)
		}
	}

	have, err := os.ReadDir(txtDir)
	if err != nil {
		log.Fatalf("%s: cannot read dir: %v\n", txtDir, err)
	}

	files, err := os.ReadDir(pdfDir)
	if err != nil {
		log.Fatalf("%s: cannot read dir: %v\n", pdfDir, err)
	}

loop:
	for _, f := range files {
		for _, h := range have {
			if strings.Compare(h.Name(), f.Name()) == 0 {
				log.Printf("%s was already existed\n", f.Name())
				continue loop
			}
		}

		pathPdf := filepath.Join(pdfDir, f.Name())
		pathTxt := filepath.Join(txtDir, f.Name()+".txt")

		for retry := 0; retry < 100; retry++ {
			log.Println("processing: pdftotext", pathPdf, pathTxt)
			if err := exec.Command("pdftotext", pathPdf, pathTxt).Run(); err != nil {
				log.Printf("failed to convert %s from PDF to text: %v\n", pathPdf, err)
				continue
			}
			break
		}

		if _, err := os.Stat(pathTxt); err != nil {
			log.Printf("there was a problem with parsing %s to text, creating an empty text file.\n", pathPdf)
		}
	}
}

func main() {
	pdf2txt("data/pdf", "data/txt")
}
