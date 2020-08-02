package core

import (
	"fmt"
	"libvirt.org/libvirt-go-xml"
)

func CreateVm() {
	conn := Connect()

	domcfg := &libvirtxml.Domain{
		Type: "kvm",
		Name: "testvm",
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
							File: "/home/yevgeny/workspace/evcli/testvm/testvm.img",
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
							File: "/home/yevgeny/workspace/evcli/testvm/testvm.iso",
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
		},
	}

	xml, err := domcfg.Marshal()
	if err != nil {
		panic(err)
	}

	fmt.Println(xml)

	dom, err := conn.DomainDefineXML(xml)
	if err != nil {
		panic(err)
	}

	fmt.Println(dom.GetName())
}
