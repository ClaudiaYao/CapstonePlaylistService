package main

import (
	"errors"
	"fmt"
	"net/http"
)

// C: this PlaylistService is responsible for transfering information request/response
// C: the database operation is conducted by its member *sql.DB
// C: when designing API or micro-service, the service request passes data via JSON
func (app *PlaylistService) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	var requestPayload Playlist

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// C: since the requestPayload is struct Playlist format, we need to convert it into
	// C: slice of interface{}, so that the CreatePlaylist could use
	// data := []interface{}{requestPayload.ID, requestPayload.PlaylistName,
	// 	requestPayload.CategoryCode, requestPayload.Price, requestPayload.DietaryInfo,
	// 	requestPayload.Status, requestPayload.StartDate, requestPayload.EndDate,
	// 	requestPayload.Popularity}

	// validate the user against the database
	playlistID, err := app.DBConnection.InsertPlaylist(r.Context(), requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	responsePayload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("playlist is created: %s", playlistID),
	}

	// C: this means the success response
	app.writeJSON(w, http.StatusAccepted, responsePayload)
}

// func GetPlaylistByCrietia(w http.ResponseWriter, r *http.Request) error {

// }

func (app *PlaylistService) Welcome(w http.ResponseWriter, r *http.Request) {
	app.writeJSON(w, http.StatusAccepted, "Welcome to Playlist service!")
}

func (app *PlaylistService) GetPlaylistByID(w http.ResponseWriter, r *http.Request) {
	var playlistID string

	err := app.readJSON(w, r, &playlistID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	// C: since the micro service's request from internal development team, the validity checking
	// C: of the playlistID could be less strict. If the micro service is facing the public,
	// C: more stringent checking should be applied to avoid any malicious query.

	playlist, err := app.DBConnection.GetPlaylistByID(r.Context(), playlistID)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	responsePayload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("playlist is created: %s", playlist.ID),
		Data:    playlist,
	}

	// C: this means the success response
	app.writeJSON(w, http.StatusAccepted, responsePayload)
}

func (app *PlaylistService) GetRestaurantByID(w http.ResponseWriter, r *http.Request) {
	var restaurantID string

	err := app.readJSON(w, r, &restaurantID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	// C: since the micro service's request from internal development team, the validity checking
	// C: of the playlistID could be less strict. If the micro service is facing the public,
	// C: more stringent checking should be applied to avoid any malicious query.

	restaurant, err := app.DBConnection.GetPlaylistByID(r.Context(), restaurantID)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	responsePayload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("restaurant is retrieved: %s", restaurant.ID),
		Data:    restaurant,
	}

	// C: this means the success response
	app.writeJSON(w, http.StatusAccepted, responsePayload)
}

func (app *PlaylistService) GetDishByID(w http.ResponseWriter, r *http.Request) {

	var dishID string
	err := app.readJSON(w, r, &dishID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	// C: since the micro service's request from internal development team, the validity checking
	// C: of the playlistID could be less strict. If the micro service is facing the public,
	// C: more stringent checking should be applied to avoid any malicious query.

	dish, err := app.DBConnection.GetPlaylistByID(r.Context(), dishID)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	responsePayload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("dish is retrieved: %s", dish.ID),
		Data:    dish,
	}

	// C: this means the success response
	app.writeJSON(w, http.StatusAccepted, responsePayload)
}

// func (app *PlaylistService) GetPlaylist(w http.ResponseWriter, r *http.Request) {
// 	var requestPayload struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}

// 	err := app.readJSON(w, r, &requestPayload)
// 	if err != nil {
// 		app.errorJSON(w, err, http.StatusBadRequest)
// 		return
// 	}

// 	// validate the user against the database
// 	user, err := app.Models.User.GetByEmail(requestPayload.Email)
// 	if err != nil {
// 		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
// 		return
// 	}

// 	valid, err := user.PasswordMatches(requestPayload.Password)
// 	if err != nil || !valid {
// 		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
// 		return
// 	}

// 	// log authentication
// 	err = app.logRequest("authentication", fmt.Sprintf("%s logged in", user.Email))
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}

// 	payload := jsonResponse{
// 		Error:   false,
// 		Message: fmt.Sprintf("Logged in user %s", user.Email),
// 		Data:    user,
// 	}

// 	app.writeJSON(w, http.StatusAccepted, payload)
// }

// func (app *Config) logRequest(name, data string) error {
// 	var entry struct {
// 		Name string `json:"name"`
// 		Data string `json:"data"`
// 	}

// 	entry.Name = name
// 	entry.Data = data

// 	jsonData, _ := json.MarshalIndent(entry, "", "\t")
// 	logServiceURL := "http://logger-service/log"

// 	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return err
// 	}

// 	client := &http.Client{}
// 	_, err = client.Do(request)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
