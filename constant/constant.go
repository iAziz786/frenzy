package constant

import "errors"

const (
	// LogRoot is the default location where frenzy will store the logs
	LogRoot = "/var/log/frenzy"

	// MaxPollWait waits for the specified milliseconds after the threshold
	// it will flush the data
	MaxPollWait = 1000

	// MaxBufData is buffer data which will either be filled to flush the data
	// to a writer
	MaxBufData = 1e6

	// MsgEndIdent is the identifier which indicates the end of the message in
	// array of bytes
	MsgEndIdent = '\n'
)

var (
	// ErrStatNotFound occurs when there is error while reading the status from
	// an folder
	ErrStatNotFound = errors.New("unable to get the folder stats")
	// ErrMakeDir signify error during creating a folder
	ErrMakeDir = errors.New("unable to create the folder")
)
