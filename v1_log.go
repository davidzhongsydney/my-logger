package mylogger

import "log"

type Log_v1 struct {
}

func NewLog_v1() *Log_v1 {
	return &Log_v1{}
}

func (l *Log_v1) LogInfo_v1(message string) {
	log.Printf("v1 INFO - %v", message)
}