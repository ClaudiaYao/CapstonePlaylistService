
\c playlist;
\COPY restaurant FROM myData/restaurant.txt WITH (FORMAT text, DELIMITER '|');
\COPY category FROM myData/category.txt WITH (FORMAT text, DELIMITER '|');
\COPY playlist FROM myData/playlist.txt WITH (FORMAT text, DELIMITER '|');
\COPY dish FROM myData/dish.txt WITH (FORMAT text, DELIMITER '|');
\COPY playlist_dish FROM myData/playlist_dish.txt WITH (FORMAT text, DELIMITER '|');