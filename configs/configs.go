package configs

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)
type Configs struct {
	DistPerPulse float64 `yaml:"distance-per-pulse"`
}

func ReadConfigs(path string) Configs {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var configs Configs
	yaml.Unmarshal([]byte(content), &configs)
	return configs
}
