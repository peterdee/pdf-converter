package handlers

import "converter/database"

type DeleteEntryResult struct {
	Deleted bool
}

type DownloadArchiveResult struct {
	Bytes       string
	Filename    string
	ProcessedAt int64
	UID         string
}

type GetInfoResult struct {
	DownloadedAt   int64
	Filename       string
	QueuedItems    int64
	Status         string
	TotalDownloads int
	UID            string
}

type GetQueueJSON struct {
	Items      []database.QueueEntry `json:"items"`
	Limit      int                   `json:"limit"`
	Page       int                   `json:"page"`
	TotalItems int                   `json:"totalItems"`
	TotalPages int                   `json:"totalPages"`
}

type GetQueueResult struct {
	JSON string
}

type QueueFileResult struct {
	Filename    string
	QueuedItems int64
	Status      string
	UID         string
}
