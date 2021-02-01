package snowflake

import (
	"testing"
)

var (
	worker *IDWorker
)

func TestMain(m *testing.M) {
	worker = NewIDWorker(0x1, 0x2)
	m.Run()
}

func TestIDWorker_NextID(t *testing.T) {
	_, err := worker.NextID()
	if err != nil {
		t.Fatal(err)
	}
}
