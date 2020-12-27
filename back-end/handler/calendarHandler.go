package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Bharathstark/lourde/back-end/caching"
	"github.com/Bharathstark/lourde/back-end/database"
	"github.com/gorilla/mux"
)

//HandleCalendar is used  to Handle Notification End Point
func HandleCalendar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response *Response
	switch r.Method {
	case "GET":
		response = handleGETAllCalendar(w, r)
	case "POST":
		response = handlePOSTCalendar(w, r)
	case "PUT":
		response = handlePUTCalendar(w, r)
	case "DELETE":
		response = handleDELETECalendar(w, r)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)

}

func handleGETAllCalendar(w http.ResponseWriter, r *http.Request) (response *Response) {
	var calendar []database.Calendar
	data, found := caching.GetCache("calendar")
	if found {
		return ConstructResponse(&data, "success")
	}
	_, error := database.GetAllEntity(&calendar)
	if *error != nil {
		return ConstructResponse(error, "failed")
	}
	caching.SetCache("calendar", calendar)
	return ConstructResponse(&calendar, "success")
}

func handlePOSTCalendar(w http.ResponseWriter, r *http.Request) (response *Response) {
	var calendar database.Calendar
	decerror := decodeJSONBody(w, r, &calendar)
	if decerror != nil {
		return ConstructResponse(decerror, "custom")
	}
	_, error := database.PostEntity(&calendar)
	if *error != nil {
		return ConstructResponse(error, "failed")
	}
	caching.DeleteCache("calendar")
	return ConstructResponse(&calendar, "success")

}

func handlePUTCalendar(w http.ResponseWriter, r *http.Request) (response *Response) {

	vars := mux.Vars(r)
	id := vars["id"]
	var calendar database.Calendar
	database.GetEntity(&calendar, &id)
	var updatecalendar database.Notification
	decerror := decodeJSONBody(w, r, &updatecalendar)
	if decerror != nil {
		return ConstructResponse(decerror, "failed")
	}
	_, error := database.UpdateEntity(&calendar, &updatecalendar)
	if *error != nil {
		return ConstructResponse(error, "failed")
	}
	caching.DeleteCache("calendar")
	return ConstructResponse(&calendar, "success")

}

func handleDELETECalendar(w http.ResponseWriter, r *http.Request) (response *Response) {

	vars := mux.Vars(r)
	id := vars["id"]
	var updatecalendar database.Calendar
	database.GetEntity(&updatecalendar, &id)
	_, error := database.DeleteEntity(&updatecalendar)
	if *error != nil {
		return ConstructResponse(error, "failed")
	}
	caching.DeleteCache("calendar")
	return ConstructResponse(&updatecalendar, "success")
}
