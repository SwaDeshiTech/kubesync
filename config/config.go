package config

import (
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var config Config

type Config struct {
	Port                  int64 `yaml:"port"`
	Mongo                 Mongo `yaml:"mongo"`
	DisableCronJob        bool  `yaml:"disableCronJob"`
	DisableRESTController bool  `yaml:"disableRESTController"`
}

type Mongo struct {
	URI             string `yaml:"uri"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	MaxPoolSize     uint64 `yaml:"maxPoolSize"`
	MinPoolSize     uint64 `yaml:"minPoolSize"`
	MaxIdleTime     int64  `yaml:"maxIdleTime"`
	MaxConnIdleTime int64  `yaml:"maxConnIdleTime"`
	ConnectTimeout  int64  `yaml:"connectTimeout"`
}

func ReadConfig() error {

	//Default path of the config folder
	configFolder := ""

	flag.StringVar(&configFolder, "configFolder", "conf", "Config folder")
	flag.Parse()

	configYaml := filepath.Join(configFolder, "config.yml")

	configData, err := ioutil.ReadFile(configYaml)
	if err != nil {
		log.Println(err)
		return err
	}

	config = Config{}
	err = yaml.Unmarshal([]byte(configData), &config)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetConfig() Config {
	return config
}
