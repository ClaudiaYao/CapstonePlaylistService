CREATE TABLE "restaurant" (
  "id" varchar PRIMARY KEY,
  "name" varchar(30) NOT NULL,
  "unit_number" varchar(10) NOT NULL,
  "address_line1" varchar(50) NOT NULL,
  "address_line2" varchar(50) NOT NULL,
  "postal_code" int NOT NULL,
  "tag" varchar NOT NULL,
  "operate_hours" varchar NOT NULL,
  "logo_url" varchar NOT NULL,
  "header_url" varchar NOT NULL
);

CREATE TABLE "playlist" (
  "id" varchar PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "category_code" varchar NOT NULL,
  "dietary_info" varchar(100) NOT NULL,
  "status" varchar(100) NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "popularity" int NOT NULL
);


CREATE TABLE "category" (
  "code" varchar PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "features" varchar NOT NULL
);

CREATE TABLE "playlist_dish" (
  "id" varchar PRIMARY KEY, 
  "dish_id" varchar NOT NULL,
  "playlist_id" varchar NOT NULL
);

CREATE TABLE "dish" (
  "id" varchar PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "restaurant_id" varchar NOT NULL,
  "price" float8 NOT NULL,
  "cuisine_style" varchar(100) NOT NULL,
  "ingredient" varchar(100) NOT NULL,
  "dish_options" varchar NOT NULL,
  "comment" varchar(200) NOT NULL,
  "image_url" varchar(200) NOT NUll
);

CREATE INDEX ON "restaurant" ("id");

CREATE INDEX ON "playlist" ("id");

CREATE INDEX ON "category" ("code");

CREATE INDEX ON "playlist_dish" ("dish_id", "playlist_id");

CREATE INDEX ON "dish" ("id");


ALTER TABLE "playlist" ADD FOREIGN KEY ("category_code") REFERENCES "category" ("code");

ALTER TABLE "playlist_dish" ADD FOREIGN KEY ("dish_id") REFERENCES "dish" ("id");

ALTER TABLE "playlist_dish" ADD FOREIGN KEY ("playlist_id") REFERENCES "playlist" ("id");

ALTER TABLE "dish" ADD FOREIGN KEY ("restaurant_id") REFERENCES "restaurant" ("id");