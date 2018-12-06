package cmd

import (
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

        namespaces, err := k8sClient.ListNamespaces()

        if err != nil {
			fmt.Println(err)
			return
		}

        namespace, err := prompts.NamespaceName(namespaces, viper.GetString("repo"))
        deployment, err := prompts.DeploymentName(viper.GetStringMap("deployments"))

		if err != nil {
			fmt.Println(err)
			return
		}

        fmt.Printf("Your username is %q\n", deployment)
        fmt.Printf("Your username is %q\n", namespace)
	},
}

func init() {
	rootCmd.AddCommand(upCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
