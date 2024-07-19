package config

var Config *Conf

type Conf struct {
	Mysql MysqlConfig
}

type MysqlConfig struct {
	Dsn string `toml:"dsn"`
}
