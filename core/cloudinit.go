package core

import (
	"bytes"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type User struct {
	Name        string   `yaml:"name"`
	Groups      string   `yaml:"groups,omitempty"`
	Sudo        []string `yaml:"sudo,omitempty"`
	LockPasswd  bool     `yaml:"lock_passwd"`
	Passwd      string   `yaml:"passwd,omitempty"`
	Shell       string   `yaml:"shell"`
	SshAuthKeys []string `yaml:"ssh_authorized_keys,omitempty"`
}

type CloudConfig struct {
	Hostname         string   `yaml:"hostname"`
	PreserveHostname bool     `yaml:"preserve_hostname"`
	Users            []User   `yaml:"users,omitempty"`
	SshAuthKeys      []string `yaml:"ssh_authorized_keys,omitempty"`
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
