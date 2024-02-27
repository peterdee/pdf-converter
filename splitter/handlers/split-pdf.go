package handlers

import "encoding/hex"

func SplitPDF(encoded string, filename string) (string, error) {
	_, decodeError := hex.DecodeString(encoded)
	if decodeError != nil {
		return "", decodeError
	}

	return "id", nil
}
