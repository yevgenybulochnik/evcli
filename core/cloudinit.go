package core

import (
    "bytes"
    "os"
    "io/ioutil"

    "gopkg.in/yaml.v3"
)

type User struct {
    Name string `yaml:"name"`
}

type CloudConfig struct {
    Hostname string `yaml:"hostname"`
    PreserveHostname bool `yaml:"preserve_hostname"`
    Users []User `yaml:"users"`
}

func (cloudConfig *CloudConfig) generateFile() *os.File {
    var data bytes.Buffer

    encoder := yaml.NewEncoder(&data)
    encoder.SetIndent(2)
    encoder.Encode(&cloudConfig)

    file, _ := ioutil.TempFile("", "user-data")
    file.WriteString("#cloud-config\n")
    file.Write(data.Bytes())
    return file
}
