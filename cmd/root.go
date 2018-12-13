package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aaa-ncnu/telepresence-launcher/pkg/tplauncher"

	"github.com/spf13/cobra"
)

var cfgFile string

// Config contains the LauncherConfig data
var Config tplauncher.LauncherConfig

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tl",
	Short: "telepresence launcher",
	Long:  `Use this tool to launch local code into a test environment via telepresence.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $PWD/.tl.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	var data []byte
	var err error

	if cfgFile != "" {
		// Use config file from the flag.
		data, err = ioutil.ReadFile(cfgFile)
		handleErr(err)
	} else {
		// Use default config file ($PWD/.tl.yaml)
		dir, err := os.Getwd()
		handleErr(err)
		data, err = ioutil.ReadFile(dir + "/.tl.yaml")
		handleErr(err)
	}

	Config, err = tplauncher.NewConfig(data)
}

func handleErr(err error) {
	if err != nil {
		fmt.Println("ERROR!!!")
		log.Fatalln(err)
	}
}
