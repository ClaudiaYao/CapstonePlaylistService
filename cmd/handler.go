package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func (app *PlaylistService) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	var requestPayload Playlist

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	playlist, err := app.DBConnection.InsertNewPlaylist(r.Context(), requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
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

func (app *PlaylistService) Welcome(w http.ResponseWriter, r *http.Request) {
	app.writeJSON(w, http.StatusAccepted, "Welcome to Playlist service!")
}

func (app *PlaylistService) GetRestaurantByID(w http.ResponseWriter, r *http.Request) {
	var restaurantID string

	err := app.readJSON(w, r, &restaurantID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	restaurant, err := app.DBConnection.GetRestaurantsByID(r.Context(), restaurantID)
	if err != nil {
		app.errorJSON(w, errors.New("wrong query for restaurant table"), http.StatusBadRequest)
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

	dishID := chi.URLParam(r, "dishId")

	dish, err := app.DBConnection.GetDishByID(r.Context(), dishID)
	if err != nil {
		app.errorJSON(w, errors.New("wrong query for the Dish table"), http.StatusBadRequest)
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

func (app *PlaylistService) GetDishesByRestaurant(w http.ResponseWriter, r *http.Request) {

	restaurantID := chi.URLParam(r, "restaurantId")

	responseDTOs, err := app.GetRestaurantInfo(r.Context(), restaurantID)
	if err != nil {
		app.errorJSON(w, errors.New("wrong query for the restaurant table"), http.StatusBadRequest)
		return
	}

	responsePayload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("retrieve all the restaurants info correctly"),
		Data:    *responseDTOs,
	}

	// C: this means the success response
	app.writeJSON(w, http.StatusAccepted, responsePayload)

}

func (app *PlaylistService) GetAllRestaurantsInfo(w http.ResponseWriter, r *http.Request) {

	dto, err := app.GetAllRestaurantInfo(r.Context())
	if err != nil {
		app.errorJSON(w, errors.New("wrong query for the restaurant table"), http.StatusBadRequest)
		return
	}

	responsePayload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("restaurant information is retrieved"),
		Data:    *dto,
	}

	// C: this means the success response
	app.writeJSON(w, http.StatusAccepted, responsePayload)

}

func (app *PlaylistService) wrapGetMultiplePlaylists(w http.ResponseWriter, r *http.Request, playlists *[]Playlist) ([]PlaylistServiceResponseDataDTO, error) {
	responseData := []PlaylistServiceResponseDataDTO{}

	for _, playlist := range *playlists {
		fmt.Println(playlist.ID)
		dishesDTO, restaurantsDTO, err := app.GetPlaylistInfo(r.Context(), playlist.ID)

		if err != nil {
			return nil, errors.New("invalid query for the playlist and other tables,")
		}

		responseDTO := PlaylistServiceResponseDataDTO{
			Playlist:       playlist,
			DishIncluded:   dishesDTO,
			RestaurantInfo: restaurantsDTO,
		}
		responseData = append(responseData, responseDTO)
	}

	return responseData, nil

}

func (app *PlaylistService) GetPopularPlaylists(w http.ResponseWriter, r *http.Request) {

	playlists, err := app.DBConnection.GetPlaylistsByPopularity(r.Context())
	if err != nil {
		app.errorJSON(w, errors.New("invalid query to playlist table"), http.StatusBadRequest)
		return
	}

	responseData, err := app.wrapGetMultiplePlaylists(w, r, &playlists)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	responsePayload := jsonResponse{
		Error:   false,
		Message: "playlists are retrieved",
		Data:    responseData,
	}

	// C: this means the success response
	app.writeJSON(w, http.StatusAccepted, responsePayload)

}

func (app *PlaylistService) GetPlaylistByCategory(w http.ResponseWriter, r *http.Request) {

	// planType := entities.SourceB2C
	// source := strings.TrimSpace(r.URL.Query().Get("source"))
	fmt.Println("enter into category check")

	categoryCode := chi.URLParam(r, "categoryCode")
	fmt.Println("catch category:", categoryCode)

	playlists, err := app.DBConnection.GetPlaylistsByCategory(r.Context(), categoryCode)
	if err != nil {
		app.errorJSON(w, errors.New("invalid query for playlist table"), http.StatusBadRequest)
		return
	}

	responseData, err := app.wrapGetMultiplePlaylists(w, r, &playlists)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	responsePayload := jsonResponse{
		Error:   false,
		Message: "playlists by category are retrieved: " + categoryCode,
		Data:    responseData,
	}

	// C: this means the success response
	app.writeJSON(w, http.StatusAccepted, responsePayload)

}

func (app *PlaylistService) GetPlaylistByID(w http.ResponseWriter, r *http.Request) {

	// planType := entities.SourceB2C
	// source := strings.TrimSpace(r.URL.Query().Get("source"))
	fmt.Println(r.URL)
	playlistID := chi.URLParam(r, "id")
	fmt.Println("catch playlist ID:", playlistID)

	playlist, err := app.DBConnection.GetPlaylistByID(r.Context(), playlistID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	dishes, restaurants, err := app.GetPlaylistInfo(r.Context(), playlistID)

	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	responseDTO := PlaylistServiceResponseDataDTO{
		Playlist:       playlist,
		DishIncluded:   dishes,
		RestaurantInfo: restaurants,
	}

	responsePayload := jsonResponse{
		Error:   false,
		Message: "playlists are retrieved",
		Data:    responseDTO,
	}

	// C: this means the success response
	app.writeJSON(w, http.StatusAccepted, responsePayload)
}

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
