package core

import (
	"bytes"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type MetaData struct {
	InstanceId    string `yaml:"instance-id"`
	LocalHostname string `yaml:"local-hostname"`
}

func (metaData *MetaData) generateFile(tempDir string) *os.File {
	var data bytes.Buffer

	encoder := yaml.NewEncoder(&data)
	encoder.SetIndent(2)
	encoder.Encode(&metaData)

	file, _ := os.Create(filepath.Join(tempDir, "meta-data"))
	file.Write(data.Bytes())
	return file
}
