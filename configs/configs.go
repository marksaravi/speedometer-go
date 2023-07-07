package configs

import (
	"encoding/json"
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
	json.Unmarshal([]byte(content), &configs)
	return configs
}
