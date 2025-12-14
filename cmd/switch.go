package cmd

import (
	profiles "github.com.almeidazs/abacatepay-cli/internal/config"
	"github.com.almeidazs/abacatepay-cli/internal/logger"
	"github.com.almeidazs/abacatepay-cli/internal/utils"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch",
	Args:  cobra.ExactArgs(1),
	Short: "Switch to a different profile",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := profiles.Load()

		if err != nil {
			return err
		}

		name := args[0]

		if err := config.SetCurrent(name); err != nil {
			return err
		}

		key, err := utils.GetKeyring(name)

		if err != nil {
			return err
		}

		if err := utils.SaveKeyring("current", key); err != nil {
			return err
		}

		logger.Success("Switched successfully to \"%s\"", name)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
}
