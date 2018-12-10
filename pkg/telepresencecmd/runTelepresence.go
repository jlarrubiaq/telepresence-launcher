package telepresencecmd

import (
	"fmt"
	"os"
	"os/exec"
)

func RunTelepresence(method string, volumes [][]string, namespace string, deployment string, image string, commands []string) {
	args := []string{}
	args = append(args, "--swap-deployment")
	args = append(args, deployment)
	args = append(args, "--namespace")
	args = append(args, namespace)
	args = append(args, "--"+method)
	args = append(args, "--rm")

	for _, volume := range volumes {
		args = append(args, "-v")
		args = append(args, escapableEnvVarReplace(volume[0])+":"+volume[1])
	}

	args = append(args, image)

	if len(commands) > 0 {
		args = append(args, commands...)
	}

	cmd := exec.Command("telepresence", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println(cmd.Args)
	cmd.Run()
}

//escapableEnvVarReplace wraps os.Getenv to allow for escaping with $$.
func escapableEnvVarReplace(s string) string {
	return os.Expand(s, func(s string) string {
		if s == "$" {
			return "$"
		}
		realEnvVal := os.Getenv(s)

		return realEnvVal
	})
}

// telepresence --swap-deployment rd8200b9595-mwg-landing-page-frontend-service --namespace pr85-aaa-ncnu-landing-page-frontend-service-mwg --docker-run --rm -it -v $(pwd):/usr/src/app aaadigital/landing-page-frontend-service:dev npm run debug
