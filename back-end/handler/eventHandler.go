package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Bharathstark/lourde/back-end/caching"
	"github.com/Bharathstark/lourde/back-end/database"
	"github.com/Bharathstark/lourde/back-end/logger"
	"github.com/gorilla/mux"
)

const (
	//MB - Size
	MB = 1 << 20
)

var eventmodules = map[string]int{
	"events":       1,
	"achievements": 10,
}

//HandleEvents - to handle events entity
func HandleEvents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var response *Response
	vars := mux.Vars(r)

	switch r.Method {
	case "GET":
		if len(vars["id"]) > 0 {
			response = handleGETAllEventImages(w, r)
		} else {
			response = handleGETAllEvents(w, r)
		}

	case "POST":
		response = handlePOSTEvents(w, r)
	case "PUT":
		response = handlePUTEvents(w, r)
	case "DELETE":
		response = handleDELETEEvents(w, r)
	default:
		logger.Log.Printf("Operations is not Supported")

	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)

}

func handleGETAllEvents(w http.ResponseWriter, r *http.Request) *Response {

	vars := mux.Vars(r)
	var events []database.Event
	event := database.Event{
		Entity: vars["module"],
	}
	data, found := caching.GetCache(event.Entity)
	if found {
		return ConstructResponse(&data, "success")
	}
	_, error := database.GetAllEntityFilterOrder(&event, &events, "events.date desc,events.id")
	if *error != nil {
		return ConstructResponse(error, "failed")
	}
	caching.SetCache(event.Entity, events)
	return ConstructResponse(&events, "success")

}

func handlePOSTEvents(w http.ResponseWriter, r *http.Request) *Response {

	var event database.Event
	vars := mux.Vars(r)
	entity := vars["module"]
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseMultipartForm(20 * MB)
	if err != nil {
		return ConstructResponse(&ResponseMessage{
			Status: http.StatusInternalServerError,
			Msg:    "Huge File size found during post",
		}, "custom")

	}

	m := r.MultipartForm
	//get the *fileheaders
	files := m.File["files"]

	path, errorFile := uploadFile(&entity, files[0])
	if errorFile != nil {

		return ConstructResponse(&ResponseMessage{
			Status: http.StatusInternalServerError,
			Msg:    "Huge File size found during post",
		}, "custom")
	}
	createddate, _ := strconv.Atoi(r.FormValue("date"))
	event = database.Event{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Path:        path,
		Entity:      entity,
		Date:        createddate,
	}
	caching.DeleteCache(event.Entity)
	_, error := database.PostEntity(&event)

	if *error != nil {
		return ConstructResponse(error, "failed")
	}
	return ConstructResponse(&event, "success")

}

func handlePUTEvents(w http.ResponseWriter, r *http.Request) (response *Response) {
	vars := mux.Vars(r)
	id := vars["id"]
	var oldevent database.Event
	var event database.Event
	rows, error := database.GetEntity(&oldevent, &id)
	if *rows <= 0 || *error != nil {
		return ConstructResponse(&ResponseMessage{
			Status: http.StatusInternalServerError,
			Msg:    "Invalid URL - no data found",
		}, "custom")
	}
	entity := "events"
	err := r.ParseMultipartForm(20 * MB)
	if err != nil {
		return ConstructResponse(&ResponseMessage{
			Status: http.StatusInternalServerError,
			Msg:    "Huge File size found during post",
		}, "failed")
	}
	m := r.MultipartForm

	//get the *fileheaders
	files := m.File["files"]
	if len(files) > 0 {
		deleteFile(&oldevent.Path)
		path, errorFile := uploadFile(&entity, files[0])
		if errorFile != nil {
			return ConstructResponse(&ResponseMessage{
				Status: http.StatusInternalServerError,
				Msg:    "Huge File size found during post",
			}, "custom")
		}
		event = database.Event{
			Path: path,
		}

	}
	caching.DeleteCache(oldevent.Entity)
	caching.DeleteCache("events" + id)
	createddate, _ := strconv.Atoi(r.FormValue("date"))
	event.Name = r.FormValue("name")
	event.Description = r.FormValue("description")
	event.Date = createddate

	_, errorpost := database.UpdateEntity(&oldevent, &event)

	if *errorpost != nil {

		return ConstructResponse(errorpost, "failed")
	}
	return ConstructResponse(&oldevent, "success")
}

func handleDELETEEvents(w http.ResponseWriter, r *http.Request) (response *Response) {

	vars := mux.Vars(r)
	id := vars["id"]
	var event database.Event
	var images []database.EventImages
	rows, error := database.GetEntity(&event, &id)
	if *rows <= 0 || *error != nil {

		return ConstructResponse(&ResponseMessage{
			Status: http.StatusInternalServerError,
			Msg:    "Invalid URL - no data found",
		}, "custom")
	}
	database.GetAllChildEntity(&event, &images, "EventImages")
	for _, j := range images {
		database.DeleteEntity(&j)
		deleteFile(&j.Path)
	}
	caching.DeleteCache(event.Entity)
	caching.DeleteCache("events" + id)
	_, error = database.DeleteEntity(&event)
	if *error != nil {
		return ConstructResponse(error, "failed")
	}
	deleteFile(&event.Path)
	return ConstructResponse(&event, "success")

}
