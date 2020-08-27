package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/yevgenybulochnik/evcli/core"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Ssh into vm",
	Long:  "Ssh into virtual machine given virtual machine name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ssh(args[0])
	},
}

func ssh(vmName string) {
	conn := core.Connect()

	domain, err := conn.LookupDomainByName(vmName)
	if err != nil {
		fmt.Println("An error has occured, are you sure this vm is running?")
		os.Exit(1)
	}
	interfaces, _ := domain.ListAllInterfaceAddresses(0)

	var ip string

	if len(interfaces) > 0 {
		ip = interfaces[0].Addrs[0].Addr
	} else {
		os.Exit(1)
	}

	sshPath, _ := exec.LookPath("ssh")

	// Future refactor allow a default username to be configured
	args := []string{sshPath, fmt.Sprintf("ubuntu@%s", ip)}
	syscall.Exec(sshPath, args, os.Environ())
}

func init() {
	rootCmd.AddCommand(sshCmd)
}
