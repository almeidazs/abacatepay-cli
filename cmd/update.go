package cmd

import (
	"errors"
	"fmt"

	"os"

	"github.com/blang/semver"
	"github.com/fatih/color"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/cobra"
)

const version = "1.0.0"

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Auto update AbacatePay CLI when new version is available",
	RunE: func(cmd *cobra.Command, args []string) error {
		return Update()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func Update() error {
	latest, err := CheckForUpdate()

	if err != nil {
		return err
	}

	if latest == nil {
		fmt.Printf("You are already on the latest version of AbacatePay CLI (%s)", version)

		return nil
	}

	exe, err := os.Executable()

	if err != nil {
		return err
	}

	if err := selfupdate.UpdateTo(latest.AssetURL, exe); err != nil {
		if os.IsPermission(err) {
			return errors.New("permission denied, please run the update command with sudo: sudo abacate update")
		}

		return err
	}

	fmt.Printf("%s\n\nhttps://github.com/almeidazs/abacatepay-cli", color.New(color.FgHiGreen).Sprintf("New CLI version downloaded (%v)", latest.Version.Major))

	return nil
}

func CheckForUpdate() (*selfupdate.Release, error) {
	latest, found, err := selfupdate.DetectLatest("almeidazs/abacatepay-cli")

	if err != nil {
		return nil, err
	}

	currentVersion, err := semver.Parse(version)

	if err != nil {
		return nil, err
	}

	if !found || latest.Version.LTE(currentVersion) {
		return nil, nil
	}

	return latest, nil
}
