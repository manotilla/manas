package main

type CompareObj struct {
	cmd string `bson:"cmd"`
	ipc string `bson:"ipc"`
	pid int `bson:"pid"`
}

type CompareMap struct {
	compareObj []CompareObj
}

func generateCompareObject(objectProcess Process) CompareObj{

	cmd := getCmd(objectProcess)
	ns := getProcessNS(objectProcess)

	object := CompareObj{cmd: cmd, ipc: ns, pid: objectProcess.pid}

	return object
}