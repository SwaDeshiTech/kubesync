package config

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var syncerConfig SyncerConfig

type SyncerConfig struct {
	SyncerConfigs []Syncer `yaml:"syncer"`
}

type Syncer struct {
	Name                 string   `yaml:"name"`
	SourceNamespace      string   `yaml:"sourceNamespace"`
	DestinationNamespace []string `yaml:"destinationNamespace"`
	ConfigMapList        []string `yaml:"configMapList"`
	SecretList           []string `yaml:"secretList"`
	K8sClusterName       string   `yaml:"k8sClusterName"`
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
