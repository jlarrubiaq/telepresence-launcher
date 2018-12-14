package cmd

import (
	"github.com/aaa-ncnu/telepresence-launcher/pkg/telepresencecmd"
	"github.com/aaa-ncnu/telepresence-launcher/pkg/gitcmd"
	"os"
	"github.com/aaa-ncnu/telepresence-launcher/pkg/k8sClient"
	"fmt"
	"github.com/aaa-ncnu/telepresence-launcher/pkg/prompts"
	"github.com/spf13/cobra" 

)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Bring up the service in the current dir",
	Long:  `Starts an interactive process which will guide you through bringing up a local service using telepresence. Assumes you want to bring up the service in the current working directory.`,
	Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Howdy! Let's get started.")

        client := k8sClient.NewKubeClient()

        namespaces, err := client.ListNamespaces(Config.LabelSelector)
        handleErr(err)

        dir, err := os.Getwd()
        handleErr(err)

        branch, err := gitcmd.GetBranchName(dir)
        handleErr(err)

        _, err = prompts.IsCorrectBranch(branch)
        handleErr(err)

        namespace, err := prompts.NamespaceName(namespaces, Config.Repo)
        handleErr(err)

        deployment, err := prompts.DeploymentName(Config.GetDeploymentPartials())
        handleErr(err)

        k8sdeployment, err := client.GetDeployment(namespace, deployment)
        handleErr(err)

        buildMethods := Config.GetAvailableLaunchMethods(deployment)
        selectedMethod, err := prompts.LaunchMethod(buildMethods)
        handleErr(err)

        methodData, err := Config.GetMethodData(deployment, selectedMethod)
        handleErr(err)

        fmt.Printf("You are on branch %q\n", branch)
        fmt.Printf("You have chosen deployment %q\n", k8sdeployment)
        fmt.Printf("You have chosen namespace %q\n", namespace)
        fmt.Printf("You have chosen launch method %q\n", selectedMethod)

        _, err = prompts.Continue()
        handleErr(err)

        err = methodData.DoPreLaunch()
        handleErr(err)
 
        commandPartial := methodData.GetCommandPartial()

        err = telepresencecmd.RunTelepresence("container", namespace, k8sdeployment, commandPartial)
        handleErr(err)

	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}

