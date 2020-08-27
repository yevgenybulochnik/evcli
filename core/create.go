package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	// "os/user"

	"libvirt.org/libvirt-go-xml"
)

func CreateImage(vmName string, backingFilePath string, diskSize int, poolPath string) {

	tempDir, _ := ioutil.TempDir("", "evcli-build")

	metaData := MetaData{
		InstanceId:    vmName,
		LocalHostname: vmName,
	}
	metaDataFile := metaData.generateFile(tempDir)

	// user, _ := user.Current()

	cloudConfig := CloudConfig{
		Hostname:         vmName,
		PreserveHostname: false,
		SshAuthKeys: []string{
			GetUserSshPublicKey(),
		},
		// Users: []User{
		//     {
		//         Name:   user.Username,
		//         Shell:  "/bin/bash",
		//         Groups: "sudo",
		//         Sudo: []string{
		//             "All=(ALL) NOPASSWD:ALL",
		//         },
		//         SshAuthKeys: []string{
		//             GetUserSshPublicKey(),
		//         },
		//         LockPasswd: false,
		//         Passwd:     "$6$1V4l1w/9D8/.Jv$vfCnT6fAQ2yfJ5GBfGVF4AsRMmdDzv2L/catZpFLLoqlIPr2DsOr.uNG7lqSxlWUPfmNHliD9t0A3f5i.etn60",
		//     },
		// },
	}

	cloudConfigFile := cloudConfig.generateFile(tempDir)

	networkConfig := NetworkConfig{
		Version: 2,
		Ethernets: map[string]Ethernet{
			"enp0s2": {Dhcp4: true},
		},
	}

	networkConfigFile := networkConfig.generateFile(tempDir)

	fmt.Println(networkConfigFile.Name())

	qemuImgPath, _ := exec.LookPath("qemu-img")

	qemuImgCmd := &exec.Cmd{
		Path: qemuImgPath,
		Args: []string{
			qemuImgPath,
			"create",
			"-f",
			"qcow2",
			"-o",
			fmt.Sprintf("backing_file=%v", backingFilePath),
			fmt.Sprintf("%s/%s.img", poolPath, vmName),
			fmt.Sprintf("%vG", diskSize),
		},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	genIsoImgPath, _ := exec.LookPath("genisoimage")

	genIsoImgCmd := &exec.Cmd{
		Path: genIsoImgPath,
		Args: []string{
			genIsoImgPath,
			"-output",
			fmt.Sprintf("%s/%s.iso", poolPath, vmName),
			"-volid",
			"cidata",
			"-joliet",
			"-rock",
			cloudConfigFile.Name(),
			metaDataFile.Name(),
			networkConfigFile.Name(),
		},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	fmt.Println(qemuImgCmd.String())
	fmt.Println(genIsoImgCmd.String())
	qemuImgCmd.Run()
	genIsoImgCmd.Run()
    os.RemoveAll(tempDir)
}

func CreateVm(vmName string, poolPath string) {
	conn := Connect()

	domcfg := &libvirtxml.Domain{
		Type: "kvm",
		Name: vmName,
		Memory: &libvirtxml.DomainMemory{
			Value: 4096,
			Unit:  "MB",
		},
		VCPU: &libvirtxml.DomainVCPU{
			Value: 1,
		},
		OS: &libvirtxml.DomainOS{
			Type: &libvirtxml.DomainOSType{
				Arch: "x86_64",
				Type: "hvm",
			},
		},
		Devices: &libvirtxml.DomainDeviceList{
			Disks: []libvirtxml.DomainDisk{
				{
					Device: "disk",
					Driver: &libvirtxml.DomainDiskDriver{
						Name: "qemu",
						Type: "qcow2",
					},
					Source: &libvirtxml.DomainDiskSource{
						File: &libvirtxml.DomainDiskSourceFile{
							File: fmt.Sprintf("%s/%s.img", poolPath, vmName),
						},
					},
					Target: &libvirtxml.DomainDiskTarget{
						Dev: "vda",
						Bus: "virtio",
					},
				},
				{
					Device: "cdrom",
					Driver: &libvirtxml.DomainDiskDriver{
						Name: "qemu",
					},
					Source: &libvirtxml.DomainDiskSource{
						File: &libvirtxml.DomainDiskSourceFile{
							File: fmt.Sprintf("%s/%s.iso", poolPath, vmName),
						},
					},
					Target: &libvirtxml.DomainDiskTarget{
						Dev: "hdd",
						Bus: "ide",
					},
				},
			},
			Serials: []libvirtxml.DomainSerial{
				{
					Target: &libvirtxml.DomainSerialTarget{},
				},
			},
			Consoles: []libvirtxml.DomainConsole{
				{
					Target: &libvirtxml.DomainConsoleTarget{
						Type: "serial",
					},
				},
			},
			Interfaces: []libvirtxml.DomainInterface{
				{
					Source: &libvirtxml.DomainInterfaceSource{
						Network: &libvirtxml.DomainInterfaceSourceNetwork{
							Network: "default",
						},
					},
					Model: &libvirtxml.DomainInterfaceModel{
						Type: "virtio",
					},
				},
			},
		},
	}

	xml, err := domcfg.Marshal()
	if err != nil {
		panic(err)
	}

	dom, err := conn.DomainDefineXML(xml)
	if err != nil {
		panic(err)
	}

	dom.Create()

	fmt.Println(dom.GetXMLDesc(0))

}
