package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/Bharathstark/lourde/back-end/database"
	"github.com/Bharathstark/lourde/back-end/logger"
	"github.com/Bharathstark/lourde/back-end/session"
	"github.com/dgrijalva/jwt-go"
)

//HandleLogin is used to handle Login Calls
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user database.User
	if r.Header.Get("Content-Type") != "" {
		if !strings.Contains(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
			json.NewEncoder(w).Encode(ConstructResponse(&ResponseMessage{
				Status: http.StatusBadRequest,
				Msg:    "Input Format Error",
			}, "custom"))
			return

		}
	}
	r.ParseMultipartForm(0.5 * MB)
	user.Name = r.FormValue("username")
	user.Hash = r.FormValue("password")
	valid, dbUser := database.IsValidUser(&user)
	if valid {
		sess, _ := session.Store.Get(r, "lourde-cookie")
		if sess.Values["authenticated"] == true {
			logger.Log.Print("User is already Authenticated")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(ConstructResponse(&ResponseMessage{
				Status: http.StatusOK,
				Msg:    "User is already authenticated",
			}, "custom"))
		} else {
			sess.Values["authenticated"] = true
			sessUser := session.User{
				UserID:        dbUser.ID,
				Username:      dbUser.Name,
				Authenticated: true,
			}
			token, _ := CreateToken(dbUser.ID)
			sessUser.Token = token
			sess.Values["userID"] = sessUser.UserID
			sess.Values["token"] = sessUser.Token
			err := sess.Save(r, w)
			if err != nil {
				logger.Log.Print("Error during save session")
				logger.Log.Print(err)
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(ConstructResponse(&ResponseMessage{
					Status: http.StatusBadRequest,
					Msg:    "Error During - Save Session",
				}, "custom"))
				return
			}
			logger.Log.Print("Session Establisted for User :")
			logger.Log.Print(sessUser.UserID)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(ConstructResponse(&sessUser, "sucess"))

		}

	} else {
		logger.Log.Print("User Login - UserName password error")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ConstructResponse(&ResponseMessage{
			Status: http.StatusBadRequest,
			Msg:    "Unauthorised - Access Denied - UserName or Password Error",
		}, "custom"))

	}
}

//HandleLogout is used to handle Login Calls
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Store.Get(r, "lourde-cookie")
	if session.Values["authenticated"] == false {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ConstructResponse(&ResponseMessage{
			Status: http.StatusBadRequest,
			Msg:    "Unauthorised - Access Denied - UserName or Password Error",
		}, "custom"))
	}
	session.Values["authenticated"] = false
	session.Values["user"] = nil
	session.Options.MaxAge = -1
	session.Save(r, w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ConstructResponse(&ResponseMessage{
		Status: http.StatusOK,
		Msg:    "User Logout - Success",
	}, "custom"))
}

//ChangePassWord is used to change passwords
func ChangePassWord(w http.ResponseWriter, r *http.Request) {
	var apiUser database.APIUser
	var user database.User
	if r.Header.Get("Content-Type") != "" {
		if !strings.Contains(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
			json.NewEncoder(w).Encode(ConstructResponse(&ResponseMessage{
				Status: http.StatusBadRequest,
				Msg:    "Input Format Error",
			}, "custom"))
			return

		}
	}
	r.ParseMultipartForm(0.5 * MB)
	apiUser.Name = r.FormValue("username")
	apiUser.Hash = r.FormValue("oldpassword")
	apiUser.NewHash = r.FormValue("newpassword")
	user.Name = apiUser.Name
	user.Hash = apiUser.Hash
	valid, dbUser := database.IsValidUser(&user)
	if valid {
		logger.Log.Print("Password Updated for User - " + dbUser.Name)
		user.Hash, _ = database.HashUserPassword(apiUser.NewHash)
		database.UpdateEntity(&dbUser, &user)
		HandleLogout(w, r)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ConstructResponse(&ResponseMessage{
			Status: http.StatusBadRequest,
			Msg:    "Incorrect User Name Password - cannot change password",
		}, "custom"))
	}

}

//CreateToken is used to create Access Tokens
func CreateToken(userid uint) (string, error) {
	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("mysecretkeyironman"))
	if err != nil {
		return "", err
	}
	return token, nil
}
