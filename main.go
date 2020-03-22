package main

import (
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)
func main()  {


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
	podcastsCollection := quickstartDatabase.Collection("processes")

	containers := ContainerList()
	procList   := SearchFullProc()
	j := 0



	for range containers {
		containerPid := getContainerPid(containers[j])

		objectProcess := Process{containerPid}
		obj := generateCompareObject(objectProcess)
		fmt.Print(obj)

		containerProcessList, err := podcastsCollection.InsertOne(ctx, bson.D{
			{"pid", obj.pid},
			{"ipc", obj.ipc},
			{"cmd", obj.cmd},
		})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(containerProcessList.InsertedID)

		j++
	}

	k := 0

	for range procList {

		objectProcess := Process{procList[k]}
		mainPid := generateCompareObject(objectProcess)
		fmt.Print(mainPid)

		mainPidList, err := podcastsCollection.InsertOne(ctx, bson.D{
			{"pid", mainPid.pid},
			{"ipc", mainPid.ipc},
			{"cmd", mainPid.cmd},
		})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(mainPidList.InsertedID)
		k++
	}


}