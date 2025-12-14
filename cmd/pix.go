package cmd

import "github.com/spf13/cobra"

var pixCmd = &cobra.Command{
	Use: "pix",
	Short: "Manage your PIX payments",
}

func init() {
	rootCmd.AddCommand(pixCmd)
}
