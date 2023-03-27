
-- name: GetRestaurantsByID :one
SELECT * FROM restaurant
WHERE id=$1;

-- name: GetDishByID :one
SELECT * FROM dish
WHERE id=$1;

-- name: GetPlaylistByID :one
SELECT * FROM playlist where id=$1;


-- name: GetPlaylistsByPopularity :many
SELECT * FROM playlist where status='Active' order by popularity DESC LIMIT 10;

-- name: GetRestanrantsByTag :many
SELECT * FROM restaurant where tag=$1 order by name;

-- name: GetAllRestaurants :many
SELECT * FROM restaurant order by name DESC LIMIT 10;

-- name: GetDishesByRestaurantID :many
SELECT * FROM dish where restaurant_id=$1 order by name;

-- name: GetPlaylistsByCategory :many
SELECT * FROM playlist where category_code=$1 LIMIT 10;

-- name: GetDishesByPlaylistID :many
SELECT dish_id from playlist_dish where playlist_id=$1;


-- name: InsertNewPlaylist :one
Insert into playlist (id, name, category_code,
  dietary_info, status, start_date, end_date,
  popularity) values 
  ($1, $2, $3, $4, $5, $6, $7, $8)
  Returning *;

-- name: InsertNewRestaurant :one
Insert into restaurant (id, name,
  unit_number, address_line1,address_line2,
  postal_code, tag, operate_hours, logo_url, header_url) values 
  ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
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
  comment, dish_options, image_url) values 
  ($1, $2, $3, $4, $5, $6, $7, $8, $9)
  Returning *;


