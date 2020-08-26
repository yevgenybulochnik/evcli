package core

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Profile struct {
	Image string `yaml:"image"`
	Base  string `yaml:"base,omitempty"`
}

func (profile *Profile) GetImagePath() string {
	conn := Connect()

	pools, _ := conn.ListAllStoragePools(0)

	for _, pool := range pools {
		volumes, _ := pool.ListAllStorageVolumes(0)
		for _, volume := range volumes {
			name, _ := volume.GetName()
			if profile.Image == name {
				imagePath, _ := volume.GetPath()
				return imagePath
			}
		}
	}

	return ""
}

func (profile *Profile) CreateVM(vmName string, diskSize int, pool string) {
	user_home, _ := os.UserHomeDir()
	poolDir := filepath.Join(user_home, pool)
	CreateImage(vmName, profile.GetImagePath(), diskSize, poolDir)
	CreateVm(vmName, poolDir)
}

type ProfileConfig struct {
	List map[string]Profile `yaml:"profiles"`
}

func (profileConfig *ProfileConfig) Parse(data []byte) error {
	return yaml.Unmarshal(data, profileConfig)
}

func (profileConfig *ProfileConfig) Save(filePath string) {
	var b bytes.Buffer

	enc := yaml.NewEncoder(&b)

	enc.SetIndent(2)
	enc.Encode(&profileConfig)

	ioutil.WriteFile(filePath, b.Bytes(), 0755)
}

func (profileConfig *ProfileConfig) ProfileExists(name string) bool {
	if _, found := profileConfig.List[name]; found {
		return true
	}

	return false
}

func (profileConfig *ProfileConfig) GetProfile(name string) Profile {
	return profileConfig.List[name]
}

func (profileConfig *ProfileConfig) AddProfile(profileName string, profile Profile) {
	if len(profileConfig.List) == 0 {
		profileConfig.List = map[string]Profile{}
	}
	if _, found := profileConfig.List[profileName]; found {
		fmt.Printf("Profile name %v already exists\n", profileName)
		os.Exit(0)
	} else {
		profileConfig.List[profileName] = profile
	}
}
