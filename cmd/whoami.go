package cmd

import (
	"errors"

	profiles "github.com.almeidazs/abacatepay-cli/internal/config"
	"github.com.almeidazs/abacatepay-cli/internal/logger"
	"github.com/spf13/cobra"
)

var raw bool

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "View the name of the current profile that you are using",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := profiles.Load()

		if err != nil {
			return err
		}

		crr := config.Current

		if crr == "" {
			return errors.New("there is no profile logged yet")
		}

		if raw {
			println(crr)

			return nil
		}

		logger.Success("You are currently using \"%s\"", crr)

		return nil
	},
}

func init() {
	whoamiCmd.Flags().BoolVar(&raw, "raw", false, "Raw profile name")

	rootCmd.AddCommand(whoamiCmd)
}
