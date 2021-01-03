package reddit

import (
	"github.com/1995parham/gosimac/cmd/common"
	"github.com/1995parham/gosimac/reddit"
	"github.com/spf13/cobra"
)

const (
	flagQuery      = "query"
	flagResolution = "resolution"
)

// Register registers unsplash command.
func Register(rootCommand *cobra.Command) {
	cmd := &cobra.Command{
		Use:     "reddit",
		Aliases: []string{"r"},
		Short:   "fetches images from https://www.reddit.com/r/wallpaper/hot",

		RunE: func(cmd *cobra.Command, args []string) error {
			count, err := cmd.Flags().GetInt(common.FlagCount)
			if err != nil {
				return err
			}

			q, err := cmd.Flags().GetString(flagQuery)
			if err != nil {
				return err
			}

			r, err := cmd.Flags().GetString(flagResolution)
			if err != nil {
				return err
			}

			s := &reddit.Source{
				Count:      count,
				Query:      q,
				Resolution: r,
			}

			return common.Run(s, cmd)
		},
	}

	cmd.Flags().StringP(flagQuery, "q", "", "Limit selection to photos matching a search term.")
	cmd.Flags().StringP(flagResolution, "r", "3840x2160", "Filter search results by resolution, possible values are 1920x1080 (Full HD), 2560x1440 (2K), 3840x2160 (4K), or any other resolution you like")
	rootCommand.AddCommand(cmd)
}
