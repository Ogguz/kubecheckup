package model

type Config struct {
	Kubernetes struct {
		ConfigFile string `yaml:"configFile"`
	} `yaml:"kubernetes"`

}