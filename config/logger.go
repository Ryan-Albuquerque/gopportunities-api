package config

import (
	"fmt"
	"log"
)

func LogError(context string, format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	log.Printf("[Error][%s] - %s",  context, message)
}

func LogInfo(context string, format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	log.Printf("[Info][%s] - %s",  context, message)
}