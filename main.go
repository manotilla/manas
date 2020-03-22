package main

import (
	"fmt"
)

type CompareObj struct {
	cmd string
	ipc string
	pid int
}


func main()  {

	containers := ContainerList()

	j := 0

	for range containers {
		containerPid := getContainerPid(containers[j])

		objectProcess := Process{containerPid}
		cmd := getCmd(objectProcess)
		ns := getProcessNS(objectProcess)

		object := CompareObj{cmd: cmd, ipc: ns, pid: objectProcess.pid}

		fmt.Print(object)
		j++

	}

}