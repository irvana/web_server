package configs

type Config struct {
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
	Level   string `mapstructure:"level"`
}
