package api

import (
	"encoding/json"
	"github.com/IceSlanserUni/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//	doLogin login
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the username from the requestBody
	user, err := GetMyUser(r)
	if err != nil {
		ctx.Logger.WithError(err).Error("Failed to parse request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// User Login
	retrievedUser, err := rt.db.LoginUser(user.Username, user.Password)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during func LoginUser")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	err = json.NewEncoder(w).Encode(retrievedUser)
	if err != nil {
		ctx.Logger.WithError(err).Error("Failed to encode user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// doRegister the profile identifier is returned.
func (rt *_router) doRegister(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the username from the requestBody
	user, err := GetMyUser(r)
	if err != nil {
		ctx.Logger.WithError(err).Error("Failed to parse request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if the input username is legal

	// User Sign in
	ID, err := rt.db.RegisterUser(user.Username, user.Password)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during func RegisterUser")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(ID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Failed to encode userID")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetMyUser gets UserName from the request body
func GetMyUser(r *http.Request) (User, error) {
	var temp User
	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
		return User{}, err
	}
	return temp, nil
}
