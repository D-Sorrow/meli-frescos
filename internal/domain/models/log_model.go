package models

import "time"

type LogLevel string

const (
	INFO    LogLevel = "INFO"
	ERROR   LogLevel = "ERROR"
	WARNING LogLevel = "WARNING"
	DEBUG   LogLevel = "DEBUG"
)

type LogAttributes struct {
	Level  LogLevel
	Source string
	Detail string
}

type Log struct {
	ID        int64
	Timestamp time.Time
	LogAttributes
}
