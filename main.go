package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/joho/godotenv"
)

type PixabayResponse struct {
	Total     int
	TotalHits int
	Hits      []struct {
		ID              int `json:"id"`
		PageURL         string
		Type            string
		Tags            string
		PreviewURL      string
		WebformatURL    string
		WebformatWidth  int
		WebformatHeight int
		LargeImageURL   string
		ImageWidth      int
		ImageHeight     int
		ImageSize       int
		Views           int
		Downloads       int
		Favorites       int
		Likes           int
		Comments        int
		UserID          int `json:"user_id"`
		User            string
		UserImageURL    string
	}
}

type Page struct {
	Title    string
	Pixabay  PixabayResponse
	Pexel    PexelsResponse
	Unsplash UnsplashResponse
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// fs := http.FileServer(http.Dir("/public"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("REQ: %s\n", r.URL.Path)

		// http.Handle("/public", fs)

		if r.URL.Path == "/" {
			fmt.Println("root")
			page, err := ioutil.ReadFile("public/index.html")
			if err != nil {
				panic(err)
			}

			td := Page{"Main", makeRequest(), makeRequest2(), makeRequest3()}
			template, err := template.New("page").Parse(string(page))
			if err != nil {
				panic(err)
			}

			if err := template.Execute(w, td); err != nil {
				panic(err)
			}
		} else if r.URL.Path == "/public/style.css" {
			w.Header().Add("content-type", "text/css")
			page, _ := ioutil.ReadFile("public/style.css")
			io.WriteString(w, string(page))
		}
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
