package utilities

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func CreateZipArchive(uid string) error {
	dirPath := fmt.Sprintf("./processing/%s", uid)
	entries, readDirError := os.ReadDir(dirPath)
	if readDirError != nil {
		return readDirError
	}
	if len(entries) == 0 {
		return nil
	}

	archive, fileError := os.Create(fmt.Sprintf("%s/archive.zip", dirPath))
	if fileError != nil {
		log.Fatal(fileError)
	}
	defer archive.Close()

	zipWriter := zip.NewWriter(archive)
	for _, entry := range entries {
		partials := strings.Split(entry.Name(), ".")
		extention := partials[len(partials)-1]
		if !entry.IsDir() && strings.ToLower(extention) == "jpg" {
			file, fileError := os.Open(fmt.Sprintf("%s/%s", dirPath, entry.Name()))
			if fileError == nil {
				writer, writerError := zipWriter.Create(entry.Name())
				if writerError == nil {
					if _, err := io.Copy(writer, file); err != nil {
						panic(err)
					}
				}
			}
			defer file.Close()
		}
	}
	zipWriter.Close()
	return nil
}
