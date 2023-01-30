package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

// C: Although DataService struct only contains one *sql.DB, using this struct
// C: Could allow to create own service
type DataQuery struct {
	db *sql.DB
}

// var db *sql.DB

func (dq *DataQuery) CreatePlaylist(ctx context.Context, playlistParam Playlist) (string, error) {
	query := `Insert into playlist (id, playlist_name, category_code,
  price, dietary_info, status, start_date, end_date,
  popularity) values 
  ($1, $2, $3, $4, $5, $6, $7, $8, $9)
  Returning id`
	fmt.Println("playlist param:", playlistParam)
	row := dq.db.QueryRowContext(ctx, query, playlistParam.ID, playlistParam.PlaylistName,
		playlistParam.CategoryCode, playlistParam.Price, playlistParam.DietaryInfo,
		playlistParam.Status,
		playlistParam.StartDate, playlistParam.EndDate, playlistParam.Popularity)

	var playlistID string
	err := row.Scan(&playlistID)
	return playlistID, err
}

// func (ds *DataService) GetPlaylistByCrietia(ctx context.Context, criteria map[string]string) ([]Playlist, error) {

// }

func (dq *DataQuery) GetPlaylistByID(ctx context.Context, id string) (*Playlist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// C: could refer to the article of golang website:
	// C: https://go.dev/doc/database/prepared-statements
	query := `SELECT id, playlist_name, category_code, price, dietary_info, 
	status, start_date, end_date, popularity FROM playlist where id=$1`
	row := dq.db.QueryRowContext(ctx, query, id)

	var item Playlist
	if err := row.Scan(
		&item.ID,
		&item.PlaylistName,
		&item.CategoryCode,
		&item.Price,
		&item.DietaryInfo,
		&item.Status,
		&item.StartDate,
		&item.EndDate,
		&item.Popularity,
	); err != nil {
		return nil, err
	}

	return &item, nil
}

// C: get multiple restaurants information based on the ID slices
func (dq *DataQuery) GetMultiRestaurantsByID(ctx context.Context, restaurantIDs []interface{}) (map[string]Restaurant, error) {

	query := `SELECT FROM restaurant WHERE id in ($1)`

	rows, err := dq.db.QueryContext(ctx, query, restaurantIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var restaurants map[string]Restaurant

	for rows.Next() {
		var item Restaurant
		err := rows.Scan(
			&item.ID,
			&item.RestaurantName,
			&item.UnitNumber,
			&item.AddressLine1,
			&item.AddressLine2,
			&item.PostalCode,
		)

		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		restaurants[item.ID] = item

	}

	return restaurants, nil

}

func (dq *DataQuery) GetMultipleDishesByID(ctx context.Context, dishesID []interface{}) (map[string]Dish, error) {
	query := `SELECT id, name, restaurant_id, price, cuisine_style, 
	ingredient, comment, serve_time FROM dish where id in ($1)`

	rows, err := dq.db.QueryContext(ctx, query, dishesID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dishes map[string]Dish

	for rows.Next() {
		var item Dish
		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.RestaurantID,
			&item.Price,
			&item.CuisineStyle,
			&item.Ingredient,
			&item.Comment,
			&item.ServeTime,
		); err != nil {
			return nil, err
		}
		dishes[item.ID] = item
	}

	return dishes, nil
}

// func (ds *DataService) GetRestaurantByID(ctx context.Context, id string) (*Restaurant, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
// 	defer cancel()

// 	query := `select id, email, first_name, last_name, password, user_active, created_at, updated_at from users where email = $1`

// 	rows, err := db.QueryContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var restaurantsInfo map[string]RestaurantInfo

// 	for rows.Next() {
// 		var restaurantinfo RestaurantInfo
// 		err := rows.Scan(
// 			&restaurantinfo.ID,
// 			&restaurantinfo.Name,
// 			&restaurantinfo.AddressLine,
// 			&restaurantinfo.PostalCode,
// 			&restaurantinfo.UnitNumber,
// 		)

// 		if err != nil {
// 			log.Println("Error scanning", err)
// 			return nil, err
// 		}

// 		restaurantsInfo[restaurantinfo.ID] = restaurantinfo

// 	}

// 	return restaurantsInfo, nil

// }
// func (db *sql.DB) GetDishByID(ctx context.Context, id string) (*Dish, error) {

// }

// func (db *sql.DB) UpdatePlaylist(ctx context.Context, playlistID string) (*Playlist, error) {

// }
// func (db *sql.DB) UpdateRestaurant(ctx context.Context, restaurantID string) (*Restaurant, error) {

// }
// func (db *sql.DB) UpdateDish(ctx context.Context) (*Dish, error) {

// }

// func (db *sql.DB) DeletePlaylist(ctx context.Context, id int64) error {

// }
// func (db *sql.DB) DeleteRestaurant(ctx context.Context) error {

// }
// func (db *sql.DB) DeleteDish(ctx context.Context) error {

// }

// New is the function used to create an instance of the data package. It returns the type
// Model, which embeds all the types we want to be available to our application.
// func New(dbPool *sql.DB) playlistService {
// 	db = dbPool
// 	return
// }

// GetAll returns a slice of all playlists, sorted by Popularity
// func (p *sql.DB) GetAll() ([]*Playlist, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
// 	defer cancel()

// 	query := `select id, email, first_name, last_name, password, user_active, created_at, updated_at
// 	from users order by last_name`

// 	rows, err := db.QueryContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var playlists []*Playlist

// 	for rows.Next() {
// 		var playlist Playlist
// 		err := rows.Scan(
// 			&playlist.ID,
// 			&playlist.Name,
// 			&playlist.Price,
// 			&playlist.StartDate,
// 			&playlist.EndDate,
// 			&playlist.Status,
// 		)
// 		if err != nil {
// 			log.Println("Error scanning", err)
// 			return nil, err
// 		}

// 		playlists = append(playlists, &playlist)
// 	}

// 	return playlists, nil
// }

// C: for this micro service, it is responsible for returning the playlist
// C: which status is active
// C: any customization to the existing playlist will be put into another database
// C: subscription
// func (u *Playlist) GetRestaurantInfo(restaurants_ID []*string) (map[string]RestaurantInfo, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
// 	defer cancel()

// 	query := `select id, email, first_name, last_name, password, user_active, created_at, updated_at from users where email = $1`

// 	rows, err := db.QueryContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var restaurantsInfo map[string]RestaurantInfo

// 	for rows.Next() {
// 		var restaurantinfo RestaurantInfo
// 		err := rows.Scan(
// 			&restaurantinfo.ID,
// 			&restaurantinfo.Name,
// 			&restaurantinfo.AddressLine,
// 			&restaurantinfo.PostalCode,
// 			&restaurantinfo.UnitNumber,
// 		)

// 		if err != nil {
// 			log.Println("Error scanning", err)
// 			return nil, err
// 		}

// 		restaurantsInfo[restaurantinfo.ID] = restaurantinfo

// 	}

// 	return restaurantsInfo, nil

// }

// // GetByEmail returns one user by email
// func (u *Playlist) GetPlaylistInfo(playlistID string) (*PlaylistInfo, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
// 	defer cancel()

// 	query := `select id, email, first_name, last_name, password, user_active, created_at, updated_at from users where email = $1`

// 	rows, err := db.QueryContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var dishes map[string]Dish

// 	var restaurantID []string

// 	for rows.Next() {
// 		var dish Dish
// 		err := rows.Scan(
// 			&dish.ID,
// 			&dish.Name,
// 			&dish.RestaurantId,
// 			&dish.Comment,
// 			&dish.CuisineStyle,
// 			&dish.Price,
// 			&dish.ServeTime,
// 			&dish.Ingredient,
// 		)

// 		restaurantID = append(restaurantID, dish.RestaurantId)

// 		if err != nil {
// 			log.Println("Error scanning", err)
// 			return nil, err
// 		}

// 		dishes[dish.ID] = dish

// 	}

// 	return PlaylistInfo, nil

// }

// // GetOne returns one user by id
// func (u *User) GetOne(id int) (*User, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
// 	defer cancel()

// 	query := `select id, email, first_name, last_name, password, user_active, created_at, updated_at from users where id = $1`

// 	var user User
// 	row := db.QueryRowContext(ctx, query, id)

// 	err := row.Scan(
// 		&user.ID,
// 		&user.Email,
// 		&user.FirstName,
// 		&user.LastName,
// 		&user.Password,
// 		&user.Active,
// 		&user.CreatedAt,
// 		&user.UpdatedAt,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }
