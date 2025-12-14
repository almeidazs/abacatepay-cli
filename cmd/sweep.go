package cmd

import (
	"github.com.almeidazs/abacatepay-cli/internal/config"
	"github.com.almeidazs/abacatepay-cli/internal/logger"
	"github.com.almeidazs/abacatepay-cli/internal/utils"
	"github.com/spf13/cobra"
)

var sweepCmd = &cobra.Command{
	Use: "sweep",
	Short: "Remove all logged AbacatePay accounts",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := config.Sweep(); err != nil {
			return err
		}

		if err := utils.Sweep(); err != nil {
			return err
		}

		logger.Success("Profiles were successfully deleted")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(sweepCmd)
}
