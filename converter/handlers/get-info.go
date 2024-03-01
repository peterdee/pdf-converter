package handlers

func GetInfo(uid string) (*GetInfoResult, error) {
	response := &GetInfoResult{
		Filename:    "file.pdf",
		QueuedItems: 1,
		Status:      "status",
		UID:         "uid",
	}
	return response, nil
}
