package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/alexeyco/teletype-crawler/pkg/cleaner"
	"github.com/alexeyco/teletype-crawler/pkg/client"
	"github.com/alexeyco/teletype-crawler/pkg/saver"
	"github.com/alexeyco/teletype-crawler/pkg/teletype"
	"github.com/alexeyco/teletype-crawler/pkg/wordcounter"
)

var blogID int64

func init() {
	rootCmd.Flags().Int64Var(&blogID, "blog-id", 0, "teletype blog id")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "teletype-crawler",
	Short: "A tool to save all specified teletype blog content",
	Run: func(cmd *cobra.Command, args []string) {
		if blogID == 0 {
			printError(`--blog-id flag is required`)
			return
		}

		tt := teletype.NewProvider(client.NewClient(), cleaner.NewCleaner())
		wc := wordcounter.NewCounter()
		s := saver.NewSaver(fmt.Sprintf("./articles/%d", blogID))

		if err := s.Clean(); err != nil {
			printError(err.Error())
		}

		err := tt.Each(context.Background(), blogID, func(a teletype.Article) error {
			if wc.Count(a.Text) < 32 {
				return nil
			}

			return s.Save(fmt.Sprintf(`%d.txt`, a.ID), a.Text)
		})

		if err != nil {
			printError(err.Error())
			return
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func printError(s string) {
	println(s)
	println(``)
}
