package server

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// CreateFolderIfNotExist will create a folder if it's not already exist
func CreateFolderIfNotExist(topic string) (string, error) {
	fullPath := filepath.Join(LogRoot, topic)
	fmt.Println(fullPath)
	if _, err := os.Stat(fullPath); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(fullPath, os.ModePerm); err != nil {
				fmt.Printf("%s", err)
				return "", ErrMakeDir
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
func ReadLogFile(topic string) chan string {
	msgStream := make(chan string)

	go func() {
		file, err := os.Open(filepath.Join(LogRoot, topic, getLogFileName()))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			msgStream <- line
		}
	}()

	return msgStream
}
