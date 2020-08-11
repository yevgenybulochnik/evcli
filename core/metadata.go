package core

import (
    "bytes"
    "io/ioutil"
    "os"

    "gopkg.in/yaml.v3"
)

type MetaData struct {
    InstanceId string `yaml:"instance-id"`
    LocalHostname string `yaml:"local-hostname"`
}

func (metaData *MetaData) generateFile() *os.File {
    var data bytes.Buffer

    encoder := yaml.NewEncoder(&data)
    encoder.SetIndent(2)
    encoder.Encode(&metaData)

    file, _ := ioutil.TempFile("", "meta-data")
    file.Write(data.Bytes())
    return file
}
