package etr

import (
	"github.com/go-openapi/swag"
)

var (
	ExitNormally  *int64
	ExitWithError *int64
)

func init() {
	ExitNormally = swag.Int64(0)
	ExitWithError = swag.Int64(1)
}

// Output is the result of this application
type Output struct {
	ExitCode  *int64
	SyncLogs  map[string]interface{}
	RequestID string     `json:"RequestID,omitempty"`
	Meta      OutputMeta `json:"meta"`
}

// OutputMeta are the set of logs
type OutputMeta struct {
	ExitCodes []OutputExitCodes `json:"5.exitcodes,omitempty"`
}

// OutputExitCodes represent applications status
type OutputExitCodes struct {
	LastStatus string `json:"TaskLastStatus,omitempty"`
}
