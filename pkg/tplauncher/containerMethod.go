package tplauncher

import (
	"os"
	"os/exec"

	"github.com/aaa-ncnu/telepresence-launcher/pkg/dockercmd"
)

// ContainerMethod is a LaunchMethod. Describes the necessary data to launch a container with telepresence.
type ContainerMethod struct {
	Method     string      `json:"method"`
	BuildSteps []BuildStep `json:"buildSteps"`
	Volumes    []Volume    `json:"volumes"`
	Mounts     []Mount     `json:"mounts"`
	Envs       []string    `json:"env"`
	Image      string      `json:"image"`
	Commands   []string    `json:"commands"`
}

// BuildStep describes a step that should run before launching this container.
type BuildStep struct {
	Cmd  string   `json:"cmd"`
	Args []string `json:"args"`
}

// Volume describes a volume that should be created before launching this container.
type Volume struct {
	Name string `json:"name"`
	Src  string `json:"src"`
}

// Mount describes a mount that should be declared in the for the container to operate correctly.
type Mount struct {
	Type        string `json:"type"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

// DoPreLaunch is the logic for tasks that should run before a container is launched.
func (m ContainerMethod) DoPreLaunch() error {

	for _, buildStep := range m.BuildSteps {
		cmd := exec.Command(buildStep.Cmd, escapableEnvVarReplaceSlice(buildStep.Args)...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	err := m.createVolumes()
	if err != nil {
		return err
	}

	return nil
}

// GetCommandPartial returns the part of the telepresence command specific to the "container" method.
func (m ContainerMethod) GetCommandPartial() []string {
	args := []string{}

	args = append(args, "--method")
	args = append(args, "container")

	args = append(args, "--docker-run")
	args = append(args, "--rm", "--init")

	for _, mount := range m.Mounts {
		args = append(args, "--mount")
		args = append(args, escapableEnvVarReplace("type="+mount.Type+",source="+mount.Source+",destination="+mount.Destination))
	}

	for _, env := range m.Envs {
		args = append(args, "-e")
		args = append(args, env)
	}

	args = append(args, m.Image)

	if len(m.Commands) > 0 {
		args = append(args, m.Commands...)
	}

	return args
}

func (m ContainerMethod) DoPostLaunch() error {
	dockercmd.DockerExec(m.Image, "/bin/sh")
	return nil
}
