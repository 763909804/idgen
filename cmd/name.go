package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/mritd/chinaid"
	"github.com/spf13/cobra"
)

var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "Generate name",
	Long: `
Generate Chinese name, length 2-4 digits (including complex surname)`,
	Run: func(cmd *cobra.Command, args []string) {
		name := chinaid.Name()
		fmt.Println(name)
		_ = clipboard.WriteAll(name)
	},
}

func init() {
	rootCmd.AddCommand(nameCmd)
}
