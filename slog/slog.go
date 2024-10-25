package slog

import (
	"edit-your-project-name/config"
	"edit-your-project-name/utils"
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const fatalTag = "FATAL"

var StartTime time.Time
var waitQuitFn []func()

func init() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	go func() {
		<-quit
		for _, fn := range waitQuitFn {
			fn()
		}
		log.Println("APP STOPPED | RunDuration:", time.Now().Sub(StartTime))
		os.Exit(0)
	}()
}

func InitLog() {
	log.SetOutput(io.MultiWriter(&lumberjack.Logger{
		Filename:  config.Log.FilePath,
		MaxSize:   config.Log.MaxSize,
		LocalTime: true,
	}, os.Stdout))
	StartTime = time.Now()
	log.Println("APP RUNNING | RunPath:", utils.RunPath())
}

func WaitQuitExec(fn func()) {
	waitQuitFn = append(waitQuitFn, fn)
}

func PS(postscript ...any) string {
	if postscript == nil || len(postscript) == 0 {
		return ""
	}
	return strings.TrimRight(fmt.Sprintln(postscript...), "\n")
}
