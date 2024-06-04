package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Crawl is a simple tool to crawl websites and output the occurence of words in a csv file.",
	Long:  `Crawl is a simple tool to crawl websites and output the occurence of words in a csv file.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
