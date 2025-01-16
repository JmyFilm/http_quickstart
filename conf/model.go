package conf

type sLog = struct {
	FilePath  string // 日志路径
	MaxSize   int    // 轮转日志文件大小上限 MB 0则不限制
	DebugInfo bool   // 调试日志输出
}

type SMySQL struct {
	Addr     string // 地址
	User     string // 用户名
	Pwd      string // 密码
	DB       string // 数据库名
	MaxConns int    // 最大打开连接数 0则不限制
}
type SRedis struct {
	Addr        string // 地址
	Pwd         string // 密码
	DB          int    // 数据库号
	Prefix      string // 前缀
	Sep         string // 分隔符
	PoolSize    int    // 连接池大小 0则自动设置为核心数的10倍
	DialTimeout int    // 超时时间 0则为5秒
}

type sApp struct {
	DebugMode bool   // 开发调试模式 正式环境需为false
	AppId     int64  // 集群ID
	AppName   string // 应用名
}

type sFiber struct {
	ListenOption    string // HTTP 或 HTTPS
	HTTPListenAddr  string // HTTP监听地址 ListenOption为HTTPS时此地址将重定向为HTTPS协议
	HTTPSListenAddr string // HTTPS监听地址
	TLSCertFile     string // HTTPS证书 Pem地址
	TLSKeyFile      string // HTTPS证书 Key地址
	RequestStdOut   bool   // 请求日志输出
	MaxConns        int    // 最大连接数 1024的倍数 0则为256
	LimitKey        string // 限流根据请求头字段名 阿里ALB|Nginx：X-Forwarded-For 空则为c.IP()
	LimitMax        int    // 限制每 LimitExp 秒可请求 LimitMax 次接口 0则不限流
	LimitExp        int    // 限制每 LimitExp 秒可请求 LimitMax 次接口 0则不限流
}
