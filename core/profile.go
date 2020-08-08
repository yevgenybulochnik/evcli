package core

import (
    "fmt"
    "os"
    "bytes"

    "io/ioutil"
    "gopkg.in/yaml.v3"
)


type ProfileConfig struct {
    Profiles map[string]Profile `yaml:"profiles"`
}

type Profile struct {
    Image string `yaml:"image"`
}

func getImagePath(profile *Profile) string {
    return GetImagePath(profile.Image)
}

func ProfileExists(name string) bool {
    file, _:= GetProfilesFile()

    var profileConfig ProfileConfig

    yaml.Unmarshal(file, &profileConfig)

    if _, found := profileConfig.Profiles[name]; found {
        return true
    }

    return false
}

func AddProfile(name string, image string) {

    file, path := GetProfilesFile()

    var profileConfig ProfileConfig

    if len(file) == 0 {
        profileConfig.Profiles = map[string]Profile{}
    }

    yaml.Unmarshal(file, &profileConfig)

    if _, found := profileConfig.Profiles[name]; found {
        fmt.Printf("Profile name %v already exists\n", name)
        os.Exit(0)
    } else {
        profileConfig.Profiles[name] = Profile{
            Image: image,
        }
    }

    var b bytes.Buffer

    enc := yaml.NewEncoder(&b)

    enc.SetIndent(2)
    enc.Encode(&profileConfig)

    ioutil.WriteFile(path, b.Bytes(), 0755)
}
