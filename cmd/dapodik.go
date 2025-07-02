package cmd

import (
	"fmt"

	"github.com/hansputera/kemdikbud-on-linux/cmd/subcmd"
	"github.com/hansputera/kemdikbud-on-linux/services/dapodik"
	"github.com/spf13/cobra"
)

var dapodikCmd = &cobra.Command{
	Use:   "dapodik",
	Short: "Manage Dapodik patches, updates, and downloads",
	Long:  "Manage Dapodik patches and downloads easily from the command line. This command helps you check for available updates, download required files, and apply patches to keep your Dapodik system up-to-date and running smoothly.",
	Run: func(cmd *cobra.Command, args []string) {
		isVersion, err := cmd.Flags().GetBool("version")
		if err != nil {
			panic(err)
		}

		if isVersion {
			version := dapodik.GetCurrentVersion()
			fmt.Printf("Currently new version: %s with %d patches\n", version.Version, len(version.Patches))
			for i, p := range version.Patches {
				fmt.Printf("%d. Patch %s -> %s\n", i+1, p.PatchName, p.PatchUrl)
			}
		}
	},
}

func GetDapodikCmd() *cobra.Command {
	dapodikCmd.AddCommand(subcmd.GetDapodikDownloadCmd())
	dapodikCmd.Flags().BoolP("version", "v", false, "")

	return dapodikCmd
}
