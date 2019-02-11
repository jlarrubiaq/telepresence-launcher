package dockercmd

import (
	"fmt"
	"time"

	"github.com/fsouza/go-dockerclient"

)

// DockerExec attempts to exec against the container image name supplied
func DockerExec(containerImg string, command string) error {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}
	
	// found := false
	// foundID := ""
	for x := 0; x < 10; x++ {
		fmt.Printf("waiting for running container attempt:%v\n", x)
		containers, err := client.ListImages(docker.ListImagesOptions{All: false})
		if err != nil {
			panic(err)
		}

		for _, container := range containers {
			// if container.Image == containerImg && strings.Contains(container.Names[0], "telepresence") {
			// 	found = true
			// 	foundID = container.ID
			// }
			fmt.Println("ID: ", img.ID)
			fmt.Println("RepoTags: ", img.RepoTags)
			fmt.Println("Created: ", img.Created)
			fmt.Println("Size: ", img.Size)
			fmt.Println("VirtualSize: ", img.VirtualSize)
			fmt.Println("ParentId: ", img.ParentID)
		}

		// if found {
		// 	break
		// }
		time.Sleep(3000 * time.Millisecond)
	}

	// fmt.Printf("container found: %s\n", foundID)
	// fmt.Println("Attempting to attach a shell...")
	// config := types.ExecConfig{
	// 	Cmd:          []string{"/bin/bash", "-c", "echo", "hello"},
	// 	Tty:          true,
	// 	AttachStderr: true,
	// 	AttachStdin:  true,
	// 	AttachStdout: true,
	// 	Detach:       false,
	// 	User:         "root",
	// 	Privileged:   false,
	// }
	// idResponse, err := cli.ContainerExecCreate(context.Background(), foundID, config)

	// if err != nil {
	// 	panic(err)
	// }

	// config2 := types.
	// response, err := cli.ContainerExecAttach(context.Background(), idResponse.ID, types.ExecStartCheck{Tty: true, Detach: false})

	// if err != nil {
	// 	panic(err)
	// }

	// defer response.Close()

	return nil
}
