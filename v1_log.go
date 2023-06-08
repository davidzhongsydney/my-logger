package mylogger

import "log"

type log_v1 struct {
}

func NewLog_v1() *log_v1 {
	return &log_v1{}
}

func (l *log_v1) LogInfo_v1(message string) {
	log.Printf("v1 INFO - %v", message)
}
