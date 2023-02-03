package conf

type Config struct {
	OkexConfigs []*OkexConfig `toml:"okex_config"`
}

type OkexConfig struct {
	Name      string `toml:"name"`
	SecretKey string `toml:"secret_key"`
}
