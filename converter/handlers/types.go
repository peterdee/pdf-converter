package handlers

type DownloadArchiveResult struct {
	Bytes       string
	Filename    string
	ProcessedAt int64
	UID         string
}

type GetInfoResult struct {
	Filename    string
	QueuedItems int64
	Status      string
	UID         string
}

type QueueFileResult struct {
	Filename    string
	QueuedItems int64
	QueueId     string
	Status      string
}
