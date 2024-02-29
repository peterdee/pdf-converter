package handlers

import "github.com/julyskies/gohelpers"

func DownloadArchive(uid string) (DownloadArchiveResult, error) {
	// TODO: finalize
	response := DownloadArchiveResult{
		Bytes:       "bytes",
		Filename:    "filename.pdf",
		ProcessedAt: gohelpers.MakeTimestamp(),
		UID:         "uid",
	}
	return response, nil
}
