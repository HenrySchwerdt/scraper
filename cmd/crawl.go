package cmd

import (
	"fmt"

	"github.com/HenrySchwerdt/scraper/crawler"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(crawlCmd)
	crawlCmd.Flags().StringP("in", "i", "", "Input file txt with each url in each line")
	crawlCmd.Flags().StringP("out", "o", "", "Output file csv with the word count of each url")
}

var crawlCmd = &cobra.Command{
	Use:   "c",
	Short: "Crawls a list of webpages",

	Run: func(cmd *cobra.Command, args []string) {
		in, err := cmd.Flags().GetString("in")
		if err != nil {
			fmt.Println(err)
			return
		}
		out, err := cmd.Flags().GetString("out")
		if err != nil {
			fmt.Println(err)
			return
		}

		if in == "" || out == "" {
			fmt.Println("Please provide both input and output files")
			return
		}

		crawler.NewCrawler(in, out).Crawl()
	},
}
