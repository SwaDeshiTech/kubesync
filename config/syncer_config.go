package config

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var syncerConfig SyncerConfig

type SyncerConfig struct {
	Name                 string   `yaml:"name"`
	WhitelistedNamespace []string `yaml:"whitelistedNamespace"`
	ConfigMapList        []string `yaml:"configMapList"`
	SecretList           []string `yaml:"secretList"`
}

func ReadSyncerConfig() error {

	//Default path of the config folder
	configFolder := "conf"

	configYaml := filepath.Join(configFolder, "syncer.yml")

	configData, err := ioutil.ReadFile(configYaml)
	if err != nil {
		log.Println(err)
		return err
	}

	syncerConfig = SyncerConfig{}
	err = yaml.Unmarshal([]byte(configData), &syncerConfig)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetSyncerConfig() SyncerConfig {
	return syncerConfig
}
