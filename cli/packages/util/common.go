package util

import (
	"fmt"
	"os"
)

const (
	CONFIG_FILE_NAME                     = "infisical-config.json"
	CONFIG_FOLDER_NAME                   = ".infisical"
	INFISICAL_WORKSPACE_CONFIG_FILE_NAME = ".infisical.json"
	INFISICAL_SERVICE_TOKEN              = "INFISICAL_SERVICE_TOKEN"
)

var INFISICAL_URL = "https://api.infisical.com"

func GetHomeDir() (string, error) {
	directory, err := os.UserHomeDir()
	return directory, err
}

func WriteToFile(fileName string, dataToWrite []byte, filePerm os.FileMode) error {
	err := os.WriteFile(fileName, dataToWrite, filePerm)
	if err != nil {
		return fmt.Errorf("Unable to wrote to file", err)
	}

	return nil
}