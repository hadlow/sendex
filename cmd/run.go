package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/hadlow/sendex/internal/display"
	"github.com/hadlow/sendex/internal/request"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [FILE]",
	Short: "Run a request file",
	Long:  ``,

	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		response, err := request.Run(args[0])

		if err != nil {
			os.Exit(1)
		}

		display.Response(response)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().BoolP("status", "s", false, "Show only the status")
	runCmd.Flags().BoolP("body", "b", false, "Show only the body")
	runCmd.Flags().BoolP("headers", "e", false, "Show only the headers")
}
