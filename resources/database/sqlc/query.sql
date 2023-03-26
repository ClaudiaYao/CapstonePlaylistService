
-- name: GetRestaurantsByID :one
SELECT * FROM restaurant
WHERE id=$1;

-- name: GetDishByID :one
SELECT * FROM dish
WHERE id=$1;

-- name: GetPlaylistByID :one
SELECT * FROM playlist where id=$1;


-- name: GetPlaylistByPopularity :many
SELECT * FROM playlist where status=Active order by popularity DESC LIMIT 10;

-- name: GetPlaylistByCategory :many
SELECT * FROM playlist where category_code=$1 LIMIT 10;

-- name: GetDishesByPlaylistID :many
SELECT dish.ID, Name, restaurant_id, price, cuisine_style, Ingredient,
Comment FROM dish inner join playlist_dish on playlist_dish.playlist_id=$1;

-- name: GetDishesByRestaurantID :many
SELECT * FROM dish where restaurant_id=$1;

-- name: InsertNewPlaylist :one
Insert into playlist (id, name, category_code,
  dietary_info, status, start_date, end_date,
  popularity) values 
  ($1, $2, $3, $4, $5, $6, $7, $8)
  Returning *;

-- name: InsertNewRestaurant :one
Insert into restaurant (id, name,
  unit_number, address_line1,address_line2,
  postal_code) values 
  ($1, $2, $3, $4, $5, $6)
  Returning *;

-- name: InsertNewCategory :one
Insert into category (code, name, features) values  
  ($1, $2, $3)
  Returning code;

-- name: InsertNewPlaylistDishRelation :one
Insert into playlist_dish (id, dish_id, playlist_id) values 
  ($1, $2, $3)
  Returning *;

-- name: InsertNewDish :one
Insert into dish (id, name, restaurant_id, price,
  cuisine_style, ingredient,
  comment) values 
  ($1, $2, $3, $4, $5, $6, $7)
  Returning *;


