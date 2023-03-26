package main

import (
	"context"
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 3

// C: Although DataService struct only contains one *sql.DB, using this struct
// C: Could allow to create own service
type DataQuery struct {
	db *sql.DB
}

const getDishByID = `
SELECT id, name, restaurant_id, price, cuisine_style, ingredient, comment FROM dish
WHERE id=$1
`

func (q *DataQuery) GetDishByID(ctx context.Context, id string) (Dish, error) {
	row := q.db.QueryRowContext(ctx, getDishByID, id)
	var i Dish
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.RestaurantID,
		&i.Price,
		&i.CuisineStyle,
		&i.Ingredient,
		&i.Comment,
	)
	return i, err
}

const getDishesByPlaylistID = `
SELECT dish.ID, Name, restaurant_id, price, cuisine_style, Ingredient,
Comment FROM dish inner join playlist_dish on playlist_dish.playlist_id=$1
`

func (q *DataQuery) GetDishesByPlaylistID(ctx context.Context, playlistID string) ([]Dish, error) {
	rows, err := q.db.QueryContext(ctx, getDishesByPlaylistID, playlistID)
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

const getDishesByRestaurantID = `
SELECT id, name, restaurant_id, price, cuisine_style, ingredient, comment FROM dish where restaurant_id=$1
`

func (q *DataQuery) GetDishesByRestaurantID(ctx context.Context, restaurantID string) ([]Dish, error) {
	rows, err := q.db.QueryContext(ctx, getDishesByRestaurantID, restaurantID)
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

const getPlaylistByCategory = `
SELECT id, name, category_code, dietary_info, status, start_date, end_date, popularity FROM playlist where category_code=$1 LIMIT 10
`

func (q *DataQuery) GetPlaylistByCategory(ctx context.Context, categoryCode string) ([]Playlist, error) {
	rows, err := q.db.QueryContext(ctx, getPlaylistByCategory, categoryCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Playlist
	for rows.Next() {
		var i Playlist
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CategoryCode,
			&i.DietaryInfo,
			&i.Status,
			&i.StartDate,
			&i.EndDate,
			&i.Popularity,
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

const getPlaylistByID = `
SELECT id, name, category_code, dietary_info, status, start_date, end_date, popularity FROM playlist where id=$1
`

func (q *DataQuery) GetPlaylistByID(ctx context.Context, id string) (Playlist, error) {
	row := q.db.QueryRowContext(ctx, getPlaylistByID, id)
	var i Playlist
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CategoryCode,
		&i.DietaryInfo,
		&i.Status,
		&i.StartDate,
		&i.EndDate,
		&i.Popularity,
	)
	return i, err
}

const getPlaylistByPopularity = `
SELECT id, name, category_code, dietary_info, status, start_date, end_date, popularity FROM playlist where status='Active' order by popularity DESC LIMIT 10
`

func (q *DataQuery) GetPlaylistByPopularity(ctx context.Context) ([]Playlist, error) {
	rows, err := q.db.QueryContext(ctx, getPlaylistByPopularity)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Playlist
	for rows.Next() {
		var i Playlist
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CategoryCode,
			&i.DietaryInfo,
			&i.Status,
			&i.StartDate,
			&i.EndDate,
			&i.Popularity,
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

const getRestaurantsByID = `
SELECT id, name, unit_number, address_line1, address_line2, postal_code FROM restaurant
WHERE id=$1
`

func (q *DataQuery) GetRestaurantsByID(ctx context.Context, id string) (Restaurant, error) {
	row := q.db.QueryRowContext(ctx, getRestaurantsByID, id)
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

const insertNewCategory = `
Insert into category (code, name, features) values  
  ($1, $2, $3)
  Returning code
`

type InsertNewCategoryParams struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Features string `json:"features"`
}

func (q *DataQuery) InsertNewCategory(ctx context.Context, arg InsertNewCategoryParams) (string, error) {
	row := q.db.QueryRowContext(ctx, insertNewCategory, arg.Code, arg.Name, arg.Features)
	var code string
	err := row.Scan(&code)
	return code, err
}

const insertNewDish = `
Insert into dish (id, name, restaurant_id, price,
  cuisine_style, ingredient,
  comment) values 
  ($1, $2, $3, $4, $5, $6, $7)
  Returning id, name, restaurant_id, price, cuisine_style, ingredient, comment
`

func (q *DataQuery) InsertNewDish(ctx context.Context, arg Dish) (Dish, error) {
	row := q.db.QueryRowContext(ctx, insertNewDish,
		arg.ID,
		arg.Name,
		arg.RestaurantID,
		arg.Price,
		arg.CuisineStyle,
		arg.Ingredient,
		arg.Comment,
	)
	var i Dish
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.RestaurantID,
		&i.Price,
		&i.CuisineStyle,
		&i.Ingredient,
		&i.Comment,
	)
	return i, err
}

const insertNewPlaylist = `
Insert into playlist (id, name, category_code,
  dietary_info, status, start_date, end_date,
  popularity) values 
  ($1, $2, $3, $4, $5, $6, $7, $8)
  Returning id, name, category_code, dietary_info, status, start_date, end_date, popularity
`

func (q *DataQuery) InsertNewPlaylist(ctx context.Context, arg Playlist) (Playlist, error) {
	row := q.db.QueryRowContext(ctx, insertNewPlaylist,
		arg.ID,
		arg.Name,
		arg.CategoryCode,
		arg.DietaryInfo,
		arg.Status,
		arg.StartDate,
		arg.EndDate,
		arg.Popularity,
	)
	var i Playlist
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CategoryCode,
		&i.DietaryInfo,
		&i.Status,
		&i.StartDate,
		&i.EndDate,
		&i.Popularity,
	)
	return i, err
}

const insertNewPlaylistDishRelation = `
Insert into playlist_dish (id, dish_id, playlist_id) values 
  ($1, $2, $3)
  Returning id, dish_id, playlist_id
`

type InsertNewPlaylistDishRelationParams struct {
	ID         string `json:"id"`
	DishID     string `json:"dishID"`
	PlaylistID string `json:"playlistID"`
}

func (q *DataQuery) InsertNewPlaylistDishRelation(ctx context.Context, arg PlaylistDish) (PlaylistDish, error) {
	row := q.db.QueryRowContext(ctx, insertNewPlaylistDishRelation, arg.ID, arg.DishID, arg.PlaylistID)
	var i PlaylistDish
	err := row.Scan(&i.ID, &i.DishID, &i.PlaylistID)
	return i, err
}

const insertNewRestaurant = `
Insert into restaurant (id, name,
  unit_number, address_line1,address_line2,
  postal_code) values 
  ($1, $2, $3, $4, $5, $6)
  Returning id, name, unit_number, address_line1, address_line2, postal_code
`

func (q *DataQuery) InsertNewRestaurant(ctx context.Context, arg Restaurant) (Restaurant, error) {
	row := q.db.QueryRowContext(ctx, insertNewRestaurant,
		arg.ID,
		arg.Name,
		arg.UnitNumber,
		arg.AddressLine1,
		arg.AddressLine2,
		arg.PostalCode,
	)
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
