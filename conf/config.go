package conf

import (
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
	"reflect"
)

var cfg *ini.File
var afterConfigFn []func()

func InitConfig() {
	var err error

	ex, _ := os.Executable()
	rp := filepath.Dir(ex)
	if cfg, err = ini.Load(rp + "/config.ini"); err != nil {
		Fatal("config ERROR:", err)
	}

	loadConfigSection(&Log)
	loadConfigSection(&Redis)
	loadConfigSection(&MySQL)
	loadConfigSection(&Fiber)

	for _, fn := range afterConfigFn {
		fn()
	}
}

// AfterConfigExec Only called by init()
func AfterConfigExec(fn func()) {
	afterConfigFn = append(afterConfigFn, fn)
}

func loadConfigSection(v any) {
	var name string

	if _type := reflect.TypeOf(v); _type.Kind() == reflect.Ptr {
		name = _type.Elem().String()
	} else {
		name = _type.Name()
	}

	if err := cfg.Section(name).MapTo(v); err != nil {
		Fatal("config Section: "+name+" ERROR:", err)
	}
}

var Log = struct {
	FileName         string
	MaxSize          int
	Debug            bool
	StatusTickerTime int
}{}

var Redis = struct {
	Addr   string
	Pwd    string
	DB     int
	Prefix string
	Sep    string
}{}

var MySQL = struct {
	User string
	Pwd  string
	Host string
	Port string
	DB   string
}{}

var Fiber = struct {
	Addr string
}{}
