package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func makeRequest() PixabayResponse {
	client := http.Client{
		Timeout: time.Duration(3 * time.Second),
	}

	req, err := http.NewRequest("GET", "https://pixabay.com/api", nil)
	if err != nil {
		panic(err)
	}

	query := req.URL.Query()
	query.Add("key", os.Getenv("PIXABAY_KEY"))
	query.Add("q", "Red Panda")

	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Could not make request")
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var pixabayResponse PixabayResponse
	err = json.Unmarshal(body, &pixabayResponse)
	if err != nil {
		panic(err)
	}

	return pixabayResponse
}

type PexelPhotoSrc struct {
	Original  string `json:"original"`
	Large2x   string `json:"large2x"`
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Small     string `json:"small"`
	Portrait  string `json:"portrait"`
	Landscape string `json:"landscape"`
	Tiny      string `json:"tiny"`
}

type PexelPhoto struct {
	ID              int    `json:"id"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	URL             string `json:"url"`
	Photographer    string `json:"photographer"`
	PhotographerURL string `json:"photographer_url"`
	PhotographerID  int    `json:"photographer_id"`
	Src             PexelPhotoSrc
	Liked           bool `json:"liked"`
}

type PexelsResponse struct {
	TotalResults int `json:"total_results"`
	Page         int `json:"page"`
	PerPage      int `json:"per_page"`
	Photos       []PexelPhoto
	NextPage     string `json:"next_page"`
}

func makeRequest2() PexelsResponse {
	client := http.Client{
		Timeout: time.Duration(3 * time.Second),
	}

	req, err := http.NewRequest("GET", "https://api.pexels.com/v1/search", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", os.Getenv("PEXELS_KEY"))

	query := req.URL.Query()
	query.Add("query", "red panda")
	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var pexelsResponse PexelsResponse

	err = json.Unmarshal(body, &pexelsResponse)
	if err != nil {
		panic(err)
	}

	return pexelsResponse
}

type UnsplashUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type UnsplashImage struct {
	Raw     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
	Small   string `json:"small"`
	Thumb   string `json:"thumb"`
}

type UnsplashPhoto struct {
	ID          string        `json:"id"`
	CreatedAt   string        `json:"created_at"`
	Width       int           `json:"width"`
	Height      int           `json:"height"`
	Color       string        `json:"color"`
	Likes       int           `json:"likes"`
	LikedByUser bool          `json:"liked_by_user"`
	Description string        `json:"description"`
	User        UnsplashUser  `json:"user"`
	Urls        UnsplashImage `json:"urls"`
}

type UnsplashResponse struct {
	Total      int `json:"total"`
	Totalpages int `json:"total_pages"`
	Results    []UnsplashPhoto
}

func makeRequest3() UnsplashResponse {
	client := http.Client{
		Timeout: time.Duration(3 * time.Second),
	}

	req, err := http.NewRequest("GET", "https://api.unsplash.com/search/photos", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Client-ID "+os.Getenv("UNSPLASH_KEY"))
	req.Header.Add("Accept-Version", "v1")

	query := req.URL.Query()
	query.Add("query", "red panda")
	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var unsplashResponse UnsplashResponse
	err = json.Unmarshal(body, &unsplashResponse)
	if err != nil {
		panic(err)
	}

	return unsplashResponse
}
