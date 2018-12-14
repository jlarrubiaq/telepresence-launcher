package telepresencecmd

import (
	"fmt"
	"os"
	"os/exec"
)

// RunTelepresence runs the telepresence command specified in the config.
func RunTelepresence(method string, namespace string, deployment string, methodArgs []string) error {
	args := []string{}
	args = append(args, "--logfile", "/tmp/telepresence.log")
	args = append(args, "--swap-deployment")
	args = append(args, deployment)
	args = append(args, "--namespace")
	args = append(args, namespace)
	args = append(args, methodArgs...)

	cmd := exec.Command("telepresence", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println(cmd.Args)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
