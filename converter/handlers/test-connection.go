package handlers

func TestConnection(timestamp int64) (*TestConnectionResult, error) {
	return &TestConnectionResult{Timestamp: timestamp}, nil
}
