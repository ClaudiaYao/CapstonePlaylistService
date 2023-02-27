# CapstonePlaylistService

## How to run the system

### 1: Start the Colima by typing "colima start" in command line

### 2. type "make tools" to install the necessary tools, such as Goose and sqlc

> Note: Goose is used to manage sql migration (including the creation of database tables)
> sqlc is used to generate sql queries and database related code automatically.
> The code generated by sqlc is put under resources/database/sqlc/internal
> This app is not using the auto generated sqlc code directly. It is just for reference to save boilerplate code

### 3. type "make up_build" to build the application and run docker-compose.yml to load all the dependencies

> Note: If local application has not any change, just skip this step and go to the next step

### 4. type "make up" to initiate the whole application and dependencies

> Note: If you did the step 3, this step could be skipped

### 5. type "make migrateup" to create database tables and initilize migration

> Note: for this step, if you meet permission or authentication errors, check in terminal window as below steps

- type "sudo lsof -i tcp:5432"
- type "sudo kill -9 <pid-number>"
- Then type "make migrateup" again

### 6. type "make copy_data" to copy sample data to the database.

> Note: here, you need to **_manually copy_** the following lines under the postgres=# command. **I will solve this issue later.**

      \c playlist;

      \COPY restaurant FROM Data/restaurant.txt WITH (FORMAT text, DELIMITER '|');

      \COPY category FROM Data/category.txt WITH (FORMAT text, DELIMITER '|');

      \COPY playlist FROM Data/playlist.txt WITH (FORMAT text, DELIMITER '|');

      \COPY dish FROM Data/dish.txt WITH (FORMAT text, DELIMITER '|');

      \COPY playlist_dish FROM Data/playlist_dish.txt WITH (FORMAT text, DELIMITER '|');

# Test API:

Open postman, and try to test the micro-service

- Get, http://localhost:8080, return a welcome message "welcome to playlist service".
- Get, http://localhost:8080/playlists, return the top 10 most popular playlists.
- Post http://localhost:8080/playlist/new, pass the JSON file as body (Not finish yet)

## Stop the whole APP:

type "make down" to stop the system.

## Generate Sample Data

type "make generate_data".

> All the generated txt files are put under folder: cmd/GenerateData/Generated

## Generate SQL Queries

This task is not done automatically. Just follow these steps:

- Change work directory to resouces/database/sqlc
- Add your query to query.sql
- type "sqlc generate"

> You will see the automatically generated files under resources/database/sqlc/internal. Copy the code snippet to the .go files when you need them.

> You might refer to [sqlc documentation] (https://docs.sqlc.dev/en/stable/) for more information.
