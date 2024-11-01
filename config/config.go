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
	Port                 int64    `yaml:"port"`
	KubeConfigPath       string   `yaml:"kubeConfigPath"`
	Mongo                Mongo    `yaml:"mongo"`
	DisableCronJob       bool     `yaml:"disableCronJob"`
	WhitelistedNamespace []string `yaml:"whitelistedNamespace"`
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

	log.Printf("----reading config from %s----", configYaml)

	configData, err := ioutil.ReadFile(configYaml)
	if err != nil {
		log.Println("----error in reading config file----", err)
		return err
	}

	log.Printf("----finished reading config from %s----", configYaml)

	config = Config{}
	err = yaml.Unmarshal([]byte(configData), &config)
	if err != nil {
		log.Println("----error in Unmarshaling config----", err)
		return err
	}

	return nil
}

func GetConfig() Config {
	return config
}
