package cmd

import (
	"fmt"
	"os"

	"github.com/quickstar/wally/cmd/bing"
	"github.com/quickstar/wally/cmd/common"
	"github.com/quickstar/wally/cmd/reddit"
	"github.com/quickstar/wally/cmd/unsplash"
	"github.com/spf13/cobra"
)

// ExitFailure status code.
const ExitFailure = 1

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	command := &cobra.Command{
		Use:     "wally",
		Short:   "Fetch the wallpaper from Bing, Reddit, Unsplash ...",
		Version: "4.0.0",
	}

	unsplash.Register(command)
	bing.Register(command)
	reddit.Register(command)

	command.PersistentFlags().StringP(common.FlagPath, "p", common.DefaultPath(), "A path to store the photos")
	command.PersistentFlags().IntP(common.FlagCount, "n", common.DefaultCount, "The number of photos to return")

	if err := command.Execute(); err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(ExitFailure)
	}
}
