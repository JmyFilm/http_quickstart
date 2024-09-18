package conf

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"syscall"
)

func InitLog() {
	log.SetOutput(io.MultiWriter(&lumberjack.Logger{
		Filename:  Log.FileName,
		MaxSize:   Log.MaxSize,
		LocalTime: true,
	}, os.Stdout))
	log.Println("APP RUNNING")
}

func Debug(info ...any) {
	if !Log.Debug {
		return
	}
	log.Println("[DEBUG]", info)
}

func Info(info ...any) {
	log.Println("[INFO]", info)
}

func Err(info ...any) {
	doLog("ERROR", nil, info...)
}

func ErrWithStack(info ...any) {
	doLog("ERROR", debug.Stack(), info...)
}

func Fatal(info ...any) {
	if info == nil || len(info) == 0 || (len(info) == 1 && info[0] == nil) {
		return
	}
	doLog("FATAL", nil, info...)
	os.Exit(1)
}

func doLog(tag string, stack []byte, err ...any) {
	if err == nil || len(err) == 0 || (len(err) == 1 && err[0] == nil) {
		return
	}

	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "Unknown"
		line = 0
	}
	funcName := runtime.FuncForPC(pc).Name()

	if stack != nil && len(stack) != 0 {
		log.Printf("[%s] %v %s %s:%d\n%s\n", tag, err, funcName, file, line, stack)
	} else {
		log.Printf("[%s] %v %s %s:%d\n", tag, err, funcName, file, line)
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
