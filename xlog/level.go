package xlog

import (
	"PROJECTNAME/conf"
	"fmt"
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
	infoLog.Println(append([]any{fmt.Sprint("[", debugTag, "]")}, info...)...)
}

func Info(info ...any) {
	infoLog.Println(append([]any{fmt.Sprint("[", infoTag, "]")}, info...)...)
}

func Err(postscript string, err ...any) {
	logging(errorLog, errorTag, nil, postscript, err...)
}

func ErrWithStack(postscript string, err ...any) {
	logging(errorLog, errorTag, debug.Stack(), postscript, err...)
}

func Fatal(postscript string, err ...any) {
	logging(errorLog, fatalTag, nil, postscript, err...)
}

func FatalWithStack(postscript string, err ...any) {
	logging(errorLog, fatalTag, debug.Stack(), postscript, err...)
}
