package core

import (
	"libvirt.org/libvirt-go"
	"libvirt.org/libvirt-go-xml"
)

func GetPoolInfo(pool *libvirt.StoragePool) (string, string) {
	xmlData, _ := pool.GetXMLDesc(0)
	poolXml := &libvirtxml.StoragePool{}
	poolXml.Unmarshal(xmlData)

	return poolXml.Name, poolXml.Target.Path
}
