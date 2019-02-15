package telepresencecmd

import (
	"os"
	"os/exec"
)

//RunTelepresenceOptions contains the options necessary to run the telepresence command.
type RunTelepresenceOptions struct {
	Method     string
	Namespace  string
	Deployment string
	MethodArgs []string
	TpChan     chan bool
}

// RunTelepresence runs the telepresence command specified in the config.
func RunTelepresence(options RunTelepresenceOptions) error {
	args := []string{}
	args = append(args, "--logfile", "/tmp/telepresence.log")
	args = append(args, "--swap-deployment")
	args = append(args, options.Deployment)
	args = append(args, "--namespace")
	args = append(args, options.Namespace)
	args = append(args, options.MethodArgs...)

	cmd := exec.Command("telepresence", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		return err
	}

	options.TpChan <- true

	return nil

}
