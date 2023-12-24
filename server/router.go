package main

import (
	"html/template"
	"net/http"

	"github.com/JyJy32/logger"
	"github.com/gorilla/mux"
)

func server_init() {
	router := mux.NewRouter()
	router.HandleFunc("/current_video", current_video).Methods("GET")
	router.HandleFunc("/video", get_videohtml).Methods("GET")
	router.HandleFunc("/set_video/{videoID}", youtube_info_to_obs).Methods("POST")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:6969",
	}

	logger.Http("Server started on port 6969")
	srv.ListenAndServe()

}

func youtube_info_to_obs(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
    if params["videoID"] == "" {
        return 
    }
    if params["videoID"] == currentVideo.VideoId {
        return
    }
    logger.Http("Received request")
	logger.Debug("Video ID: " + params["videoID"])
	video, err := get_video(params["videoID"])
	if err != nil {
		logger.Error("Error getting video")
		logger.Error(err.Error())
		return
	}
	info := VideoInfo{
		video.Id,
		video.Snippet.Title,
		video.Snippet.ChannelTitle,
		video.Snippet.Thumbnails.Default.Url,
	}

	logger.Debug("Video title: " + info.Title)

    logger.Debug("video changed")
    currentVideo = info
}

func get_videohtml(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("video.gohtml")
	if err != nil {
		logger.Error("Error parsing template")
		logger.Error(err.Error())
		return
	}
	//logger.Http("Serving video info")
	err = template.Execute(w, currentVideo)
	if err != nil {
		logger.Error("Error executing template")
		logger.Error(err.Error())
		return
	}
}

func current_video(w http.ResponseWriter, r *http.Request) {
	logger.Http("Serving current video")
	templ, err := template.ParseFiles("index.gohtml")
	if err != nil {
		logger.Error("Error parsing template")
		logger.Error(err.Error())
		panic(err)
	}

	templ.Execute(w, nil)
}
