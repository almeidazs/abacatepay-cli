package cmd

import (
	"github.com.almeidazs/abacatepay-cli/internal/http"
	"github.com.almeidazs/abacatepay-cli/internal/logger"
	"github.com.almeidazs/abacatepay-cli/internal/utils"
	"github.com/spf13/cobra"
)

var pixStatusCmd = &cobra.Command{
	Use:   "status",
	Args:  cobra.ExactArgs(1),
	Short: "Check QRCode Pix payment status",
	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := utils.GetAPIKey()

		if err != nil {
			return err
		}

		id := args[0]

		client := http.NewClient(key)

		data, err := client.Get("/pixQrCode/check?id=" + id)

		if err != nil {
			return err
		}

		status, _ := data["status"].(string)

		logger.Success("Your PIX \"%s\" status is \"%s\"\n", id, status)

		return nil
	},
}

func init() {
	pixCmd.AddCommand(pixStatusCmd)
}
