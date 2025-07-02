package subcmd

import (
	"errors"
	"fmt"
	"path"

	"github.com/hansputera/kemdikbud-on-linux/services/dapodik"
	"github.com/hansputera/kemdikbud-on-linux/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var dapodikDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download dapodik files",
	Run: func(cmd *cobra.Command, args []string) {
		downloadType, err := cmd.Flags().GetString("download-type")

		if err != nil || len(downloadType) < 1 {
			prompt := promptui.Select{
				Label: "Download type",
				Items: []string{"Main file (Dapodik)", "Patch files (Patch)"},
			}

			_, result, err := prompt.Run()
			if err != nil {
				panic(err)
			}

			if result == "Main file (Dapodik)" {
				downloadType = "main"
			} else {
				downloadType = "patch"
			}
		}

		targetpath, err := cmd.Flags().GetString("targetpath")
		if err != nil || len(targetpath) < 1 {
			prompt := promptui.Prompt{
				Label: "Folder Path (Must be folder and exists)",
				Validate: func(s string) error {
					if !utils.IsExistsFolder(s) {
						return errors.New("path is not folder")
					}

					return nil
				},
				AllowEdit: true,
			}

			result, err := prompt.Run()
			if err != nil {
				panic(err)
			}

			targetpath = result
		} else {
			if !utils.IsExistsFolder(targetpath) {
				panic("path is not folder")
			}
		}

		version := dapodik.GetCurrentVersion()

		if downloadType == "main" {
			dapodik_path := path.Join(targetpath, fmt.Sprintf("%s.exe", version.Version))
			dapodik_vokasi_path := path.Join(targetpath, fmt.Sprintf("%s_vokasi.exe", version.Version))

			if err = utils.DownloadFile(version.Url, dapodik_path); err != nil {
				panic(err)
			}

			if err = utils.DownloadFile(version.VokasiUrl, dapodik_vokasi_path); err != nil {
				panic(err)
			}
		}

		if downloadType == "patch" {
			for _, patch := range version.Patches {
				vokasi_str := ""
				if patch.IsVokasi {
					vokasi_str = "_vokasi"
				}

				patch_path := path.Join(targetpath, fmt.Sprintf("patch_%s%s.exe", patch.PatchName, vokasi_str))
				if err = utils.DownloadFile(patch.PatchUrl, patch_path); err != nil {
					panic(err)
				}
			}
		}
	},
}

func GetDapodikDownloadCmd() *cobra.Command {
	dapodikDownloadCmd.Flags().StringP("download-type", "d", "", "d=main")
	dapodikDownloadCmd.Flags().StringP("targetpath", "p", "", "p=files")
	return dapodikDownloadCmd
}
