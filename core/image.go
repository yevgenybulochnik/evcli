package core

import (
	"sort"
)

var ImageDict = map[string]string{
	"centos6":    "https://cloud.centos.org/centos/6/images/CentOS-6-x86_64-GenericCloud.qcow2",
	"centos7":    "https://cloud.centos.org/centos/7/images/CentOS-7-x86_64-GenericCloud.qcow2",
	"centos8":    "https://cloud.centos.org/centos/8/x86_64/images/CentOS-8-GenericCloud-8.1.1911-20200113.3.x86_64.qcow2",
	"ubuntu1804": "https://cloud-images.ubuntu.com/bionic/current/bionic-server-cloudimg-amd64.img",
	"ubuntu1810": "https://cloud-images.ubuntu.com/releases/cosmic/release-20190628/ubuntu-18.10-server-cloudimg-amd64.img",
	"ubuntu1904": "https://cloud-images.ubuntu.com/releases/disco/release/ubuntu-19.04-server-cloudimg-amd64.img",
	"ubuntu1910": "https://cloud-images.ubuntu.com/releases/eoan/release/ubuntu-19.10-server-cloudimg-amd64.img",
	"ubuntu2004": "https://cloud-images.ubuntu.com/focal/current/focal-server-cloudimg-amd64.img",
}

func GetAvailableImageNames() string {
	var imgNames []string

	availableImagesText := "Available Images:\n"

	for imgKey := range ImageDict {
		imgNames = append(imgNames, imgKey)
	}

	sort.Strings(imgNames)

	for _, imgName := range imgNames {
		availableImagesText += "\t" + imgName + "\n"
	}

	return availableImagesText
}

func GetImagePath(imageName string) string {
	conn := Connect()

	pools, _ := conn.ListAllStoragePools(0)

	for _, pool := range pools {
		volumes, _ := pool.ListAllStorageVolumes(0)
		for _, volume := range volumes {
			name, _ := volume.GetName()
			if imageName == name {
				imagePath, _ := volume.GetPath()
				return imagePath
			}
		}
	}

	return ""
}
