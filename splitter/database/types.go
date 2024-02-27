package database

type QueueEntry struct {
	CreatedAt        int64  `json:"createdAt"`
	Id               string `json:"_id"`
	IsProcessed      bool   `json:"isProcessed"`
	OriginalFileName string `json:"originalFileName"`
	UID              string `json:"uid"`
	UpdatedAt        int64  `json:"updatedAt"`
}
