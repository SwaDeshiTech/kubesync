package config

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/SwaDeshiTech/kubesync/enums"
	"gopkg.in/yaml.v2"
)

var config Config

type Config struct {
	Port                 int64          `yaml:"port"`
	KubeConfigPath       string         `yaml:"kubeConfigPath"`
	UseServiceAccount    bool           `yaml:"useServiceAccount"`
	ServiceAccountName   string         `yaml:"serviceAccountName"`
	DisableCronJob       bool           `yaml:"disableCronJob"`
	CronSchedules        []CronSchedule `yaml:"cronSchedules"`
	K8sClusters          []K8sCluster   `yaml:"k8sClusters"`
	Syncers              []Syncer       `yaml:"syncers"`
	WhitelistedNamespace []string       `yaml:"whitelistedNamespace"`
}

type CronSchedule struct {
	UUID           string          `yaml:"uuid"`
	CronExpression string          `yaml:"cronExpression"`
	JobName        string          `yaml:"jobName"`
	JobDescription string          `yaml:"jobDescription"`
	StartDate      string          `yaml:"startDate"`
	EndDate        string          `yaml:"endDate"`
	Priority       string          `yaml:"priority"`
	Frequency      enums.Frequency `yaml:"frequency"`
	JobType        string          `yaml:"jobType"`
	Resources      []Resource      `yaml:"resources"`
	IsActive       bool            `yaml:"isActive"`
}

type Resource struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

type K8sCluster struct {
	Name                 string `yaml:"name"`
	DisplayName          string `yaml:"displayName"`
	EndPoint             string `yaml:"endpoint"`
	ServerID             string `yaml:"serverID"`
	ClientID             string `yaml:"clientID"`
	TenantID             string `yaml:"tenantID"`
	KubeServerIP         string `yaml:"kubeServerIP"`
	CertificateAuthority string `yaml:"certificateAuthority"`
	IsActive             bool   `yaml:"isActive"`
}

type Syncer struct {
	Name                 string   `yaml:"name"`
	SourceNamespace      string   `yaml:"sourceNamespace"`
	DestinationNamespace []string `yaml:"destinationNamespace"`
	ConfigMapList        []string `yaml:"configMapList"`
	SecretList           []string `yaml:"secretList"`
	K8sClusterName       string   `yaml:"k8sClusterName"`
	SkipNamespace        []string `yaml:"skipNamespace"`
}

func ReadConfig() error {

	//Default path of the config folder
	configFolder := ""

	flag.StringVar(&configFolder, "configFolder", "conf", "Config folder")
	flag.Parse()

	if os.Getenv("CONFIG_FOLDER") != "" {
		configFolder = os.Getenv("CONFIG_FOLDER")
	}

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
