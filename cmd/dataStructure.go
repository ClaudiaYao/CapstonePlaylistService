package main

import (
	"errors"
	"time"
)

type Category struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Features string `json:"features,omitempty"`
}

type Dish struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	RestaurantID string    `json:"restaurantID"`
	Price        float64   `json:"price"`
	CuisineStyle string    `json:"cuisineStyle,omitempty"`
	Ingredient   string    `json:"ingredient,omitempty"`
	Comment      string    `json:"comment,omitempty"`
	ServeTime    time.Time `json:"serveTime,omitempty"`
	DishOptions  string    `json:"dishOptions"`
	ImageUrl     string    `json:"imageUrl"`
}

type Playlist struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	CategoryCode string    `json:"categoryCode"`
	DietaryInfo  string    `json:"dietaryInfo"`
	Status       string    `json:"status"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	Popularity   int32     `json:"popularity"`
}

type PlaylistDish struct {
	ID         int32  `json:"id"`
	DishID     string `json:"dishID"`
	PlaylistID string `json:"playlistID"`
}

type Restaurant struct {
	ID           string `json:"id"`
	Name         string `json:"restaurantName"`
	UnitNumber   string `json:"unitNumber"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	PostalCode   int    `json:"postalCode"`
	OperateHours string `json:"operateHours"`
	Tag          string `json:"tag"`
	LogoUrl      string `json:"logoURL"`
	HeaderUrl    string `json:"headerURL"`
}

type PlaylistServiceResponseDataDTO struct {
	Playlist       Playlist
	DishIncluded   []Dish
	RestaurantInfo []Restaurant
}

type RestaurantResponseDataDTO struct {
	RestaurantInfo Restaurant
	DishIncluded   []Dish
}

// This struct includes all the data returned to the request
// DishIncluded is a map structure, the key is the DishID
// RestaurantInfo is a map structure, the key is the RestaurantID
// RestaurantAddress is a map structure, the key is the AddressID

type PlaylistService struct {
	DBConnection *DataQuery
	Model        *PlaylistServiceData
}

type PlaylistServiceData struct {
	Playlist       []Playlist
	DishIncluded   map[string]Dish
	RestaurantInfo map[string]Restaurant
}

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExist     = errors.New("row does not exist")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

// // C: For the normal users, the request allows retrieving the playlist information
// // C: For the backend admin users, the request allows to update the playlist
// // C: and restaurant, dish information
// // define the interface functions
// type playlistServiceInterface interface {
// 	CreatePlaylist(ctx context.Context) (*Playlist, error)
// 	GetPlaylistByCrietia(ctx context.Context, criteria ...string) ([]Playlist, error)

// 	GetPlaylistByID(ctx context.Context, name string) (*Playlist, error)
// 	GetRestaurantByID(ctx context.Context, id string) (*Restaurant, error)
// 	GetDishByID(ctx context.Context, id string) (*Dish, error)

// 	UpdatePlaylist(ctx context.Context, playlistID string) (*Playlist, error)
// 	UpdateRestaurant(ctx context.Context, restaurantID string) (*Restaurant, error)
// 	UpdateDish(ctx context.Context) (*Dish, error)

// 	DeletePlaylist(ctx context.Context, id int64) error
// 	DeleteRestaurant(ctx context.Context) error
// 	DeleteDish(ctx context.Context) error
// }
