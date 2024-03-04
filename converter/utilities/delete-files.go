package utilities

import (
	"fmt"
	"os"
)

func DeleteFiles(uid string) error {
	dirPath := fmt.Sprintf("./processing/%s", uid)
	files, dirError := os.ReadDir(dirPath)
	if dirError != nil {
		return dirError
	}
	for _, file := range files {
		if file.Name() == fmt.Sprintf("%s.zip", uid) {
			continue
		}
		os.Remove(fmt.Sprintf("%s/%s", dirPath, file.Name()))
	}
	return nil
}
