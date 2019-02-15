package dockercmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/fsouza/go-dockerclient"
)

// DockerExec attempts to exec against the container image name supplied
func DockerExec(containerImg string, command string, notes string) error {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		return err
	}

	found := false
	foundID := ""
	fmt.Println("waiting for running container")
	for x := 0; x < 15; x++ {
		containers, err := client.ListContainers(docker.ListContainersOptions{All: false})
		if err != nil {
			return err
		}

		for _, container := range containers {
			if container.Image == containerImg && strings.Contains(container.Names[0], "telepresence") {
				found = true
				foundID = container.ID
			}
		}

		if found {
			fmt.Println("found")
			break
		}
		time.Sleep(3000 * time.Millisecond)
	}

	fmt.Printf("container found: %s\n", foundID)
	fmt.Println("Attempting to attach a shell...")

	state, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	defer terminal.Restore(int(os.Stdin.Fd()), state)

	exec, err := client.CreateExec(docker.CreateExecOptions{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
		Cmd:          []string{"/bin/sh"},
		Container:    foundID,
		Context:      context.Background(),
	})

	if err != nil {
		return err
	}

	fmt.Println(notes)
	err = client.StartExec(exec.ID, docker.StartExecOptions{
		InputStream:  os.Stdin,
		OutputStream: os.Stdout,
		ErrorStream:  os.Stderr,
		Detach:       false,
		Tty:          true,
		RawTerminal:  true,
		Context:      context.Background(),
	})

	if err != nil {
		return err
	}
	fmt.Println("your shell has terminated. Press ctrl+c to stop telepresence")

	return nil
}
