package telepresencecmd

import (
	"fmt"
	"os"
	"os/exec"
)

//RunTelepresenceOptions contains the options necessary to run the telepresence command.
type RunTelepresenceOptions struct {
	Method     string
	Namespace  string
	Deployment string
	Expose     string
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

	if options.Expose != "" {
		args = append(args, "--expose")
		args = append(args, options.Expose)
	}

	args = append(args, options.MethodArgs...)

	cmd := exec.Command("telepresence", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("For your reference, this is the command being run behind the scenes...")
	fmt.Println(cmd.Args)

	err := cmd.Run()

	if err != nil {
		return err
	}

	options.TpChan <- true

	return nil

}
