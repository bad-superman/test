package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/bad-superman/test/logging"
)

type Config struct {
	OkexConfigs []*OkexConfig `toml:"okex_config"`

	DTalkToken string `toml:"dingtalk_token"`
}

type OkexConfig struct {
	Name       string `toml:"name"`
	ApiKey     string `toml:"api_key"`
	SecretKey  string `toml:"secret_key"`
	Passphrase string `toml:"passphrase"`
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
