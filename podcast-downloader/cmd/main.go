package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"podcast-downloader/internal/downloader"
	"podcast-downloader/internal/rss"
	"podcast-downloader/internal/scheduler"
)

func main() {
	// Define multiple RSS feed URLs
	feedURLs := []string{
		"https://feeds.megaphone.fm/POLTD2316968013",
		"https://www.spreaker.com/show/4552198/episodes/feed",
		// Add more feed URLs here
	}

	// Schedule tasks for each feed
	for _, feedURL := range feedURLs {
		scheduler.Schedule(24*time.Hour, func() {
			fmt.Printf("Checking for new episodes from feed: %s\n", feedURL)
			feed, err := rss.FetchRSSFeed(feedURL)
			if err != nil {
				fmt.Printf("Error fetching RSS feed: %v\n", err)
				return
			}

			// Create a folder for the podcast
			podcastFolder := feed.Channel.Title
			if err := os.MkdirAll(podcastFolder, os.ModePerm); err != nil {
				fmt.Printf("Error creating folder: %v\n", err)
				return
			}

			// Track downloaded episodes
			recordFile := fmt.Sprintf("%s/downloaded.txt", podcastFolder)
			downloadedEpisodes := make(map[string]bool)
			if _, err := os.Stat(recordFile); err == nil {
				file, err := os.Open(recordFile)
				if err == nil {
					defer file.Close()
					var episodeTitle string
					for {
						_, err := fmt.Fscanf(file, "%s\n", &episodeTitle)
						if err != nil {
							break
						}
						downloadedEpisodes[episodeTitle] = true
					}
				}
			}

			// Download only the latest episode
			if len(feed.Channel.Items) > 0 {
				latestEpisode := feed.Channel.Items[0]
				if !downloadedEpisodes[latestEpisode.Title] {
					// Sanitize the filename to remove invalid characters
					sanitizedTitle := sanitizeFilename(latestEpisode.Title)
					fileName := fmt.Sprintf("%s/%s.mp3", podcastFolder, sanitizedTitle)
					fmt.Printf("Downloading latest episode: %s\n", latestEpisode.Title)
					err := downloader.DownloadFile(latestEpisode.Enclosure.URL, fileName)
					if err == nil {
						file, err := os.OpenFile(recordFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err == nil {
							defer file.Close()
							fmt.Fprintf(file, "%s\n", latestEpisode.Title)
						}
					} else {
						fmt.Printf("Error downloading file: %v\n", err)
					}
				}
			}
		})
	}

	// Keep the application running
	select {}
}

// Helper function to extract file extension
func getFileExtension(url string) string {
	parts := strings.Split(url, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return ""
}

// Helper function to sanitize filenames
func sanitizeFilename(name string) string {
	invalidChars := []string{":", "<", ">", "|", "?", "*", "\\", "/", "[", "]"}
	for _, char := range invalidChars {
		name = strings.ReplaceAll(name, char, "_")
	}
	return name
}
