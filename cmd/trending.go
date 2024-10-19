// cmd/trending.go
package cmd

import (
	"github.com/spf13/cobra"
	"yt-cli/app/internal/youtube"
)

var trendingCmd = &cobra.Command{
	Use:     "trending",
	Aliases: []string{"popular", "tr"},
	Short:   "Show trending videos",
	Long:    `Displays the current trending videos on YouTube.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := youtube.NewYouTubeClient()
		if err != nil {
			return err
		}
		return client.GetMostPopularVideos(maxResults, outputFormat)
	},
}

func init() {
	rootCmd.AddCommand(trendingCmd)
}
