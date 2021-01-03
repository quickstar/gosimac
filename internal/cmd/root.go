package cmd

import (
	"log"
	"os"
	"path"

	"github.com/adrg/xdg"
	"github.com/pterm/pterm"
	"github.com/quickstar/wally/internal/cmd/bing"
	"github.com/quickstar/wally/internal/cmd/reddit"
	"github.com/quickstar/wally/internal/cmd/unsplash"
	"github.com/spf13/cobra"
)

// ExitFailure status code.
const (
	ExitFailure = 1
	// DirectoryPermission used for creating wally Directory.
	// nolint: gofumpt
	DirectoryPermission os.FileMode = 0755
)

// DefaultPath is a default path for storing the wallpapers.
func DefaultPath() string {
	p := path.Join(xdg.UserDirs.Pictures, "wally")
	if _, err := os.Stat(p); err != nil {
		if err := os.Mkdir(p, DirectoryPermission); err != nil {
			log.Fatalf("os.Mkdir: %v", err)
		}
	}

	return p
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// nolint: exhaustruct
	root := &cobra.Command{
		Use:   "wally",
		Short: "Fetch the wallpaper from Bings, Reddit, Unsplash...",
	}

	var path string

	root.PersistentFlags().StringVarP(&path, "path", "p", DefaultPath(), "A path to where photos are stored")

	unsplash.Register(root, path)
	reddit.Register(root, path)
	bing.Register(root, path)

	if err := root.Execute(); err != nil {
		pterm.Error.Println(err.Error())
		os.Exit(ExitFailure)
	}
}
