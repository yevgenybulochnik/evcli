package core

import (
	"gopkg.in/yaml.v3"
)

type Deployment struct {
	Hosts map[string]Profile `yaml:"hosts"`
}

func (deployment *Deployment) Parse(data []byte) error {
	return yaml.Unmarshal(data, deployment)
}
