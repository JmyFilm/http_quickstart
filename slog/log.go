package slog

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

// doLog level等级标签 stack堆栈信息 errPostscript错误附言 err错误信息
func doLog(level string, stack []byte, postscript string, err ...any) {
	if err == nil || len(err) == 0 || (len(err) == 1 && err[0] == nil) {
		return
	}

	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		file = ""
		line = 0
	}
	funcName := runtime.FuncForPC(pc).Name()

	var postscriptWithErr string
	if len(postscript) != 0 {
		postscriptWithErr = strings.TrimRight(fmt.Sprintln(append([]any{postscript, ">"}, err...)...), "\n")
	} else {
		postscriptWithErr = strings.TrimRight(fmt.Sprintln(err...), "\n")
	}

	if stack != nil && len(stack) != 0 {
		log.Printf("[%s] %s %s %s:%d\n%s\n", level, postscriptWithErr, funcName, file, line, stack)
	} else {
		log.Printf("[%s] %s %s %s:%d\n", level, postscriptWithErr, funcName, file, line)
	}

	if level == fatalTag {
		os.Exit(1)
	}
}
