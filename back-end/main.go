package main

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Bharathstark/lourde/back-end/caching"
	"github.com/Bharathstark/lourde/back-end/database"
	"github.com/Bharathstark/lourde/back-end/handler"
	"github.com/Bharathstark/lourde/back-end/logger"
	"github.com/Bharathstark/lourde/back-end/session"
	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			logger.Log.Printf(r.Host + " - " + r.Method + " - " + r.RequestURI + " - " + r.RemoteAddr)
			next.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			sess, _ := session.Store.Get(r, "lourde-cookie")
			if sess.Values["authenticated"] == true && sess.Values["token"] == r.Header.Get("Authorization") {
				next.ServeHTTP(w, r)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(handler.ConstructResponse(handler.ResponseMessage{
					Status: http.StatusUnauthorized,
					Msg:    "Unauthorised Acess..Kindly Relogin",
				}, "custom"))
			}

		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func main() {

	database.Initialize()
	database.InitializeDefaultData()
	caching.Initialize()
	session.Initialize()
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v1/").Subrouter()

	apiRouter.HandleFunc("/notifications", handler.HandleNotifications).Queries("module", "{module}").Methods("GET", "POST")
	apiRouter.HandleFunc("/notifications/{id}", handler.HandleNotifications).Methods("PUT", "DELETE")
	apiRouter.HandleFunc("/events", handler.HandleEvents).Queries("module", "{module}").Methods("GET", "POST")
	apiRouter.HandleFunc("/events/{id}", handler.HandleEvents).Methods("GET", "PUT", "DELETE")
	apiRouter.HandleFunc("/events/{parentid}/images", handler.HandleEventImages).Methods("POST")
	apiRouter.HandleFunc("/events/{parentid}/images/{id}", handler.HandleEventImages).Methods("DELETE")
	apiRouter.HandleFunc("/students", handler.HandleStudents).Queries("module", "{module}").Methods("GET", "POST")
	apiRouter.HandleFunc("/students/{id}", handler.HandleStudents).Methods("PUT", "DELETE")
	apiRouter.HandleFunc("/students/{parentid}/marks", handler.HandleStudentsMarks).Methods("POST")
	apiRouter.HandleFunc("/students/{parentid}/marks/{id}", handler.HandleStudentsMarks).Methods("DELETE")
	apiRouter.HandleFunc("/calendar", handler.HandleCalendar).Methods("GET", "POST")
	apiRouter.HandleFunc("/calendar/{id}", handler.HandleCalendar).Methods("PUT", "DELETE")
	apiRouter.HandleFunc("/change", handler.ChangePassWord).Methods("POST")
	apiRouter.HandleFunc("/health", handler.HandleHealthAPI).Methods("GET")
	apiRouter.Use(authenticationMiddleware)
	authRouter := router.PathPrefix("/api/auth/").Subrouter()
	authRouter.HandleFunc("/login", handler.HandleLogin).Methods("POST")
	authRouter.HandleFunc("/logout", handler.HandleLogout).Methods("GET")
	spa := handler.SPAHandler{StaticPath: "dist", IndexPath: "index.html"}
	apiRouter.Use(loggingMiddleware)
	router.PathPrefix("/").Handler(spa)
	cert, error := tls.LoadX509KeyPair("lourde.crt", "private.key")
	if error != nil {
		logger.Log.Print(error)
	}

	srv := &http.Server{
		Handler:      router,
		Addr:         ":443",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	go http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
	}))

	log.Fatal(srv.ListenAndServeTLS("", ""))

}
