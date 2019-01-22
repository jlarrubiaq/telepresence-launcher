// +build !darwin

package tplauncher

import (
	"fmt"
	"os"
	"os/exec"
)

// Helper function to add the volume flags to the command.
func (m ContainerMethod) createVolumes() error {
	for _, volume := range m.Volumes {
		cmdslice := []string{"volume", "create", "--driver", "local", "-o", "o=bind", "-o", "type=none", "-o"}
		cmdslice = append(cmdslice, "device="+volume.Src)
		cmdslice = append(cmdslice, volume.Name)

		cmd := exec.Command("docker", escapableEnvVarReplaceSlice(cmdslice)...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Printf("creating volme with command %q\n", cmd.Args)

		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
