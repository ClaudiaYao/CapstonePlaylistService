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
			&item.Name,
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

// The following are insertion operation. They might not be used in the frontend
// Quest, but will be used in creating mock data to display the functionality

func (dq *DataQuery) InsertPlaylist(ctx context.Context, playlistParam Playlist) (string, error) {
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

func (dq *DataQuery) InsertNewCategory(ctx context.Context, arg Category) (string, error) {

	query := "Insert into category (code, name, features) values ($1, $2, $3)"
	fmt.Println("insert new category:", arg)

	row := dq.db.QueryRowContext(ctx, query, arg.Code, arg.Name, arg.Features)
	var code string
	err := row.Scan(&code)
	return code, err
}

func (dq *DataQuery) InsertNewDish(ctx context.Context, arg Dish) (string, error) {
	query := `Insert into dish (id, name, restaurant_id, price,
		cuisine_style, ingredient,
		comment, serve_time) values 
		($1, $2, $3, $4, $5, $6, $7, $8)`

	row := dq.db.QueryRowContext(ctx, query,
		arg.ID,
		arg.Name,
		arg.RestaurantID,
		arg.Price,
		arg.CuisineStyle,
		arg.Ingredient,
		arg.Comment,
		arg.ServeTime,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

func (dq *DataQuery) InsertNewPlaylistDish(ctx context.Context, arg PlaylistDish) (int32, error) {

	query := `Insert into playlist_dish (id, dish_id, playlist_id) values 
	($1, $2, $3)`
	row := dq.db.QueryRowContext(ctx, query, arg.ID, arg.DishID, arg.PlaylistID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

func (dq *DataQuery) InsertNewRestaurant(ctx context.Context, arg Restaurant) (string, error) {

	query := `Insert into restaurant (id, name,
		unit_number, address_line1,address_line2,
		postal_code) values 
		($1, $2, $3, $4, $5, $6)`
	row := dq.db.QueryRowContext(ctx, query,
		arg.ID,
		arg.Name,
		arg.UnitNumber,
		arg.AddressLine1,
		arg.AddressLine2,
		arg.PostalCode,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}
