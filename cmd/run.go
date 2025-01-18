package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/hadlow/sendex/internal/display"
	"github.com/hadlow/sendex/internal/file"
	"github.com/hadlow/sendex/internal/helpers"
	"github.com/hadlow/sendex/internal/request"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [FILE]",
	Short: "Run a request file",
	Long:  ``,

	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		argsMap, err := helpers.CreateArgsmap(args[1:])
		if err != nil {
			os.Exit(1)
		}

		req, err := file.Get(path)
		if err != nil {
			os.Exit(1)
		}

		response, err := request.Run(*req, argsMap)
		if err != nil {
			os.Exit(1)
		}

		displayConfig := display.NewDisplayConfig(true, true, true)
		displayConfig.Request = req

		if s, _ := cmd.Flags().GetBool("status"); s {
			displayConfig.ShowHead = false
			displayConfig.ShowBody = false
		}

		if s, _ := cmd.Flags().GetBool("body"); s {
			displayConfig.ShowStatus = false
			displayConfig.ShowHead = false
		}

		if s, _ := cmd.Flags().GetBool("head"); s {
			displayConfig.ShowStatus = false
			displayConfig.ShowBody = false
		}

		display.Response(response, displayConfig)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().BoolP("status", "s", false, "Show only the status")
	runCmd.Flags().BoolP("body", "b", false, "Show only the body")
	runCmd.Flags().BoolP("head", "e", false, "Show only the headers")
}
