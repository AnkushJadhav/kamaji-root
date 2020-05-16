package app

import (
	"errors"
	"io"
	"os"

	"github.com/AnkushJadhav/kamaji-root/logger"

	"gopkg.in/yaml.v2"
)

type serverConfig struct {
	Port    int    `yaml:"port"`
	BindIP  string `yaml:"bindIP"`
	LogFile string `yaml:"logFile,omitempty"`
}

type adminConfig struct {
	RootToken string `yaml:"rootToken"`
}

type dbConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type config struct {
	Server serverConfig `yaml:"server"`
	Admin  adminConfig  `yaml:"admin"`
	Mongo  dbConfig     `yaml:"mongo"`
}

// getConfig loads the config from cfgFile
func getConfig(cfgFile string) (*config, error) {
	logger.Infof("reading config from file : %s", cfgFile)
	f, err := os.Open(cfgFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	conf, err := loadConfig(f)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func loadConfig(r io.Reader) (*config, error) {
	// get yaml decoder for input reader
	decoder := yaml.NewDecoder(r)
	conf := &config{}
	if err := decoder.Decode(conf); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			return nil, errors.New("Invalid format of config file")
		}
		return nil, err
	}

	return conf, nil
}
