package main

import (
	"context"
	"log"

	"github.com/containerd/containerd"
)

func main() {

	client, err := containerd.New("/run/containerd/containerd.sock",
		containerd.WithDefaultNamespace("default"),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// image, err := client.Pull(context.Background(), "docker.io/library/nginx:latest", containerd.WithPullUnpack)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ctx := namespaces.WithNamespace(context.Background(), "default")

	// redis, err := client.NewContainer(ctx, "nginx",
	// 	containerd.WithNewSpec(oci.WithImageConfig(image), oci.WithRootFSPath("/var/lib/containerd")),
	// 	containerd.WithNewSnapshot("snap-nginx", image),
	// )

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// task, err := redis.NewTask(context.Background(), cio.NewCreator(cio.WithStdio))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer task.Delete(context.Background())

	// // the task is now running and has a pid that can be used to setup networking
	// // or other runtime settings outside of containerd
	// pid := task.Pid()

	// log.Println(pid)

	// // start the redis-server process inside the container
	// err = task.Start(context.Background())

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // wait for the task to exit and get the exit status
	// status, err := task.Wait(context.Background())

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(status)

	conts, err := client.Containers(context.Background(), "id==foo")
	if err != nil {
		log.Fatal("containerError: ", err)
	}

	cont := conts[0]

	log.Println(cont.ID())

	t, err := cont.Task(context.Background(), nil)

	if err != nil {
		log.Fatal("taskError: ", err)
	}

	chk, err := t.Checkpoint(context.Background())

	if err != nil {
		log.Fatal("checkpointErr: ", err)
	}

	log.Println(chk)

	err = client.Push(context.Background(), "docker.io/imaskm/criu:v1", chk.Target())
	if err != nil {
		log.Fatal("pushError: ", err)
	}

	// client.TaskService().Checkpoint(context.Background(),
	// 	&tasks.CheckpointTaskRequest{
	// 		ContainerID: cont.ID(),
	// 	},
	// )

	// checkImage, err := cont.Checkpoint(context.Background(), "docker.io/imaskm/criutest:v1",
	// 	containerd.WithCheckpointImage,
	// 	containerd.WithCheckpointRW,
	// 	containerd.WithCheckpointRuntime)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(checkImage.Name())

	// client.Push(context.Background(), )

	// resp, err := client.TaskService().Get(context.Background(), &tasks.GetRequest{
	// 	ContainerID: "foo",
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(resp.Process.ID)

	// opts, err := typeurl.MarshalAny(options.CheckpointOptions{
	// 	ImagePath: "",
	// 	Exit:      false,
	// 	WorkPath:  "",
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// chkResponse, err := client.TaskService().Checkpoint(context.Background(), &tasks.CheckpointTaskRequest{
	// 	ContainerID: resp.Process.ID,
	// 	Options:     opts,
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(chkResponse)

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
