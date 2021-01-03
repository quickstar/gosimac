package reddit

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/quickstar/wally/internal/reddit"
	"github.com/spf13/cobra"
)

const (
	flagCount      = "number"
	flagQuery      = "query"
	flagResolution = "resolution"

	// DefaultCount is a default number of fetching images from sources.
	defaultCount = 10
)

// Register registers unsplash command.
func Register(root *cobra.Command, path string) {
	cmd := &cobra.Command{
		Use:     "reddit",
		Aliases: []string{"r"},
		Short:   "fetches images from https://www.reddit.com/r/wallpaper/hot",

		RunE: func(cmd *cobra.Command, args []string) error {
			n, err := cmd.Flags().GetInt(flagCount)
			if err != nil {
				return fmt.Errorf("count flag parse failed: %w", err)
			}
			pterm.Info.Printf("count: %d\n", n)

			q, err := cmd.Flags().GetString(flagQuery)
			if err != nil {
				return fmt.Errorf("query flag parse failed: %w", err)
			}
			pterm.Info.Printf("query: %s\n", q)

			s, err := cmd.Flags().GetString(flagResolution)
			if err != nil {
				return fmt.Errorf("resolution flag parse failed: %w", err)
			}
			pterm.Info.Printf("resolution: %s\n", s)

			r := reddit.New(n, q, path, s)

			if err := r.Fetch(); err != nil {
				return fmt.Errorf("reddit fetch failed %w", err)
			}

			return nil
		},
	}

	cmd.Flags().StringP(flagQuery, "q", "", "Limit selection to photos matching a search term.")
	cmd.Flags().IntP(flagCount, "n", defaultCount, "The number of photos to return")
	cmd.Flags().StringP(flagResolution, "r", "3840x2160", "Filter search results by resolution, possible values are 1920x1080 (Full HD), 2560x1440 (2K), 3840x2160 (4K), or any other resolution you like")
	root.AddCommand(cmd)
}
