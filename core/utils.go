package core

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "path/filepath"
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

    if _, err := os.Stat(config_dir); os.IsNotExist(err) {
        accept := AskForConfirmation("Evcli config dir not found would you like to create it?", 3)
        if accept {
            os.Mkdir(config_dir, 0755)
        } else {
            fmt.Println("Exiting")
            os.Exit(0)
        }
    }
}
