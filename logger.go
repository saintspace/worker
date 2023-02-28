package main

import (
	"encoding/json"
	"log"
)

type logMessage struct {
	Level   string                 `json:"level"`
	Message string                 `json:"message"`
	Params  map[string]interface{} `json:"params"`
}

func logInfo(message string, params map[string]interface{}) {
	logJSON("INFO", message, params)
}

func logError(message string, params map[string]interface{}) {
	logJSON("ERROR", message, params)
}

func logJSON(level, message string, params map[string]interface{}) {
	logData := logMessage{
		Level:   level,
		Message: message,
		Params:  params,
	}

	logJSON, err := json.Marshal(logData)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(logJSON))
}
