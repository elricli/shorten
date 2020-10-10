package generator

import (
	"fmt"
	"time"
)

// IDWorker .
type IDWorker struct {
	epoch              int64
	workerIDBits       int64
	maxWorkerID        int64
	sequenceBits       int64
	workerIDShift      int64
	timestampLeftShift int64
	sequenceMask       int64
	lastTimestamp      int64
	WorkerID           int64
	Sequence           int64
}

// NewIDWorker return a id worker.
func NewIDWorker(workerID, sequence int64) *IDWorker {
	var (
		workerIDBits       int64 = 5
		maxWorkerID        int64 = -1 ^ (-1 << workerIDBits)
		sequenceBits       int64 = 5 // Don't need too large sequence
		workerIDShift      int64 = sequenceBits
		timestampLeftShift int64 = sequenceBits + workerIDBits
		sequenceMask       int64 = -1 ^ (-1 << sequenceBits)
	)
	w := &IDWorker{
		epoch:              1602317796589,
		workerIDBits:       workerIDBits,
		maxWorkerID:        maxWorkerID,
		sequenceBits:       sequenceBits,
		workerIDShift:      workerIDShift,
		timestampLeftShift: timestampLeftShift,
		sequenceMask:       sequenceMask,
		lastTimestamp:      -1,
		WorkerID:           workerID,
		Sequence:           sequence,
	}
	return w
}

// NextID .
func (worker *IDWorker) NextID() (int64, error) {
	timestamp := timeGen()
	if timestamp < worker.lastTimestamp {
		return 0, fmt.Errorf("Clock moved backwards. Refusing to generate id for %d milliseconds", worker.lastTimestamp-timestamp)
	}

	if worker.lastTimestamp == timestamp {
		worker.Sequence = (worker.Sequence + 1) & worker.sequenceMask
		if worker.Sequence == 0 {
			timestamp = tilNextMillis(worker.lastTimestamp)
		}
	} else {
		worker.Sequence = 0
	}
	worker.lastTimestamp = timestamp
	id := ((timestamp - worker.epoch) << worker.timestampLeftShift) |
		(worker.WorkerID << worker.workerIDShift) |
		worker.Sequence
	return id, nil
}

func timeGen() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func tilNextMillis(lastTimestamp int64) int64 {
	timestamp := timeGen()
	for timestamp <= lastTimestamp {
		timestamp = timeGen()
	}
	return timestamp
}
