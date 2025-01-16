package xlog

import (
	"PROJECTNAME/conf"
	"PROJECTNAME/utils"
	"context"
	"errors"
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

var _version string
var startTime time.Time
var waitQuitFn []func()

var ShutdownCtx context.Context
var shutdownCancelFunc context.CancelFunc

func init() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	go func() {
		<-quit
		for _, fn := range waitQuitFn {
			fn()
		}
		shutdownCancelFunc()
		log.Println(conf.App.AppName, _version, "STOPPED | RunDuration:", time.Now().Sub(startTime))
		os.Exit(0)
	}()
}

func AppName() string {
	return fmt.Sprint(conf.App.AppName, " ", _version)
}

func Init(version string) {
	_version = version
	log.SetOutput(io.MultiWriter(&lumberjack.Logger{
		Filename:  conf.Log.FilePath,
		MaxSize:   conf.Log.MaxSize,
		LocalTime: true,
	}, os.Stdout))
	startTime = time.Now()
	ShutdownCtx, shutdownCancelFunc = context.WithCancel(context.Background())
	log.Println(conf.App.AppName, _version, "RUNNING | RunPath:", utils.RunPath())

	if conf.App.DebugMode == true {
		Info("\033[41m!!! App.DebugMode = true !!!\033[0m")
	}
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

func NE(errorText ...any) error {
	if errorText == nil || len(errorText) == 0 {
		return nil
	}
	return errors.New(strings.TrimRight(fmt.Sprintln(errorText...), "\n"))
}
