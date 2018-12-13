// +build darwin

package tplauncher

import (
	"fmt"
	"os"
	"os/exec"
)

// Helper function to add the volume flags to the command. This will use NFS volumes if running MacOS
func (m ContainerMethod) createVolumes() error {
	fmt.Println("NOTE: MacOS uses docker native NFS for mounting volumes. Make sure you have NFS server set up on your host machine.")
	for _, volume := range m.Volumes {
		cmdslice := []string{"volume", "create", "--driver", "local", "-o", "type=nfs", "-o", "o=addr=host.docker.internal,rw", "-o"}
		cmdslice = append(cmdslice, "device=:"+volume.Src)
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

// docker volume create --driver local -o type=nfs -o device=:${PWD} -o o=addr=host.docker.internal,rw lpfs
