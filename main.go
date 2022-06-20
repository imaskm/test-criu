package main

import (
	"context"
	"log"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/api/services/tasks/v1"
)

func main() {

	client, err := containerd.New("/run/containerd/containerd.sock",
		containerd.WithDefaultNamespace("default"),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	conts, _ := client.Containers(context.Background())

	log.Println(conts)

	resp, err := client.TaskService().Get(context.Background(), &tasks.GetRequest{
		ContainerID: "foo",
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Process.ID)

	chkResponse, err := client.TaskService().Checkpoint(context.Background(), &tasks.CheckpointTaskRequest{
		ContainerID: resp.Process.ID,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(chkResponse)
	// if resp.Process. {
	// 	log.Println("0 tasks")
	// 	return
	// }

	// for _, t := range resp.Tasks {
	// 	log.Println(t.ID)

	// }

	// tasksvc := client.TaskService()

	// myTask, err := tasksvc.Get(context.Background(), &tasks.GetRequest{
	// 	ContainerID: "testc",
	// 	ExecID:      "",
	// })

}
