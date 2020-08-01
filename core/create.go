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
	}

	xml, err := domcfg.Marshal()
	if err != nil {
		panic(err)
	}

	dom, err := conn.DomainDefineXML(xml)
	if err != nil {
		panic(err)
	}

	fmt.Println(dom.GetName())
}
