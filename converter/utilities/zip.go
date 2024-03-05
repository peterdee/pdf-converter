package utilities

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"converter/constants"
)

func CreateZipArchive(uid string) error {
	dirPath := fmt.Sprintf("./processing/%s", uid)
	entries, readDirError := os.ReadDir(dirPath)
	if readDirError != nil {
		return readDirError
	}
	if len(entries) == 0 {
		return errors.New(constants.ERROR_MESSAGES.TargetDirectoryIsEmpty)
	}

	archive, fileError := os.Create(fmt.Sprintf("%s/%s.zip", dirPath, uid))
	if fileError != nil {
		log.Fatal(fileError)
	}
	defer archive.Close()

	zipWriter := zip.NewWriter(archive)
	for _, entry := range entries {
		partials := strings.Split(entry.Name(), ".")
		if len(partials) < 2 {
			continue
		}

		extention := partials[len(partials)-1]
		if !(!entry.IsDir() && strings.ToLower(extention) == "jpg") {
			continue
		}

		file, fileError := os.Open(fmt.Sprintf("%s/%s", dirPath, entry.Name()))
		if fileError != nil {
			continue
		}
		defer file.Close()

		writer, writerError := zipWriter.Create(entry.Name())
		if writerError != nil {
			continue
		}

		if _, err := io.Copy(writer, file); err != nil {
			continue
		}
	}

	zipWriter.Close()
	return nil
}
