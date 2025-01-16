package xlog

import (
	"PROJECTNAME/conf"
	"fmt"
	"log"
	"runtime/debug"
)

const debugTag = "DEBUG"
const infoTag = "INFO"
const errorTag = "ERROR"
const fatalTag = "FATAL"

func Debug(info ...any) {
	if !conf.Log.DebugInfo {
		return
	}
	log.Println(append([]any{fmt.Sprint("[", debugTag, "]")}, info...)...)
}

func Info(info ...any) {
	log.Println(append([]any{fmt.Sprint("[", infoTag, "]")}, info...)...)
}

func Err(postscript string, err ...any) {
	logging(errorTag, nil, postscript, err...)
}

func ErrWithStack(postscript string, err ...any) {
	logging(errorTag, debug.Stack(), postscript, err...)
}

func Fatal(postscript string, err ...any) {
	logging(fatalTag, nil, postscript, err...)
}

func FatalWithStack(postscript string, err ...any) {
	logging(fatalTag, debug.Stack(), postscript, err...)
}
