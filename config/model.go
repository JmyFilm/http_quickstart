package config

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
