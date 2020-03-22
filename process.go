package main
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type Process struct {
	pid  int
}

type ProcNamespace struct {
	pid int
	ipc int
}

func getCmd(process Process) string{
	pid := process.pid

	file := fmt.Sprintf("/proc/%d/cmdline", pid)
	output, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Print(err)
	}

	return string(output)
}


func getProcessNS(process Process) string{
	pid := process.pid
	processPath := fmt.Sprintf("/proc/%d/ns", pid)

	symlink := filepath.Join(processPath, "ipc")
	target, err := os.Readlink(symlink)

	if err != nil {
		fmt.Print(err)
	}
	return target
}

func searchFullProc(){
	files, err := ioutil.ReadDir("/proc/")

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		if _, err := strconv.Atoi(f.Name()); err == nil {
			pid := f.Name()

			n, _ := strconv.Atoi(pid)

			objectProcess := Process{n}

			cmd := getCmd(objectProcess)
			ns := getProcessNS(objectProcess)

			object := CompareObj{cmd: cmd, ipc: ns, pid: objectProcess.pid}
			fmt.Print(object)
		}
	}
}
