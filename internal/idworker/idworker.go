package idworker

import "github.com/drrrMikado/shorten/pkg/snowflake"

var (
	worker = snowflake.NewIDWorker(1, 1)
	Get    = worker.NextID
)
