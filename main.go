package main

import (
	"github.com/hansputera/kemdikbud-on-linux/cmd"
	"github.com/hansputera/kemdikbud-on-linux/constants"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:     "kemdikbud-on-linux",
	Short:   "KemdikbudOnLinux is a CLI application that help you to ran Dapodik/Erapor apps on Linux without VM",
	Long:    "KemdikbudOnLinux is a CLI application that help you to ran Dapodik/Erapor apps on Linux without VM",
	Run:     func(cmd *cobra.Command, args []string) {},
	Version: constants.APP_VERSION,
}

func main() {
	rootCmd.AddCommand(cmd.GetDapodikCmd())

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
