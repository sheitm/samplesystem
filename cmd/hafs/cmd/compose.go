/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/HafslundEcoVannkraft/samplesystem/internal/compose"
	"github.com/HafslundEcoVannkraft/samplesystem/internal/model"

	"github.com/spf13/cobra"
)

// composeCmd represents the compose command
var composeCmd = &cobra.Command{
	Use:   "compose",
	Short: "Provides methods for interacting with Docker Compose.",
	Long: `Provides methods for interacting with Docker Compose. Specifically,
assembling system.yaml and underlying app.yaml files into a docker-compose.yaml file.

This file can either be written to disk ("dump") or used to start a Docker Compose stack.

To write to disk 'hafs compose dump'
To start a stack 'hafs compose up'
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("compose called")
	},
}

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dumps the Docker compose file to disk (or stdout).",
	Long: `Dumps the system to a file. The system is assembled from the system.yaml file
and the app.yaml files in the apps directory.`,

	Run: func(cmd *cobra.Command, args []string) {
		f := cmd.Flag("file").Value.String()
		if f == "" {
			fmt.Println("No file specified")
			return
		}

		sys, err := model.AssembleSystem(model.WithSystemFile(f))
		if err != nil {
			fmt.Println(err)
			return
		}

		b, err := compose.Compose(sys)
		if err != nil {
			fmt.Println(err)
			return
		}

		dest := cmd.Flag("dest").Value.String()
		if dest == "stdout" {
			fmt.Println(string(b))
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(composeCmd)
	composeCmd.AddCommand(dumpCmd)

	//composeCmd.PersistentFlags().String("file", "system.yaml", "The system.yaml file to use")
	composeCmd.PersistentFlags().StringP("file", "f", "system.yaml", "The system.yaml file to use")

	dumpCmd.PersistentFlags().String("dest", "stdout", "The file to write the compose file to. Default is stdout.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// composeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// composeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
