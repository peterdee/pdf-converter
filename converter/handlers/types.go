package handlers

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

type GetQueueResult struct {
	JSON string
}

type QueueFileResult struct {
	Filename    string
	QueuedItems int64
	Status      string
	UID         string
}
