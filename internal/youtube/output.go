// internal/youtube/output.go
package youtube

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"strings"
	"time"
)

func outputVideos(videos []Video, format string) error {
	switch format {
	case "json":
		return outputJSON(videos)
	case "simple":
		return outputSimple(videos)
	default:
		return outputPretty(videos)
	}
}

func outputJSON(videos []Video) error {
	encoded, err := json.MarshalIndent(videos, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(encoded))
	return nil
}

func outputSimple(videos []Video) error {
	for _, v := range videos {
		fmt.Printf("%s\n%s\n\n", v.Title, v.URL)
	}
	return nil
}

func outputPretty(videos []Video) error {
	titleColor := color.New(color.FgHiCyan, color.Bold)
	channelColor := color.New(color.FgHiBlue)
	urlColor := color.New(color.FgHiGreen)
	infoColor := color.New(color.FgHiBlack)

	for _, v := range videos {
		fmt.Println(strings.Repeat("─", 80))
		titleColor.Printf("▶ %s\n", v.Title)
		if v.Channel != "" {
			channelColor.Printf("📺 %s\n", v.Channel)
		}
		urlColor.Printf("🔗 %s\n", v.URL)
		infoColor.Printf("📅 Published %s", formatTime(v.Published))
		if v.Views != "" {
			infoColor.Printf(" • 👁 %s views", v.Views)
		}
		fmt.Println()
	}
	return nil
}

func formatTime(t time.Time) string {
	duration := time.Since(t)
	switch {
	case duration < 24*time.Hour:
		return "today"
	case duration < 48*time.Hour:
		return "yesterday"
	case duration < 7*24*time.Hour:
		return fmt.Sprintf("%d days ago", int(duration.Hours()/24))
	case duration < 30*24*time.Hour:
		return fmt.Sprintf("%d weeks ago", int(duration.Hours()/24/7))
	default:
		return t.Format("Jan 2, 2006")
	}
}

func formatViews(views uint64) string {
	switch {
	case views >= 1000000000:
		return fmt.Sprintf("%.1fB", float64(views)/1000000000)
	case views >= 1000000:
		return fmt.Sprintf("%.1fM", float64(views)/1000000)
	case views >= 1000:
		return fmt.Sprintf("%.1fK", float64(views)/1000)
	default:
		return fmt.Sprintf("%d", views)
	}
}
