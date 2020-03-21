package main

import "fmt"

func main()  {

	containers := ContainerList()
	fmt.Print(containers)

	j := 0
	for range containers {
		containerPid := getContainerPid(containers[j])
		fmt.Print(containerPid)
		j++
	}
	/*
	container := Container{"766f780fdeac"}
	getCmd(objectProcess) //procFS()
	objectProcess := Process{4209}
*/
}