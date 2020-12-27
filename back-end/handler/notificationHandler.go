package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Bharathstark/lourde/back-end/caching"
	"github.com/Bharathstark/lourde/back-end/database"
	"github.com/gorilla/mux"
)

var noticationmap = map[string]int{
	"notification": 1,
	"banner":       10,
	"videos":       1,
}

//HandleNotifications is to Handle Notification End Point
func HandleNotifications(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var response *Response
	switch r.Method {
	case "GET":
		response = handleGETAllNotifications(w, r)
	case "POST":
		response = handlePOSTNotification(w, r)
	case "PUT":
		response = handlePUTNotification(w, r)
	case "DELETE":
		response = handleDELETENotification(w, r)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)

}

func handleGETAllNotifications(w http.ResponseWriter, r *http.Request) (response *Response) {
	var notifications []database.Notification
	vars := mux.Vars(r)
	notication := database.Notification{
		Entity: vars["module"],
	}
	data, found := caching.GetCache(notication.Entity)
	if found {
		return ConstructResponse(&data, "success")
	}
	_, error := database.GetAllEntityFilterOrder(&notication, &notifications, "notification.date desc,notification.id")
	if *error != nil {
		return ConstructResponse(error, "failed")
	}
	caching.SetCache(notication.Entity, &notifications)
	return ConstructResponse(&notifications, "success")

}

func handlePOSTNotification(w http.ResponseWriter, r *http.Request) (response *Response) {

	var notification database.Notification
	vars := mux.Vars(r)
	entity := vars["module"]
	decodeJSONBody(w, r, &notification)
	notification.Entity = entity
	_, error := database.PostEntity(&notification)
	if *error != nil {
		return ConstructResponse(error, "failed")
	}
	caching.DeleteCache(entity)
	return ConstructResponse(&notification, "success")

}

func handlePUTNotification(w http.ResponseWriter, r *http.Request) (response *Response) {

	vars := mux.Vars(r)
	id := vars["id"]
	var notification database.Notification
	database.GetEntity(&notification, &id)
	var updatenotification database.Notification
	decerror := decodeJSONBody(w, r, &updatenotification)
	if decerror != nil {
		return ConstructResponse(decerror, "custom")
	}
	_, error := database.UpdateEntity(&notification, &updatenotification)
	if *error != nil {
		return ConstructResponse(error, "failed")
	}
	caching.DeleteCache(notification.Entity)
	return ConstructResponse(&notification, "success")

}

func handleDELETENotification(w http.ResponseWriter, r *http.Request) (response *Response) {

	vars := mux.Vars(r)
	id := vars["id"]
	var notification database.Notification
	database.GetEntity(&notification, &id)
	_, error := database.DeleteEntity(&notification)
	if *error != nil {
		return ConstructResponse(error, "failed")
	}
	caching.DeleteCache(notification.Entity)
	return ConstructResponse(&notification, "success")
}
