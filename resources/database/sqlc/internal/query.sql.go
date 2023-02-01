// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createPlaylist = `-- name: CreatePlaylist :one
Insert into playlist (id, playlist_name, category_code,
  price, dietary_info, status, start_date, end_date,
  popularity) values 
  ($1, $2, $3, $4, $5, $6, $7, $8, $9)
  Returning id
`

type CreatePlaylistParams struct {
	ID           string         `json:"id"`
	PlaylistName string         `json:"playlistName"`
	CategoryCode string         `json:"categoryCode"`
	Price        float64        `json:"price"`
	DietaryInfo  sql.NullString `json:"dietaryInfo"`
	Status       string         `json:"status"`
	StartDate    time.Time      `json:"startDate"`
	EndDate      time.Time      `json:"endDate"`
	Popularity   int32          `json:"popularity"`
}

func (q *Queries) CreatePlaylist(ctx context.Context, arg CreatePlaylistParams) (string, error) {
	row := q.db.QueryRowContext(ctx, createPlaylist,
		arg.ID,
		arg.PlaylistName,
		arg.CategoryCode,
		arg.Price,
		arg.DietaryInfo,
		arg.Status,
		arg.StartDate,
		arg.EndDate,
		arg.Popularity,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const getMultiRestaurantsByID = `-- name: GetMultiRestaurantsByID :many
SELECT  FROM restaurant
WHERE id in ($1)
`

type GetMultiRestaurantsByIDRow struct {
}

func (q *Queries) GetMultiRestaurantsByID(ctx context.Context, id string) ([]GetMultiRestaurantsByIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getMultiRestaurantsByID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMultiRestaurantsByIDRow
	for rows.Next() {
		var i GetMultiRestaurantsByIDRow
		if err := rows.Scan(); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMultipleDishesByID = `-- name: GetMultipleDishesByID :many
SELECT id, name, restaurant_id, price, cuisine_style, ingredient, comment, serve_time FROM dish where id in ($1)
`

func (q *Queries) GetMultipleDishesByID(ctx context.Context, id string) ([]Dish, error) {
	rows, err := q.db.QueryContext(ctx, getMultipleDishesByID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Dish
	for rows.Next() {
		var i Dish
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.RestaurantID,
			&i.Price,
			&i.CuisineStyle,
			&i.Ingredient,
			&i.Comment,
			&i.ServeTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPlaylistByCriteria = `-- name: GetPlaylistByCriteria :many

SELECT id, name, restaurant_id, price, cuisine_style, ingredient, comment, serve_time FROM dish where id=$1
`

// SELECT * FROM playlist where popularity>$1 and price < $2
func (q *Queries) GetPlaylistByCriteria(ctx context.Context, id string) ([]Dish, error) {
	rows, err := q.db.QueryContext(ctx, getPlaylistByCriteria, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Dish
	for rows.Next() {
		var i Dish
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.RestaurantID,
			&i.Price,
			&i.CuisineStyle,
			&i.Ingredient,
			&i.Comment,
			&i.ServeTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPlaylistByID = `-- name: GetPlaylistByID :one
SELECT id, playlist_name, category_code, price, dietary_info, status, start_date, end_date, popularity FROM playlist where id=$1
`

func (q *Queries) GetPlaylistByID(ctx context.Context, id string) (Playlist, error) {
	row := q.db.QueryRowContext(ctx, getPlaylistByID, id)
	var i Playlist
	err := row.Scan(
		&i.ID,
		&i.PlaylistName,
		&i.CategoryCode,
		&i.Price,
		&i.DietaryInfo,
		&i.Status,
		&i.StartDate,
		&i.EndDate,
		&i.Popularity,
	)
	return i, err
}

const getRestaurantByID = `-- name: GetRestaurantByID :one
SELECT id, name, unit_number, address_line1, address_line2, postal_code FROM restaurant
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRestaurantByID(ctx context.Context, id string) (Restaurant, error) {
	row := q.db.QueryRowContext(ctx, getRestaurantByID, id)
	var i Restaurant
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UnitNumber,
		&i.AddressLine1,
		&i.AddressLine2,
		&i.PostalCode,
	)
	return i, err
}

const insertNewCategory = `-- name: InsertNewCategory :one
Insert into category (code, name, features) values  
  ($1, $2, $3)
  Returning code
`

type InsertNewCategoryParams struct {
	Code     string         `json:"code"`
	Name     string         `json:"name"`
	Features sql.NullString `json:"features"`
}

func (q *Queries) InsertNewCategory(ctx context.Context, arg InsertNewCategoryParams) (string, error) {
	row := q.db.QueryRowContext(ctx, insertNewCategory, arg.Code, arg.Name, arg.Features)
	var code string
	err := row.Scan(&code)
	return code, err
}

const insertNewDish = `-- name: InsertNewDish :one
Insert into dish (id, name, restaurant_id, price,
  cuisine_style, ingredient,
  comment, serve_time) values 
  ($1, $2, $3, $4, $5, $6, $7, $8)
  Returning id
`

type InsertNewDishParams struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	RestaurantID string         `json:"restaurantID"`
	Price        float64        `json:"price"`
	CuisineStyle sql.NullString `json:"cuisineStyle"`
	Ingredient   sql.NullString `json:"ingredient"`
	Comment      sql.NullString `json:"comment"`
	ServeTime    time.Time      `json:"serveTime"`
}

func (q *Queries) InsertNewDish(ctx context.Context, arg InsertNewDishParams) (string, error) {
	row := q.db.QueryRowContext(ctx, insertNewDish,
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

const insertNewPlaylistDish = `-- name: InsertNewPlaylistDish :one
Insert into playlist_dish (id, dish_id, playlist_id) values 
  ($1, $2, $3)
  Returning id
`

type InsertNewPlaylistDishParams struct {
	ID         int32  `json:"id"`
	DishID     string `json:"dishID"`
	PlaylistID string `json:"playlistID"`
}

func (q *Queries) InsertNewPlaylistDish(ctx context.Context, arg InsertNewPlaylistDishParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertNewPlaylistDish, arg.ID, arg.DishID, arg.PlaylistID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const insertNewRestaurant = `-- name: InsertNewRestaurant :one
Insert into restaurant (id, name,
  unit_number, address_line1,address_line2,
  postal_code) values 
  ($1, $2, $3, $4, $5, $6)
  Returning id
`

type InsertNewRestaurantParams struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	UnitNumber   string         `json:"unitNumber"`
	AddressLine1 string         `json:"addressLine1"`
	AddressLine2 sql.NullString `json:"addressLine2"`
	PostalCode   sql.NullInt32  `json:"postalCode"`
}

func (q *Queries) InsertNewRestaurant(ctx context.Context, arg InsertNewRestaurantParams) (string, error) {
	row := q.db.QueryRowContext(ctx, insertNewRestaurant,
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
