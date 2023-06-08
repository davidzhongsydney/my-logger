package mylogger

import "log"

type Log_v2 struct {
}

func NewLog_v2() *Log_v2 {
	return &Log_v2{}
}

func (l *Log_v2) LogInfo_v2(message string) {
	log.Printf("v2 INFO - %v", message)
}
