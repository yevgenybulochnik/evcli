package core

import (
	libvirt "libvirt.org/libvirt-go"
)

func Connect() *libvirt.Connect {
	conn, err := libvirt.NewConnect("qemu:///system")

	if err != nil {
		panic(err)
	}

	return conn
}
