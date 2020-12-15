package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	userAPI "github.com/mycontroller-org/backend/v2/pkg/api/user"
	json "github.com/mycontroller-org/backend/v2/pkg/json"
	handlerML "github.com/mycontroller-org/backend/v2/pkg/model/handler"
)

func registerAuthRoutes(router *mux.Router) {
	router.HandleFunc("/api/user/login", login).Methods(http.MethodPost)
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	d, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		postErrorResponse(w, err.Error(), 500)
		return
	}

	userLogin := handlerML.UserLogin{}

	err = json.Unmarshal(d, &userLogin)
	if err != nil {
		postErrorResponse(w, err.Error(), 500)
		return
	}

	// get user details
	userDB, err := userAPI.GetByUsername(userLogin.Username)
	if err != nil {
		postErrorResponse(w, "Invalid user or password!", 401)
		return
	}

	//compare the user from the request, with the one we defined:
	if userLogin.Username != userDB.Username || userLogin.Password != userDB.Password {
		postErrorResponse(w, "Please provide valid login details", 401)
		return
	}
	token, err := createToken(userDB)
	if err != nil {
		postErrorResponse(w, err.Error(), 500)
		return
	}

	// update in cookies
	// expiration := time.Now().Add(7 * 24 * time.Hour)
	// tokenCookie := http.Cookie{Name: "authToken", Value: token, Expires: expiration, Path: "/"}
	// http.SetCookie(w, &tokenCookie)

	tokenResponse := &handlerML.JwtTokenResponse{
		ID:       userDB.ID,
		Username: userDB.Username,
		Email:    userDB.Email,
		FullName: userDB.FullName,
		Token:    token,
	}
	postSuccessResponse(w, tokenResponse)
}
