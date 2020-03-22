package main

import (
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
	"time"
)
func main()  {
	var containerPids []string

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database("manas")
	processCollection := quickstartDatabase.Collection("processes")

	containers := ContainerList()
	procList   := SearchFullProc()

	container_terator := 0



	for range containers {
		containerPid := getContainerPid(containers[container_terator])

		objectProcess := Process{containerPid}
		obj := generateCompareObject(objectProcess)

		containerPids = append(containerPids, obj.ipc)

		container_terator++
	}

	process_iterator := 0

	for range procList {

		objectProcess := Process{procList[process_iterator]}
		mainPid := generateCompareObject(objectProcess)
		fmt.Print(mainPid)


		if strings.Contains(strings.Join(containerPids, ","), mainPid.ipc) {
			mainPidList, err := processCollection.InsertOne(ctx, bson.D{
				{"pid", mainPid.pid},
				{"ipc", mainPid.ipc},
				{"cmd", mainPid.cmd},
				{"is_container", true},

			})
			fmt.Print(mainPidList.InsertedID)

			if err != nil {
				log.Fatal(err)
			}

		}else{

			fmt.Print("Operating System Process")
		}

		process_iterator++
	}


}