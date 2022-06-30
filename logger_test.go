package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestDefaultLogger(test *testing.T) {
	logger := NewDefaultLogger()
	defer cleanup(logger.GetLogFilePath())

	entry1 := "Test log entry to file only"
	entry2 := "Test log entry to file and console"
	entry3 := "Test log entry with Logf : 12345"

	logger.Log(entry1, false)
	logger.Log(entry2, true)
	logger.Logf(false, "%s : %d", "Test log entry with Logf", 12345)

	logger.Close()

	if _, pathErr := os.Stat(logger.GetLogFilePath()); pathErr != nil {
		test.Error(pathErr)
	}

	var logFile []byte
	var logFileErr error
	if logFile, logFileErr = ioutil.ReadFile(logger.GetLogFilePath()); logFileErr != nil {
		test.Error(logFileErr)
	}

	actualContent := string(logFile)
	expectedContent := fmt.Sprintf("%s\n%s\n%s\n", entry1, entry2, entry3)

	if actualContent != expectedContent {
		test.Errorf("Unexpected log file contents: %s", actualContent)
	}
}

func cleanup(filename string) {
	err := os.Remove(filename)
	if err != nil {
		fmt.Printf("Unable to delete log file %s\n", filename)
		fmt.Println(err.Error())
	}
}
