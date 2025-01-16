package conf

func load() {
	loadConfigSection("Log", &Log)
	loadConfigSection("MySQL", &MySQL)
	loadConfigSection("Redis", &Redis)
	loadConfigSection("App", &App)
	loadConfigSection("Fiber", &Fiber)
}

var Log = sLog{}
var MySQL = SMySQL{}
var Redis = SRedis{}
var App = sApp{}
var Fiber = sFiber{}
