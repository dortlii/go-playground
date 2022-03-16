package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// 'videos get' subcommand
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// inputs for 'videos get' command
	getAll := getCmd.Bool("all", false, "Get all videos")
	getId := getCmd.String("id", "", "YouTube Video ID")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addID := addCmd.String("id", "", "YouTube video ID")
	addTitle := addCmd.String("title", "", "YouTube video Title")
	addUrl := addCmd.String("url", "", "YouTube video URL")
	addImageUrl := addCmd.String("imageurl", "", "YouTube video Image URL")
	addDesc := addCmd.String("desc", "", "YouTube video description")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'add' subcommands")
		os.Exit(1)
	}

	// look at the 2nd argument's value
	switch os.Args[1] {
	case "get": //if it's the 'get' command
		HandleGet(getCmd, getAll, getId)
	case "add": // if it's the 'add' command
		HandleAdd(addCmd, addID, addTitle, addUrl, addImageUrl, addDesc)
	default: // if we don't understand the input
	}
}

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {
	getCmd.Parse(os.Args[2:])

	if *all == false && *id == "" {
		fmt.Print("id is required or specify --all for all videos")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *all {
		//return all videos
		videos := getVideos()

		fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
		for _, video := range videos {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.ImageUrl, video.Description)
		}

		return
	}

	if *id != "" {
		videos := getVideos()
		id := *id
		for _, video := range videos {
			if id == video.Id {
				fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
				fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.ImageUrl, video.Description)

			}
		}
	}
}

func ValidateVideo(addCmd *flag.FlagSet, id *string, title *string, url *string, imageUrl *string, description *string) {
	if *id == "" || *title == "" || *url == "" || *imageUrl == "" || *description == "" {
		fmt.Print("all fields are required for adding a video")
		addCmd.PrintDefaults()
		os.Exit(1)
	}

	video := video{
		Id:          *id,
		Title:       *title,
		Description: *description,
		ImageUrl:    *imageUrl,
		Url:         *url,
	}

	videos := getVideos()
	videos = append(videos, video)

	saveVideos(videos)
}

func HandleAdd(addCmd *flag.FlagSet, id *string, title *string, url *string, imageurl *string, desc *string) {
	ValidateVideo(addCmd, id, title, url, imageurl, desc)
}
