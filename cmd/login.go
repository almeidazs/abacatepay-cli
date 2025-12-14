package cmd

import (
	profiles "github.com.almeidazs/abacatepay-cli/internal/config"
	"github.com.almeidazs/abacatepay-cli/internal/logger"
	"github.com.almeidazs/abacatepay-cli/internal/prompts"
	"github.com/spf13/cobra"
)

var (
	key  string
	name string
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in on Abacatepay with an API key",
	Run: func(cmd *cobra.Command, args []string) {
		if key == "" {
			input, err := prompts.AskAPIKey()

			if err != nil {
				logger.Error(err)

				return
			}

			key = input
		}

		if name == "" {
			input, err := prompts.AskProfileName()

			if err != nil {
				logger.Error(err)

				return
			}

			name = input
		}

		config, err := profiles.Load()

		if err != nil {
			logger.Error(err)

			return
		}

		if err := config.Save(name, key); err != nil {
			logger.Error(err)

			return
		}
	},
}

func init() {
	loginCmd.Flags().StringVar(&key, "key", "", "Abacate Pay's API Key")
	loginCmd.Flags().StringVar(&name, "name", "", "Name for the profile (Min 3, Max 50 chars.)")

	rootCmd.AddCommand(loginCmd)
}
