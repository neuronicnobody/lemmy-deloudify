package main

import (
	"encoding/json"
	"fmt"

	xtptest "github.com/dylibso/xtp-test-go"
)

type CreatePost struct {
	Name             string  `json:"name"`
	Body             *string `json:"body,omitempty"`
	Community_id     int32   `json:"community_id"`
	Url              *string `json:"url,omitempty"`
	Alt_text         *string `json:"alt_text,omitempty"`
	Honeypot         *string `json:"honeypot,omitempty"`
	Nsfw             *bool   `json:"nsfw,omitempty"`
	Language_id      *int32  `json:"language_id,omitempty"`
	Custom_thumbnail *string `json:"custom_thumbnail,omitempty"`
}

//go:export test
func test() int32 {
	body := "I AM SO LOUD. What will YOU do about it?"
	post := CreatePost{
		Name:         "Example Post",
		Body:         &body,
		Community_id: 123,
	}

	jsonPost, err := json.Marshal(post)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return 1
	}

	output := fromJson((xtptest.CallBytes("api_before_post_post", jsonPost)))
	xtptest.AssertEq("basic test", *output.Body, "I am so loud. What will you do about it?")
	return 0
}

func fromJson(data []byte) CreatePost {
	var res CreatePost
	_ = json.Unmarshal(data, &res)
	return res
}

func main() {}
