package main

import (
	"github.com/spf13/cobra"
	"k8s-deploy/cmd"
	"log"
)

func main() {
	command := &cobra.Command{}
	command.AddCommand(
		cmd.OSInitCmd,
		cmd.RemoveCmd,
		cmd.NewDeployRootCmd(),
		cmd.NewClearCmd(),
	)
	err := command.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
