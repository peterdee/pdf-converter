package handlers

func TestConnection(timestamp int64) *TestConnectionResult {
	return &TestConnectionResult{Timestamp: timestamp}
}
