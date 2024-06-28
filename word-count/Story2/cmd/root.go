/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	wc "example.com/Story2/pkg"
	"github.com/spf13/cobra"
)

var wordFlag bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wc",
	Short: "wc is a word, line, and character count tool",
	Long:  `wc is a word, line, and character count tool that reads from the standard input or from a file and outputs the number of lines, words, and characters`,
	Run: func(cmd *cobra.Command, args []string) {

		fileName := args[0]
		if wordFlag {
			wc.WordsCount(fileName)
		} else {
			fmt.Println("Not enough arguments")
			return
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.Story2.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolVarP(&wordFlag, "word", "w", false, "words count in file")
}
