package ocelogger

import (
	"fmt"
	"log"
)

func Info(msg string) {
	log.Println(fmt.Sprintf("level: INFO { msg: %v }", msg))
}

func Infof(key string, value interface{}) {
	log.Println(fmt.Sprintf("level: INFO { %v: %v } ", key, value))
}

func Error(err error) {
	log.Println(fmt.Sprintf("level: ERROR { msg: %v }", err.Error()))
}

func Warn(err error) {
	log.Println(fmt.Sprintf("level: WARN { msg: %v }", err.Error()))
}

func Errorf(msg string, err error) {
	log.Println(fmt.Sprintf("level: ERROR { %v error: %v } ", msg, err.Error()))
}