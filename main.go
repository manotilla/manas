package main

func main()  {

	container := Container{"766f780fdeac"}
	container_pid := getContainerPid(container)

	objectProcess := Process{1915, "/proc"}
	getCmd(objectProcess)
}
