package main

import (
	"os/exec"

	log "github.com/Sirupsen/logrus"
)

func commitContainer(imageName string) {
	mntURL := "/dockerc/mnt"
	imageTar := "/dockerc/" + imageName + ".tar"
	log.Infof("%s", imageTar)
	if _, err := exec.Command("tar", "-czf", imageTar, "-C", mntURL, ".").CombinedOutput(); err != nil {
		log.Errorf("Tar folder %s error %v", mntURL, err)
	}
}
