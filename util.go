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
	query := req.URL.Query()
	query.Add("key", os.Getenv("PIXABAY_KEY"))
	query.Add("q", "Red Panda")

	req.URL.RawQuery = query.Encode()

	if err != nil {
		panic(err)
	}

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

func makeRequest2() {
	client := http.Client{
		Timeout: time.Duration(3 * time.Second),
	}

	req, err := http.NewRequest("GET", "https://api.pexels.com/v1", nil)
	if err != nil {
		panic(err)
	}
}
