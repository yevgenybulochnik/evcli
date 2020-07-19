package core

import (
    "fmt"

    libvirt "libvirt.org/libvirt-go"
)

func Connect() *libvirt.Connect {
    conn, err := libvirt.NewConnect("qemu:///system")

    if err != nil {
        fmt.Println(err)
    }

    return conn
}
