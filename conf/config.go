package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/bad-superman/test/logging"
)

type Config struct {
	InfluxConfig InfluxConfig   `toml:"influx_config"`
	OkexConfigs  []*OkexConfig  `toml:"okex_config"`
	Database     DatabaseConfig `toml:"database"`
	DTalkToken   string         `toml:"dingtalk_token"`
	Thegraph     ThegraphConfig `toml:"thegraph"`
}

type DatabaseConfig struct {
	DSN string `toml:"dsn"`
}

type OkexConfig struct {
	Name       string `toml:"name"`
	ApiKey     string `toml:"api_key"`
	SecretKey  string `toml:"secret_key"`
	Passphrase string `toml:"passphrase"`
}

type InfluxConfig struct {
	Token  string `toml:"token"`
	URL    string `toml:"url"`
	Org    string `toml:"org"`
	Bucket string `toml:"bucket"`
}

type ThegraphConfig struct {
	ApiKey string `toml:"api_key"`
}

var _config *Config = new(Config)

func Init(path string) *Config {
	_, err := toml.DecodeFile(path, _config)
	if err != nil {
		logging.Panicf("init config error,err:%v", err)
	}
	return _config
}

func GetConfig() *Config {
	return _config
}
