package main

import (
	"encoding/json"
	"io/ioutil"
)

type video struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	Url         string `json:"url"`
}

func getVideos() (videos []video) {
	fileBytes, err := ioutil.ReadFile("./videos.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &videos)

	if err != nil {
		panic(err)
	}

	/*
		fileContent := string(fileBytes)
		fmt.Println(fileContent)

	*/

	/*
			video1 := video{
				Id:          "evyLwK5C2dw",
				Title:       "Kubernetes on Azure",
				ImageUrl:    "https://i.ytimg.com/vi/evyLwk5C2dws",
				Url:         "alfja",
				Description: " ",
			}

			video2 := video{
				Id:          "evyLwK5C2dw",
				Title:       "Kubernetes on Azure",
				ImageUrl:    "https://i.ytimg.com/vi/evyLwk5C2dws",
				Url:         "alfja",
				Description: " ",
			}

		return []video{video1, video2}
	*/

	return videos
}

func saveVideos(videos []video) {
	videoBytes, err := json.Marshal(videos)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./videos-updated.json", videoBytes, 0644)

	if err != nil {
		panic(err)
	}
}
