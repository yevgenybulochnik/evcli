package core

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func AskForConfirmation(s string, tries int) bool {
	reader := bufio.NewReader(os.Stdin)
	for ; tries > 0; tries-- {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		} else {
			fmt.Println("Invalid entry")
		}
	}
	return false
}

func CheckOrCreateConfigDir() {
	user_home, _ := os.UserHomeDir()
	config_dir := filepath.Join(user_home, ".evcli")
	profiles_file := filepath.Join(config_dir, "profiles.yaml")

	if _, err := os.Stat(config_dir); os.IsNotExist(err) {
		accept := AskForConfirmation("Evcli config dir not found would you like to create it?", 3)
		if accept {
			os.Mkdir(config_dir, 0755)
			os.Create(profiles_file)
		} else {
			fmt.Println("Exiting")
			os.Exit(0)
		}
	}
}

func GetProfilesFile() ([]byte, string) {
	user_home, _ := os.UserHomeDir()
	config_dir := filepath.Join(user_home, ".evcli")
	profiles_file := filepath.Join(config_dir, "profiles.yaml")
	file, _ := ioutil.ReadFile(profiles_file)
	return file, profiles_file
}

func GetGlobalProfiles() (ProfileConfig, string) {
	profilesFile, path := GetProfilesFile()
	var profiles ProfileConfig
	profiles.Parse(profilesFile)
	return profiles, path
}

func GetUserSshPublicKey() string {
	userHome, _ := os.UserHomeDir()
	sshDir := filepath.Join(userHome, ".ssh", "id_rsa.pub")
	pubKey, _ := ioutil.ReadFile(sshDir)
	return string(pubKey)
}
