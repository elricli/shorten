package snowflake

import (
	"testing"
)

func TestNewIDWorker(t *testing.T) {
	tests := []struct {
		name     string
		workerID int64
		wantErr  bool
	}{
		{"normal", 1, false},
		{"too large workerID", 10, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewIDWorker(tt.workerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIDWorker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
