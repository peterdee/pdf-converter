package queue

type QueueEntry struct {
	CreatedAt        int64  `json:"createdAt"`
	DownloadedAt     int64  `json:"downloadedAt"`
	ID               string `bson:"_id" json:"id,omitempty"`
	OriginalFileName string `json:"originalFileName"`
	Status           string `json:"status"`
	TotalDownloads   int    `json:"totalDownloads"`
	UID              string `json:"uid"`
	UpdatedAt        int64  `json:"updatedAt"`
}

type GetQueueJSON struct {
	Items      []QueueEntry `json:"items"`
	Limit      int          `json:"limit"`
	Page       int          `json:"page"`
	TotalItems int          `json:"totalItems"`
	TotalPages int          `json:"totalPages"`
}
