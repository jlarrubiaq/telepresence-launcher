package cmd

import (
	"github.com/aaa-ncnu/telepresence-launcher/pkg/telepresencecmd"
	"github.com/aaa-ncnu/telepresence-launcher/pkg/dockercmd"
	"github.com/aaa-ncnu/telepresence-launcher/pkg/gitcmd"
	"os"
	"github.com/aaa-ncnu/telepresence-launcher/pkg/k8sClient"
	"fmt"
	"github.com/spf13/viper"
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

        namespaces, err := client.ListNamespaces()

        dir, err := os.Getwd()
        if err != nil {
            fmt.Println(err)
            return
        }

        branch, err := gitcmd.GetBranchName(dir)

        if err != nil {
			fmt.Println(err)
			return
		}

        isCorrectBranch, err := prompts.IsCorrectBranch(branch)

        if !isCorrectBranch || err != nil {
            fmt.Println("Stopping. Please checkout the correct branch and try again.")
            return
        }

        namespace, err := prompts.NamespaceName(namespaces, viper.GetString("repo"))
        deployment, err := prompts.DeploymentName(viper.GetStringMap("deployments"))

        if err != nil {
			fmt.Println(err)
			return
		}

        k8sdeployment, err := client.GetDeployment(namespace, deployment)

        if err != nil {
			fmt.Println(err)
			return
        }
        
        if buildMethod := viper.GetString("deployments."+deployment+".method"); buildMethod == "docker-run" {
            dockercmd.DockerBuild(viper.GetStringSlice("deployments."+deployment+".build"))
        }

        fmt.Printf("You are on branch %s\n", branch)
        fmt.Printf("You have chosen deployment %q\n", k8sdeployment)
        fmt.Printf("You have chosen namespace %q\n", namespace)

        volumes := [][]string{}
        viper.UnmarshalKey("deployments."+deployment+".volumes", &volumes)

        telepresencecmd.RunTelepresence(
            viper.GetString("deployments."+deployment+".method"), 
            volumes,
            namespace,
            k8sdeployment,
            viper.GetString("deployments."+deployment+".image"),
            viper.GetStringSlice("deployments."+deployment+".commands"),
        )

	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
