// cmd/subscriptions.go
package cmd

import (
	"github.com/spf13/cobra"
	"yt-cli/app/internal/youtube"
)

var subscriptionsCmd = &cobra.Command{
	Use:     "subs",
	Aliases: []string{"subscriptions"},
	Short:   "Show latest videos from subscriptions",
	Long:    `Displays the most recent videos from your YouTube subscriptions.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := youtube.NewYouTubeClient()
		if err != nil {
			return err
		}
		return client.GetSubscribedChannelsVideos(maxResults, outputFormat)
	},
}

func init() {
	rootCmd.AddCommand(subscriptionsCmd)
}
