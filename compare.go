package main

type CompareObj struct {
	cmd string
	ipc string
	pid int
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