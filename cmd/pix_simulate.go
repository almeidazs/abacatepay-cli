package cmd

import (
	"github.com.almeidazs/abacatepay-cli/internal/http"
	"github.com.almeidazs/abacatepay-cli/internal/logger"
	"github.com.almeidazs/abacatepay-cli/internal/utils"
	"github.com/spf13/cobra"
)

var simulateCmd = &cobra.Command{
	Use:   "simulate",
	Args:  cobra.ExactArgs(1),
	Short: "Simulates payment for a QRCode Pix",
	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := utils.GetAPIKey()

		if err != nil {
			return err
		}

		data, err := http.NewClient(key).Post(
			"/pixQrCode/simulate-payment?id="+args[0], nil,
		)

		if err != nil {
			return err
		}

		brCode, _ := data["brCode"].(string)

		logger.Success("brCode \"%s\"", brCode)

		return nil
	},
}

func init() {
	pixCmd.AddCommand(simulateCmd)
}
