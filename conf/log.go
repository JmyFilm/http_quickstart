package conf

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"strings"
	"syscall"
)

const fatalTag = "FATAL"

func InitLog() {
	log.SetOutput(io.MultiWriter(&lumberjack.Logger{
		Filename:  Log.FilePath,
		MaxSize:   Log.MaxSize,
		LocalTime: true,
	}, os.Stdout))
	log.Println("APP RUNNING")
}

func Debug(info ...any) {
	if !Log.Debug {
		return
	}
	log.Println(append([]any{"[DEBUG]"}, info...)...)
}

func Info(info ...any) {
	log.Println(append([]any{"[INFO]"}, info...)...)
}

func ErrExt(errPostscript string, err ...any) {
	doErr("ERROR", nil, errPostscript, err...)
}

func ErrWithStackExt(errPostscript string, err ...any) {
	doErr("ERROR", debug.Stack(), errPostscript, err...)
}

func FatalExt(errPostscript string, err ...any) {
	doErr("ERROR", debug.Stack(), errPostscript, err...)
}

func doErr(tag string, stack []byte, errPostscript string, err ...any) {
	if err == nil || len(err) == 0 || (len(err) == 1 && err[0] == nil) {
		return
	}

	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		file = ""
		line = 0
	}
	funcName := runtime.FuncForPC(pc).Name()

	var errPostscriptWithErrs string
	if len(errPostscript) != 0 {
		errPostscriptWithErrs = strings.TrimRight(fmt.Sprintln(append([]any{errPostscript}, err...)...), "\n")
	} else {
		errPostscriptWithErrs = strings.TrimRight(fmt.Sprintln(err...), "\n")
	}
	if stack != nil && len(stack) != 0 {
		log.Printf("[%s] %s %s %s:%d\n%s\n", tag, errPostscriptWithErrs, funcName, file, line, stack)
	} else {
		log.Printf("[%s] %s %s %s:%d\n", tag, errPostscriptWithErrs, funcName, file, line)
	}

	if tag == fatalTag {
		os.Exit(1)
	}
}

func init() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	go func() {
		<-quit
		log.Println("APP STOPPED")
		os.Exit(0)
	}()
}
