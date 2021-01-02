package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Bharathstark/lourde/back-end/caching"
	"github.com/Bharathstark/lourde/back-end/database"
	"github.com/gorilla/mux"
)

var student = map[string]int{
	"10": 10,
	"12": 10,
}

//HandleStudents is to Handle Notification End Point
func HandleStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response *Response
	switch r.Method {
	case "GET":
		response = handleGETAllStudents(w, r)
	case "POST":
		response = handlePOSTStudent(w, r)
	case "PUT":
		response = handlePUTStudent(w, r)
	case "DELETE":
		response = handleDELETEStudent(w, r)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)

}

//HandleStudentsMarks is used to store students marks
func HandleStudentsMarks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var response *Response
	switch r.Method {
	case "POST":
		response = handlePOSTMarks(w, r)
	case "DELETE":
		response = handleDELETEMarks(w, r)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)

}

func handleGETAllStudents(w http.ResponseWriter, r *http.Request) (response *Response) {

	vars := mux.Vars(r)
	var events []database.Student
	event := database.Student{
		Entity: vars["module"],
	}
	data, found := caching.GetCache(event.Entity)
	if found {
		return ConstructResponse(&data, "success")
	}
	_, error := database.GetAllEntityFilterOrder(&event, &events, "students.rank asc,students.id")
	for i := range events {
		var marks []database.Marks
		database.GetAllChildEntity(&events[i], &marks, "Marks")
		events[i].Marks = marks
	}
	if *error != nil {
		return ConstructResponse(error, "failed")
	}
	caching.SetCache(event.Entity, events)
	return ConstructResponse(&events, "success")
}

func handlePOSTStudent(w http.ResponseWriter, r *http.Request) (response *Response) {

	var event database.Student
	vars := mux.Vars(r)
	entity := vars["module"]
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseMultipartForm(5 * MB)

	if err != nil {

		return ConstructResponse(&ResponseMessage{
			Status: http.StatusInternalServerError,
			Msg:    "File size found during post should be less than 5 MB",
		}, "custom")
	}

	m := r.MultipartForm
	//get the *fileheaders
	files := m.File["files"]
	if len(files) > 0 {
		path, errorFile := uploadFile(&entity, files[0])
		event = database.Student{
			Path:   path,
			Entity: entity,
		}
		if errorFile != nil {
			return ConstructResponse(&ResponseMessage{
				Status: http.StatusInternalServerError,
				Msg:    "Error in attchement End Point",
			}, "custom")
		}
	}

	rank, _ := strconv.Atoi(r.FormValue("order"))

	event.Name = r.FormValue("name")
	event.Rank = rank
	event.Description = r.FormValue("description")

	_, error := database.PostEntity(&event)

	if *error != nil {
		return ConstructResponse(error, "failed")
	}
	caching.DeleteCache(event.Entity)
	return ConstructResponse(&event, "success")

}

func handlePUTStudent(w http.ResponseWriter, r *http.Request) (response *Response) {
	vars := mux.Vars(r)
	id := vars["id"]
	var oldevent database.Student
	var event database.Student
	rows, error := database.GetEntity(&oldevent, &id)
	if *rows <= 0 || *error != nil {

		return ConstructResponse(&ResponseMessage{
			Status: http.StatusInternalServerError,
			Msg:    "Invalid URL - no data found",
		}, "custom")
	}
	err := r.ParseMultipartForm(20 * MB)
	if err != nil {
		return ConstructResponse(&ResponseMessage{
			Status: http.StatusInternalServerError,
			Msg:    "Huge File size found during post",
		}, "failed")
	}
	m := r.MultipartForm

	entity := r.FormValue("entity")

	//get the *fileheaders
	files := m.File["files"]
	if len(files) > 0 {
		path, errorFile := uploadFile(&entity, files[0])
		deleteFile(&oldevent.Path)
		if errorFile != nil {
			return ConstructResponse(&ResponseMessage{
				Status: http.StatusInternalServerError,
				Msg:    "Huge File size found during post",
			}, "custom")
		}
		event = database.Student{
			Path: path,
		}

	}
	rank, _ := strconv.Atoi(r.FormValue("order"))
	event.Name = r.FormValue("name")
	event.Description = r.FormValue("description")
	event.Entity = r.FormValue("entity")
	event.Rank = rank

	_, errorpost := database.UpdateEntity(&oldevent, &event)

	if *errorpost != nil {

		return ConstructResponse(errorpost, "failed")
	}
	caching.DeleteCache(event.Entity)
	return ConstructResponse(&oldevent, "success")
}
func handleDELETEStudent(w http.ResponseWriter, r *http.Request) (response *Response) {

	vars := mux.Vars(r)
	id := vars["id"]
	var event database.Student
	var images []database.Marks
	rows, error := database.GetEntity(&event, &id)
	if *rows <= 0 || *error != nil {

		return ConstructResponse(&ResponseMessage{
			Status: http.StatusInternalServerError,
			Msg:    "Invalid URL - no data found",
		}, "custom")
	}
	database.GetAllChildEntity(&event, &images, "Marks")
	for _, j := range images {
		database.DeleteEntity(&j)
	}
	_, error = database.DeleteEntity(&event)
	if *error != nil {
		return ConstructResponse(error, "failed")
	}
	deleteFile(&event.Path)
	caching.DeleteCache(event.Entity)
	return ConstructResponse(&event, "success")

}

func handlePOSTMarks(w http.ResponseWriter, r *http.Request) (response *Response) {
	vars := mux.Vars(r)
	parentid := vars["parentid"]
	var student database.Student
	var studentMark database.Marks
	rows, error := database.GetEntity(&student, &parentid)
	if *rows <= 0 || *error != nil {
		return ConstructResponse(&ResponseMessage{
			Status: http.StatusInternalServerError,
			Msg:    "Invalid URL - no data found",
		}, "custom")
	}
	decerror := decodeJSONBody(w, r, &studentMark)
	if decerror != nil {
		return ConstructResponse(decerror, "custom")
	}
	errorpost := database.PostChildEntity(&student, &studentMark)
	if *errorpost != nil {
		return ConstructResponse(error, "failed")
	}
	caching.DeleteCache(student.Entity)
	return ConstructResponse(&studentMark, "success")
}

func handleDELETEMarks(w http.ResponseWriter, r *http.Request) (response *Response) {
	vars := mux.Vars(r)
	id := vars["id"]
	var studentMark database.Marks
	rows, error := database.GetEntity(&studentMark, &id)
	if *rows <= 0 || *error != nil {
		return ConstructResponse(&ResponseMessage{
			Status: http.StatusInternalServerError,
			Msg:    "Invalid URL - no data found",
		}, "custom")
	}
	_, error = database.DeleteEntity(&studentMark)
	if *error != nil {
		return ConstructResponse(error, "failed")
	}
	caching.DeleteCache("10")
	caching.DeleteCache("12")
	return ConstructResponse(&studentMark, "success")

}
