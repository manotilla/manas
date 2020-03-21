package main
import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Process struct {
	pid  int32
}

type procNamespace struct {
	pid int
	ipc int
}

func getCmd(process Process){
	pid := process.pid

	file := fmt.Sprintf("/proc/%d/cmdline", pid)
	output, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(output))
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
