package cmd

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var goCmd = &cobra.Command{
	Use: "go [shortcut]",
	Args: cobra.MaximumNArgs(1),
	ValidArgs: []string{"docs", "ref"},
	Short: "Open AbacatePay pages quickly",
	RunE: func(cmd *cobra.Command, args []string) error {
		url := "https://www.abacatepay.com/app"

		if len(args) == 1 {
			switch args[0] {
			case "docs":
				url = "https://docs.abacatepay.com"
			case "ref":
				url = "https://docs.abacatepay.com/api-reference"
			default:
				return fmt.Errorf("unknown shortcut \"%s\"", args[0])
			}
		}

		return open(url)
	},
}

func init() {
	rootCmd.AddCommand(goCmd)
}

func open(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}

	return cmd.Start()
}
