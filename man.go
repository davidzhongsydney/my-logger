package mylogger

import "log"

func LogInfo(message string) {
	log.Printf("INFO - %v", message)
}

func LogFatal(message string) {
	log.Printf("FATAL - %v", message)
}
