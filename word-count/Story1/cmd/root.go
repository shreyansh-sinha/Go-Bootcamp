package wordcount

import (
	"fmt"
	"os"

	wc "example.com/Story1/pkg"
	"github.com/spf13/cobra"
)

var choice bool
var rootCmd = &cobra.Command{
	Use:   "wc",
	Short: "wc is a word, line, and character count tool",
	Long:  `wc is a word, line, and character count tool that reads from the standard input or from a file and outputs the number of lines, words, and characters`,
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		if choice {
			wc.LinesCount(file)
		} else {
			fmt.Println("Not enough arguments")
			return
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&choice, "line", "l", false, "lines count flag")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error encountered '%s'", err)
		os.Exit(1)
	}
}
