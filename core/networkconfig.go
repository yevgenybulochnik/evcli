package core

import (
	"bytes"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Ethernet struct {
	Dhcp4 bool `yaml:"dhcp4"`
}

type NetworkConfig struct {
	Version   int                 `yaml:"version"`
	Ethernets map[string]Ethernet `yaml:"ethernets"`
}

func (networkConfig *NetworkConfig) generateFile(tempDir string) *os.File {
	var data bytes.Buffer

	encoder := yaml.NewEncoder(&data)
	encoder.SetIndent(2)
	encoder.Encode(&networkConfig)

	file, _ := os.Create(filepath.Join(tempDir, "network-config"))
	file.Write(data.Bytes())
	return file
}
