// internal/youtube/client.go
package youtube

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"google.golang.org/api/youtube/v3"
)

type Video struct {
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Channel   string    `json:"channel,omitempty"`
	Published time.Time `json:"published"`
	Views     string    `json:"views,omitempty"`
}

type YouTubeClient struct {
	service *youtube.Service
	spinner *spinner.Spinner
}

func NewYouTubeClient() (*YouTubeClient, error) {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = "Loading "
	s.Color("cyan")

	client, err := GetAuthenticatedClient()
	if err != nil {
		return nil, err
	}

	return &YouTubeClient{
		service: client,
		spinner: s,
	}, nil
}

func (c *YouTubeClient) GetChannelVideos(channelId, channelTitle string, maxResults int) ([]Video, error) {
	c.spinner.Suffix = fmt.Sprintf(" fetching videos from channel %s...", channelTitle)
	c.spinner.Start()
	defer c.spinner.Stop()

	call := c.service.Search.List([]string{"snippet"}).
		ChannelId(channelId).
		MaxResults(int64(maxResults)).
		Order("date")

	response, err := call.Do()
	if err != nil {
		return nil, fmt.Errorf("error fetching videos for channel %s: %v", channelTitle, err)
	}

	var videos []Video
	for _, item := range response.Items {
		publishedAt, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		if err != nil {
			return nil, fmt.Errorf("error parsing published date for video %s: %v", item.Id.VideoId, err)
		}
		videos = append(videos, Video{
			Title:     item.Snippet.Title,
			URL:       fmt.Sprintf("https://youtube.com/watch?v=%s", item.Id.VideoId),
			Channel:   item.Snippet.ChannelTitle,
			Published: publishedAt,
		})
	}

	return videos, nil
}

func (c *YouTubeClient) GetSubscribedChannelsVideos(maxResults int, format string) error {
	c.spinner.Suffix = " fetching your subscriptions..."
	c.spinner.Start()
	defer c.spinner.Stop()

	subscriptions, err := c.service.Subscriptions.List([]string{"snippet"}).
		Mine(true).
		MaxResults(int64(maxResults)).
		Do()

	if err != nil {
		return fmt.Errorf("error fetching subscriptions: %v", err)
	}

	var allVideos []Video
	for _, sub := range subscriptions.Items {
		videos, err := c.GetChannelVideos(sub.Snippet.ResourceId.ChannelId, sub.Snippet.Title, maxResults)
		if err != nil {
			continue
		}
		allVideos = append(allVideos, videos...)
	}

	return outputVideos(allVideos, format)
}

func (c *YouTubeClient) GetMostPopularVideos(maxResults int, format string) error {
	c.spinner.Suffix = " fetching trending videos..."
	c.spinner.Start()
	defer c.spinner.Stop()

	videos, err := c.service.Videos.List([]string{"snippet", "statistics"}).
		Chart("mostPopular").
		MaxResults(int64(maxResults)).
		Do()

	if err != nil {
		return fmt.Errorf("error fetching popular videos: %v", err)
	}

	var formattedVideos []Video
	for _, v := range videos.Items {
		publishedAt, err := time.Parse(time.RFC3339, v.Snippet.PublishedAt)
		if err != nil {
			return fmt.Errorf("error parsing published date for video %s: %v", v.Id, err)
		}

		formattedVideos = append(formattedVideos, Video{
			Title:     v.Snippet.Title,
			URL:       fmt.Sprintf("https://youtube.com/watch?v=%s", v.Id),
			Channel:   v.Snippet.ChannelTitle,
			Published: publishedAt,
			Views:     formatViews(v.Statistics.ViewCount),
		})
	}

	return outputVideos(formattedVideos, format)
}
