package server

import "errors"

const (
	// LogRoot is the default location where frenzy will store the logs
	LogRoot = "/var/log/frenzy"
)

var (
	// ErrStatNotFound occurs when there is error while reading the status from
	// an folder
	ErrStatNotFound = errors.New("unable to get the folder stats")
	// ErrMakeDir signify error during creating a folder
	ErrMakeDir = errors.New("unable to create the folder")
)
