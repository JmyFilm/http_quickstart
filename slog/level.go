package slog

import (
	"edit-your-project-name/config"
	"log"
	"runtime/debug"
)

func Debug(info ...any) {
	if !config.Log.DebugInfo {
		return
	}
	log.Println(append([]any{"[DEBUG]"}, info...)...)
}

func Info(info ...any) {
	log.Println(append([]any{"[INFO]"}, info...)...)
}

func Err(postscript string, err ...any) {
	doLog("ERROR", nil, postscript, err...)
}

func ErrWithStack(postscript string, err ...any) {
	doLog("ERROR", debug.Stack(), postscript, err...)
}

func Fatal(postscript string, err ...any) {
	doLog(fatalTag, debug.Stack(), postscript, err...)
}
