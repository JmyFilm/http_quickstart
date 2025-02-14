package xlog

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

// logging level等级标签 stack堆栈信息 errPostscript错误附言 err错误信息
func logging(logType *log.Logger, level string, stack []byte, postscript string, err ...any) {
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
		logType.Printf("[%s] %s < %s | %s:%d\n%s\n", level, postscriptWithErr, funcName, file, line, stack)
	} else {
		logType.Printf("[%s] %s < %s | %s:%d\n", level, postscriptWithErr, funcName, file, line)
	}

	if level == fatalTag {
		for _, fn := range waitQuitFn {
			fn()
		}
		os.Exit(1)
	}
}
