package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/hadlow/sendex/config"
	"github.com/hadlow/sendex/internal/file"
)

var newCmd = &cobra.Command{
	Use:   "new [FILE]",
	Short: "Create a Sendex request file",
	Long:  ``,

	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		var template []byte = config.DefaultTemplate

		isPost, _ := cmd.Flags().GetBool("post")

		if isPost {
			template = config.PostTemplate
		}

		err := file.NewWithTemplate(args[0], template)

		if err != nil {
			log.Fatal("Error: ", err)
		}

		fmt.Println("File created at:", args[0])
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().BoolP("post", "p", false, "Use the post template")
}
