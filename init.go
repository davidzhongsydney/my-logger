package mylogger

import "log"

func LogInfo_outer(message string) {
	log.Printf("INFO - %v", message)
}
