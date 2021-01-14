package server

import (
	"log"
	"os"
	"path/filepath"

	"github.com/iAziz786/frenzy/constant"
)

// CreateFolderIfNotExist will create a folder if it's not already exist
func CreateFolderIfNotExist(topic string) (string, error) {
	fullPath := filepath.Join(constant.LogRoot, topic)
	if _, err := os.Stat(fullPath); err != nil {
		if os.IsNotExist(err) {
			log.Println("creating new folder for topic:", topic)
			if err := os.MkdirAll(fullPath, os.ModePerm); err != nil {
				log.Printf("unable to create directory %s\n", err)
				return "", constant.ErrMakeDir
			}
		}

		return "", err
	}

	return fullPath, nil
}

func getLogFileName() string {
	return "00000000000.frenzy"
}

// CreateLogFile will create a log file in LogRoot folder. If the log file
// already exists it will open the file.
func CreateLogFile(topic string) (*os.File, error) {
	fullPath, err := CreateFolderIfNotExist(topic)

	fullFilePath := filepath.Join(fullPath, getLogFileName())

	if err != nil {
		return nil, err
	}

	return os.OpenFile(fullFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}

// ReadLogFile will return a channel where it will send the message of the
// topic
func ReadLogFile(topic string) chan []byte {
	msgStream := make(chan []byte)

	go func() {
		file, err := os.Open(filepath.Join(constant.LogRoot, topic, getLogFileName()))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		for {
			if b, err := Poll(file); err == nil && len(b) > 0 {
				msgStream <- b
			}
		}
	}()

	return msgStream
}
