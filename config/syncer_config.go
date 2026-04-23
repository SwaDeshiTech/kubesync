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
	Name           string    `yaml:"name"`
	ConfigMap      ConfigMap `yaml:"configMap"`
	Secret         Secret    `yaml:"secret"`
	K8sClusterName string    `yaml:"k8sClusterName"`
}

type ConfigMap struct {
	List                 []string `yaml:"list"`
	SourceNamespace      string   `yaml:"sourceNamespace"`
	DestinationNamespace []string `yaml:"destinationNamespace"`
}

type Secret struct {
	List                 []string `yaml:"list"`
	SourceNamespace      string   `yaml:"sourceNamespace"`
	DestinationNamespace []string `yaml:"destinationNamespace"`
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
