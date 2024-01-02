package main

import (
	"html/template"
	"net/http"

	"github.com/JyJy32/logger"
	"github.com/gorilla/mux"
)

func server_init() {
	router := mux.NewRouter()
	router.HandleFunc("/current_video", currentVideo).Methods("GET")
	router.HandleFunc("/video", getVideo).Methods("GET")
	router.HandleFunc("/set_video/{videoID}", youtube_info_to_obs).Methods("POST")
    router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    }).Methods("GET")
    router.HandleFunc("/coffee", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusTeapot)
    }).Methods("GET")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:6969",
	}

	logger.Http("Server started on port 6969")
	srv.ListenAndServe()

}

func youtube_info_to_obs(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	videoID, ok := params["videoID"]
	if !ok || videoID == "undefined" || videoID == "" {
		return
	}
	if params["videoID"] == CurrentVideoInfo.VideoId {
		return
	}
	logger.Http("Received request")
	logger.Debug("Video ID: " + videoID)
	info, err := getVideoInfo(videoID)
	if err != nil {
		logger.Error("Error getting video")
		logger.Error(err.Error())
		return
	}
	logger.Debug("Video title: " + info.Title)

	logger.Debug("video changed")
	CurrentVideoInfo = *info
}

func getVideo(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("video.gohtml")
	if err != nil {
		logger.Error("Error parsing template")
		logger.Error(err.Error())
		return
	}
	//logger.Http("Serving video info")
	err = template.Execute(w, CurrentVideoInfo)
	if err != nil {
		logger.Error("Error executing template")
		logger.Error(err.Error())
		return
	}
}

func currentVideo(w http.ResponseWriter, r *http.Request) {
	logger.Http("Serving current video")
	templ, err := template.ParseFiles("index.gohtml")
	if err != nil {
		logger.Error("Error parsing template")
		logger.Error(err.Error())
		panic(err)
	}

	templ.Execute(w, nil)
}
