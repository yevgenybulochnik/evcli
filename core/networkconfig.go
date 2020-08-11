package core

import (
	"bytes"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type Ethernet struct {
	Dhcp4 bool `yaml:"dhcp4"`
}

type NetworkConfig struct {
	Version   int                 `yaml:"version"`
	Ethernets map[string]Ethernet `yaml:"ethernets"`
}

func (networkConfig *NetworkConfig) generateFile() *os.File {
	var data bytes.Buffer

	encoder := yaml.NewEncoder(&data)
	encoder.SetIndent(2)
	encoder.Encode(&networkConfig)

	file, _ := ioutil.TempFile("", "network-config")
	file.Write(data.Bytes())
	return file
}
