package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Bharathstark/lourde/back-end/database"
)

//SPAHandler is used to serve single Page Applications
type SPAHandler struct {
	StaticPath string
	IndexPath  string
}

func (h SPAHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	if r.URL.Path == "" {

		http.Error(w, "Error thrown", http.StatusBadRequest)
		return
	}
	path = filepath.Join(h.StaticPath, path)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {

		http.ServeFile(w, r, filepath.Join(h.StaticPath, h.IndexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.FileServer(http.Dir(h.StaticPath)).ServeHTTP(w, r)
}

//HandleHealthAPI is used to handle health API's
func HandleHealthAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var count [5]int64
	event := database.Event{
		Entity: "events",
	}
	database.GetCountEntity(&event, &count[0])
	event.Entity = "achievements"
	database.GetCountEntity(&event, &count[1])
	notification := database.Notification{
		Entity: "notification",
	}
	database.GetCountEntity(&notification, &count[2])
	notification.Entity = "videos"
	database.GetCountEntity(&notification, &count[3])
	student := database.Student{}
	database.GetCountEntity(&student, &count[4])
	var countes = map[string]int64{
		"events":       count[0],
		"achievements": count[1],
		"notification": count[2],
		"videos":       count[3],
		"academics":    count[4],
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ConstructResponse(&countes, "success"))

}
