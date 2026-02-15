package logger

import (
	"encoding/json"
	"lab/internal/models"
	"log"
	"os"
)

func LogInfo(message, requestID string, additional ...interface{}) {
	var logger = log.New(os.Stdout, "", log.LstdFlags)
	entry := models.LogEntry{
		Level:     "INFO",
		Message:   message,
		RequestID: requestID,
		Function:  os.Getenv("AWS_LAMBDA_FUNCTION_NAME"),
	}
	data, _ := json.Marshal(entry)
	logger.Println(string(data))
}
