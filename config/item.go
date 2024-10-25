package config

func load() {
	loadConfigSection("Log", &Log)
	loadConfigSection("Redis", &Redis)
	loadConfigSection("MySQL", &MySQL)
	loadConfigSection("Fiber", &Fiber)
	loadConfigSection("App", &App)
}

var Log = struct {
	FilePath         string
	DebugInfo        bool
	RequestLogStdout bool
	MaxSize          int
	StatusTickerTime int
}{}

var Redis = RedisConfig{}

var MySQL = MySQLConfig{}

var Fiber = struct {
	ListenOption    string
	HTTPListenAddr  string
	HTTPSListenAddr string
	TLSCertFile     string
	TLSKeyFile      string
	LimitMax        int
	LimitExp        int
}{}

var App = struct {
	AppName string
	NodeId  int64
}{}
