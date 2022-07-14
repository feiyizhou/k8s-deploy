package cmd

import (
	"github.com/spf13/cobra"
	"k8s-deploy/common"
	"log"
	"os"
	"os/exec"
)

var OSInitCmd = &cobra.Command{
	Use:   "os-init",
	Short: "os-init",
	Long:  "os-init",
	Run: func(cmd *cobra.Command, args []string) {
		initArgs := []string{common.GetOsInitShellPath()}
		initArgs = append(initArgs, args...)
		command := exec.Command("sh", initArgs...)
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		err := command.Run()
		if err != nil {
			log.Fatalf("Os init failed, err: %v \n", err)
		}
	},
}
