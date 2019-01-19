package main

import (
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/exfly/dockerc/cgroups"
	"github.com/exfly/dockerc/cgroups/subsystems"
	"github.com/exfly/dockerc/container"
)

func Run(tty bool, comArray []string, res *subsystems.ResourceConfig) {
	parent, writePipe := container.NewParentProcess(tty)
	if parent == nil {
		log.Errorf("New parent process error")
		return
	}
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	cgroupMenaget := cgroups.NewCgroupManager("dockerc-cgroup")
	defer cgroupMenaget.Destroy()
	cgroupMenaget.Set(res)
	cgroupMenaget.Apply(parent.Process.Pid)
	sendInitCommand(comArray, writePipe)
	parent.Wait()
	os.Exit(0)
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	log.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}
