package snowflake

import (
	"errors"
	"strconv"
	"time"
)

var (
	// Epoch is set to the twitter snowflake epoch of Jun 04 2021 00:00:00 UTC in milliseconds
	Epoch int64 = 1622764800000
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
func NewIDWorker(workerID int64) (*IDWorker, error) {
	var (
		workerIDBits       int64 = 3
		maxWorkerID        int64 = -1 ^ (-1 << workerIDBits)
		sequenceBits       int64 = 3 // Don't need too large sequence
		workerIDShift            = sequenceBits
		timestampLeftShift       = sequenceBits + workerIDBits
		sequenceMask       int64 = -1 ^ (-1 << sequenceBits)
	)
	if workerID < 0 || workerID > maxWorkerID {
		return nil, errors.New("workerID must between 0 and " + strconv.FormatInt(workerID, 10))
	}
	w := &IDWorker{
		epoch:              Epoch,
		workerIDBits:       workerIDBits,
		maxWorkerID:        maxWorkerID,
		sequenceBits:       sequenceBits,
		workerIDShift:      workerIDShift,
		timestampLeftShift: timestampLeftShift,
		sequenceMask:       sequenceMask,
		lastTimestamp:      -1,
		WorkerID:           workerID,
		Sequence:           0,
	}
	return w, nil
}

// NextID .
func (worker *IDWorker) NextID() int64 {
	timestamp := timeGen()

	if worker.lastTimestamp >= timestamp {
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
	return id
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
