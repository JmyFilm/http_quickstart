package config

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"
)

var cfg *ini.File
var initialized = &sync.WaitGroup{}

func init() {
	initialized.Add(1)
}

func InitConfig() {
	var err error
	ex, _ := os.Executable()

	if cfg, err = ini.Load(path.Join(filepath.Dir(ex), "config.ini")); err != nil {
		log.Fatal("config ERROR", err)
	}
	load()
	initialized.Done()
}

func WaitInitExec(fn func()) {
	go func() {
		initialized.Wait()
		fn()
	}()
}

func loadConfigSection(name string, v any) {
	if err := cfg.Section(name).MapTo(v); err != nil {
		log.Fatal("config Section:", name, "ERROR", err)
	}
}
