package database

type QueueEntry struct {
	CreatedAt        int64  `json:"createdAt"`
	Id               string `json:"_id"`
	OriginalFileName string `json:"originalFileName"`
	Status           string `json:"status"`
	UID              string `json:"uid"`
	UpdatedAt        int64  `json:"updatedAt"`
}
