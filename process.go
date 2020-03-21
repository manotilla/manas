package main
import (
	"fmt"
	"io/ioutil"
)

type Process struct {
	pid  int32
	path string
}

func getCmd(process Process){
	pid := process.pid
	path := process.path

	file := fmt.Sprintf("%s/%d/cmdline", path, pid)
	output, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(output))
}

func getNamespace(process Process){


}