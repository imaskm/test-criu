package main

import (
	"context"
	"log"
	"os"

	"github.com/containerd/containerd"
	refdocker "github.com/containerd/containerd/reference/docker"
	"github.com/containerd/nerdctl/pkg/imgutil/commit"
	"github.com/containerd/nerdctl/pkg/imgutil/dockerconfigresolver"
	"github.com/containerd/nerdctl/pkg/imgutil/push"
	"github.com/containerd/nerdctl/pkg/platformutil"
)

func main() {

	client, err := containerd.New("/run/containerd/containerd.sock",
		containerd.WithDefaultNamespace("default"),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// client.SnapshotService("").

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

	conts, err := client.Containers(context.Background(), "id==foo3")
	if err != nil {
		log.Fatal("containerError: ", err)
	}

	cont := conts[0]

	log.Println(cont.ID())

	img := "docker.io/imaskm/testcommit:v3"

	digest, err := commit.Commit(context.Background(), client, cont, &commit.Opts{
		Pause: false,
		Ref:   img,
	})

	if err != nil {
		log.Fatal("digestErrors: ", err)
	}

	log.Println(digest)

	platMC, err := platformutil.NewMatchComparer(true, []string{})
	if err != nil {
		log.Fatal("MatchErr: ", err)
	}

	named, err := refdocker.ParseDockerRef(img)
	if err != nil {
		log.Fatal(err)
	}
	ref := named.String()
	refDomain := refdocker.Domain(named)

	reg := os.Getenv("REGISTRY")
	if reg == "" {
		reg = "docker.io"
	}

	username := os.Getenv("DOCKER_USERNAME")
	password := os.Getenv("DOCKER_PASSWORD")

	// var credFunc dockerconfigresolver.AuthCreds

	credFunc := func(registry string) (string, string, error) {
		return username, password, nil
	}

	if err != nil {
		log.Fatal("authErr: ", err)
	}

	resolver, err := dockerconfigresolver.New(context.Background(), refDomain, dockerconfigresolver.WithAuthCreds(credFunc))
	if err != nil {
		log.Fatal(err)
	}

	err = push.Push(context.Background(), client, resolver, os.Stdout, ref, ref, platMC, false)

	if err != nil {
		log.Fatal("pushError: ", err)
	}
	// t, err := cont.Task(context.Background(), nil)

	// if err != nil {
	// 	log.Fatal("taskError: ", err)
	// }

	// chk, err := t.Checkpoint(context.Background())

	// if err != nil {
	// 	log.Fatal("checkpointErr: ", err)
	// }

	// log.Println(chk)

	// err = client.Push(context.Background(), "docker.io/imaskm/criu:v1", chk.Target())
	// if err != nil {
	// 	log.Fatal("pushError: ", err)
	// }

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
