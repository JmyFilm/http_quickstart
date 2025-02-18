package conf

import (
	"PROJECTNAME/utils"
	"gopkg.in/ini.v1"
	"io/fs"
	"log"
	"path/filepath"
	"sync"
	"time"
)

var cfg *ini.File
var initialized = &sync.WaitGroup{}

func init() {
	initialized.Add(1)
}

func Init(path string) {
	var err error

	if path == "" {
		path = findLatestIniFile()
	}
	if cfg, err = ini.Load(path); err != nil {
		log.Fatal("config Load ERROR ", err)
	}

	load()
	if err := verify(); err != nil {
		log.Fatal("config Verify ERROR ", err)
	}

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
		log.Fatal("config Section: ", name, " ERROR ", err)
	}
}

func findLatestIniFile() string {
	var latestFile string
	var latestModTime time.Time
	dir := utils.RunPath()

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || filepath.Ext(d.Name()) != ".ini" {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		if info.ModTime().After(latestModTime) {
			latestFile = path
			latestModTime = info.ModTime()
		}
		return nil
	})
	if err != nil {
		log.Fatal("config ERROR ", err)
	}

	if latestFile == "" {
		latestFile = filepath.Join(dir, "config.ini")
	}
	return latestFile
}
