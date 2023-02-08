	docker cp cmd/GenerateData/Generated/. playlist-postgres:/Data

## second, open the psql in the Postgres docker container
	docker exec -it playlist-postgres psql -U postgres

# third, connect to the database and then copy the data files to each table
	\c playlist;
	\COPY restaurant FROM Data/restaurant.txt WITH (FORMAT text, DELIMITER '|');
	\COPY category FROM Data/category.txt WITH (FORMAT text, DELIMITER '|');
	\COPY playlist FROM Data/playlist.txt WITH (FORMAT text, DELIMITER '|');
	\COPY dish FROM Data/dish.txt WITH (FORMAT text, DELIMITER '|');
	\COPY playlist_dish FROM Data/playlist_dish.txt WITH (FORMAT text, DELIMITER '|');