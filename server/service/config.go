package service

import "github.com/BurntSushi/toml"

// Config server config
type Config struct {
	Mode    string `toml:"mode"`
	Servers map[string]baseConfig
}

type baseConfig struct {
	Addr       string `toml:"addr"`
	ResolveNet string `toml:"resolveNet"`
	ListenNet  string `toml:"listenNet"`
	Port       string `toml:"port"`
	Topic      string `toml:"topic"`
}

// GetConfig get server config
func GetConfig(tomlFile string) (conf Config, err error) {
	if _, err := toml.DecodeFile(tomlFile, &conf); err != nil {
		return conf, err
	}
	return conf, nil
}
