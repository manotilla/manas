package main

import "fmt"

/*
import (
	"fmt"
)*/



func main()  {


	containers := ContainerList()
	procList   := SearchFullProc()
	j := 0

	for range containers {
		containerPid := getContainerPid(containers[j])

		objectProcess := Process{containerPid}
		obj := generateCompareObject(objectProcess)
		fmt.Print(obj)

		j++
	}

	k := 0

	for range procList {
		objectProcess := Process{procList[k]}
		mainPid := generateCompareObject(objectProcess)
		fmt.Print(mainPid)
		k++
	}
}