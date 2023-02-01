-- name: GetRestaurantByID :one
SELECT * FROM restaurant
WHERE id = $1 LIMIT 1;

-- name: GetMultiRestaurantsByID :many
SELECT  FROM restaurant
WHERE id in ($1);

-- name: GetPlaylistByID :one
SELECT * FROM playlist where id=$1;

-- name: GetPlaylistByCriteria :many
--SELECT * FROM playlist where popularity>$1 and price < $2

-- name: GetDishByID :one
SELECT * FROM dish where id=$1;

-- name: GetMultipleDishesByID :many
SELECT * FROM dish where id in ($1);

-- name: CreatePlaylist :one
Insert into playlist (id, playlist_name, category_code,
  price, dietary_info, status, start_date, end_date,
  popularity) values 
  ($1, $2, $3, $4, $5, $6, $7, $8, $9)
  Returning id;

-- name: InsertNewRestaurant :one
Insert into restaurant (id, name,
  unit_number, address_line1,address_line2,
  postal_code) values 
  ($1, $2, $3, $4, $5, $6)
  Returning id;

-- name: InsertNewCategory :one
Insert into category (code, name, features) values  
  ($1, $2, $3)
  Returning code;

-- name: InsertNewPlaylistDish :one
Insert into playlist_dish (id, dish_id, playlist_id) values 
  ($1, $2, $3)
  Returning id;

-- name: InsertNewDish :one
Insert into dish (id, name, restaurant_id, price,
  cuisine_style, ingredient,
  comment, serve_time) values 
  ($1, $2, $3, $4, $5, $6, $7, $8)
  Returning id;