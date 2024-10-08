package conf

import (
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
)

var cfg *ini.File
var afterConfigFn []func()

func InitConfig() {
	var err error

	ex, _ := os.Executable()
	rp := filepath.Dir(ex)
	if cfg, err = ini.Load(rp + "/config.ini"); err != nil {
		FatalExt("config ERROR", err)
	}

	loadConfigSection("Log", &Log)
	loadConfigSection("Redis", &Redis)
	loadConfigSection("MySQL", &MySQL)
	loadConfigSection("Fiber", &Fiber)

	for _, fn := range afterConfigFn {
		fn()
	}
}

// AfterConfigExec Only called by init()
func AfterConfigExec(fn func()) {
	afterConfigFn = append(afterConfigFn, fn)
}

func loadConfigSection(name string, v any) {
	if err := cfg.Section(name).MapTo(v); err != nil {
		FatalExt("config Section: "+name+" ERROR", err)
	}
}

type MySQLConfig struct {
	Addr string
	User string
	Pwd  string
	DB   string
}
type RedisConfig struct {
	Addr   string
	Pwd    string
	DB     int
	Prefix string
	Sep    string
}

// ========= Config Item =========

var Log = struct {
	FilePath         string
	MaxSize          int
	Debug            bool
	StatusTickerTime int
}{}

var Redis = RedisConfig{}

var MySQL = MySQLConfig{}

var Fiber = struct {
	AppName          string
	Addr             string
	RequestLogStdout bool
}{}
