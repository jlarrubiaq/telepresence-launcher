package dockercmd

import (
	"os"
	"os/exec"
)

func DockerBuild(command []string) {
	cmd := exec.Command("docker", command...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
