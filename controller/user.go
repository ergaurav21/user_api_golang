package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"user_api/dao"
)

type UserHandler struct {
	message string
}

func (u UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		getUsers(w)
	case http.MethodPost:
		createUser(w, r)
	case http.MethodPut:
		updateUser(w, r)
	case http.MethodDelete:
		deleteUser(w, r)
	default:
		w.Write([]byte("Method not supported"))

	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var u dao.User
	err := json.NewDecoder(r.Body).Decode(&u) //to parse json and map to
	// the struct
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}
	if u.UserName == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing userid field"))
		return
	}

	user, err := dao.GetUser(*u.UserName)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error occurred"))
		return
	}

	if user != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("duplicate userid"))
		return
	}

	err = dao.InsertUser(u)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error occurred"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func getUsers(w http.ResponseWriter) {

	sl, err := dao.GetUserList()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error occurred"))
		return
	}

	result, err := json.Marshal(sl) // to convert a type to json

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(result)
	w.WriteHeader(http.StatusOK)

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var u dao.User
	err := json.NewDecoder(r.Body).Decode(&u) //to parse json and map to
	// the struct
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}
	if u.UserName == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing username field"))
		return
	}

	user, err := dao.GetUser(*u.UserName)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error occurred"))
		return
	}

	if user == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("username is not present in our records"))
		return
	}

	err = dao.UpdateUser(u)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error occurred"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("User Updated successfully"))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	urlSegments := strings.Split(r.URL.Path, "users/")

	if len(urlSegments[1:]) > 1 || len(urlSegments[1:]) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request is not valid"))
		return
	}

	username := urlSegments[len(urlSegments)-1]

	err := dao.RemoveUser(username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error occurred"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("User Deleted successfully"))

}

func NewUserHandler(message string) (u *UserHandler) {

	return &UserHandler{message: message}
}
