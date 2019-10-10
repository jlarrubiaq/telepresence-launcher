// +build darwin

package tplauncher

import (
	"fmt"
	"os"
	"os/exec"
)

// Helper function to add the volume flags to the command. This will use NFS volumes if running MacOS
func (m ContainerMethod) createVolumes(useBindMounts bool) error {

	var cmdslice []string
	for _, volume := range m.Volumes {
		if useBindMounts {
			cmdslice = []string{"volume", "create", "--driver", "local", "-o", "o=bind", "-o", "type=none", "-o"}
			cmdslice = append(cmdslice, "device="+volume.Src)
			fmt.Println("NOTE: Using Bind mounts due to usebindmount flag being true.")
		} else {
			cmdslice = []string{"volume", "create", "--driver", "local", "-o", "type=nfs", "-o", "o=addr=host.docker.internal,rw", "-o"}
			cmdslice = append(cmdslice, "device=:"+volume.Src)
			fmt.Println("NOTE: MacOS uses docker native NFS for mounting volumes. Make sure you have NFS server set up on your host machine.")
		}

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
