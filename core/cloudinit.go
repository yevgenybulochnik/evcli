package core

import (
	"bytes"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type User struct {
	Name       string `yaml:"name"`
	Groups     string `yaml:"groups,omitempty"`
	LockPasswd bool   `yaml:"lock_passwd"`
	Passwd     string `yaml:"passwd,omitempty"`
	Shell      string `yaml:"shell"`
}

type CloudConfig struct {
	Hostname         string `yaml:"hostname"`
	PreserveHostname bool   `yaml:"preserve_hostname"`
	Users            []User `yaml:"users"`
}

func (cloudConfig *CloudConfig) generateFile(tempDir string) *os.File {
	var data bytes.Buffer

	encoder := yaml.NewEncoder(&data)
	encoder.SetIndent(2)
	encoder.Encode(&cloudConfig)

	file, _ := os.Create(filepath.Join(tempDir, "user-data"))
	file.WriteString("#cloud-config\n")
	file.Write(data.Bytes())
	return file
}
