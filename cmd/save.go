package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/hadlow/sendex/internal/file"
	"github.com/hadlow/sendex/internal/helpers"
	"github.com/hadlow/sendex/internal/output"
	"github.com/hadlow/sendex/internal/request"
)

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save a response to a file",
	Long:  ``,
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

		response, err := request.Run(&req, argsMap)
		if err != nil {
			os.Exit(1)
		}

		// outputPath, err := cmd.Flags().GetString("output")

		outputConfig := output.NewOutputConfig()
		outputConfig.Request = &req
		outputConfig.Path = path + ".out"

		if s, _ := cmd.Flags().GetBool("status"); s {
			outputConfig.ShowHead = false
			outputConfig.ShowBody = false
		}

		if s, _ := cmd.Flags().GetBool("body"); s {
			outputConfig.ShowStatus = false
			outputConfig.ShowHead = false
		}

		if s, _ := cmd.Flags().GetBool("head"); s {
			outputConfig.ShowStatus = false
			outputConfig.ShowBody = false
		}

		output.Save(response, outputConfig)
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)

	runCmd.Flags().StringP("output", "o", "", "The filepath to save the output to")
}
