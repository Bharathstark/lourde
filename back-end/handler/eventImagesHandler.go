package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Bharathstark/lourde/back-end/caching"
	"github.com/Bharathstark/lourde/back-end/database"
	"github.com/Bharathstark/lourde/back-end/logger"
	"github.com/gorilla/mux"
)

//HandleEventImages - to handle events entity
func HandleEventImages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response *Response
	if r.Method == "POST" || r.Method == "DELETE" {
		vars := mux.Vars(r)
		if !validateURL(w, r, &database.Event{}, vars["parentid"]) {
			return
		}
		if r.Method == "DELETE" && !validateURL(w, r, &database.EventImages{}, vars["id"]) {

		}
	}
	switch r.Method {
	case "GET":
		response = handleGETAllEventImages(w, r)
	case "POST":
		response = handlePOSTEventImages(w, r)
	case "DELETE":
		response = handleDELETEEventImages(w, r)
	default:
		logger.Log.Printf("Operations is not Supported")

	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)

}

func handleGETAllEventImages(w http.ResponseWriter, r *http.Request) (response *Response) {
	var eventImages []database.EventImages
	var event database.Event
	vars := mux.Vars(r)
	id := vars["id"]
	data, found := caching.GetCache("events" + id)
	if found {
		return ConstructResponse(&data, "success")
	}
	rows, error := database.GetEntity(&event, &id)
	if *rows <= 0 || *error != nil {
		return ConstructResponse(&ResponseMessage{
			Status: http.StatusInternalServerError,
			Msg:    "Invalid URL - no data found",
		}, "custom")
	}
	error = database.GetAllChildEntity(&event, &eventImages, "EventImages")
	if *error != nil {
		return ConstructResponse(error, "failed")
	}
	event.EventImages = eventImages
	caching.SetCache("events"+id, event)
	return ConstructResponse(&event, "success")

}

func handlePOSTEventImages(w http.ResponseWriter, r *http.Request) (response *Response) {

	vars := mux.Vars(r)
	parentid := vars["parentid"]
	var event database.Event
	var eventImages database.EventImages

	var arr []interface{}

	rows, error := database.GetEntity(&event, &parentid)
	if *rows <= 0 || *error != nil {
		return ConstructResponse(&ResponseMessage{
			Status: http.StatusInternalServerError,
			Msg:    "Invalid URL - no data found",
		}, "custom")
	}
	entity := event.Entity + parentid
	err := r.ParseMultipartForm(30 * MB)

	if err != nil {
		return ConstructResponse(&ResponseMessage{
			Status: http.StatusInternalServerError,
			Msg:    "Huge File size found during post",
		}, "custom")
	}

	m := r.MultipartForm
	files := m.File["files"]

	for _, file := range files {

		path, errorFile := uploadFile(&entity, file)
		if errorFile != nil {
			return ConstructResponse(&ResponseMessage{
				Status: http.StatusInternalServerError,
				Msg:    "File Upload Failed due to unknown reasons",
			}, "custom")
		}
		eventImages = database.EventImages{
			Description: r.FormValue("description"),
			Path:        path,
		}
		database.PostChildEntity(&event, &eventImages)
		arr = append(arr, eventImages)
	}
	caching.DeleteCache("events" + parentid)
	return ConstructResponse(&arr, "success")
}

func handleDELETEEventImages(w http.ResponseWriter, r *http.Request) (response *Response) {

	vars := mux.Vars(r)
	id := vars["id"]
	parent := vars["parentid"]
	var eventImages database.EventImages
	database.GetEntity(&eventImages, &id)
	_, error := database.DeleteEntity(&eventImages)
	if *error != nil {
		return ConstructResponse(error, "failed")
	}
	deleteFile(&eventImages.Path)
	caching.DeleteCache("events" + parent)
	return ConstructResponse(&eventImages, "success")

}
