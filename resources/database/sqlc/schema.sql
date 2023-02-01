CREATE TABLE "restaurant" (
  "id" varchar PRIMARY KEY,
  "name" varchar(30) NOT NULL,
  "unit_number" varchar(10) NOT NULL,
  "address_line1" varchar(50) NOT NULL,
  "address_line2" varchar(50),
  "postal_code" int
);

CREATE TABLE "playlist" (
  "id" varchar PRIMARY KEY,
  "playlist_name" varchar(50) NOT NULL,
  "category_code" varchar NOT NULL,
  "price" float8 NOT NULL,
  "dietary_info" varchar(100),
  "status" varchar(50) NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "popularity" int NOT NULL
);


CREATE TABLE "category" (
  "code" varchar PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "features" varchar(50)
);

CREATE TABLE "playlist_dish" (
  "id" SERIAL PRIMARY KEY, 
  "dish_id" varchar NOT NULL,
  "playlist_id" varchar NOT NULL
);

CREATE TABLE "dish" (
  "id" varchar PRIMARY KEY,
  "name" varchar(20) NOT NULL,
  "restaurant_id" varchar NOT NULL,
  "price" float8 NOT NULL,
  "cuisine_style" varchar(50),
  "ingredient" varchar(50),
  "comment" varchar(200),
  "serve_time" date NOT NULL
);

CREATE INDEX ON "restaurant" ("id");

CREATE INDEX ON "playlist" ("id");

CREATE INDEX ON "address" ("id");

CREATE INDEX ON "category" ("code");

CREATE INDEX ON "playlist_dish" ("dish_id", "playlist_id");

CREATE INDEX ON "dish" ("id");

ALTER TABLE "restaurant" ADD FOREIGN KEY ("address_id") REFERENCES "address" ("id");

ALTER TABLE "playlist" ADD FOREIGN KEY ("category_code") REFERENCES "category" ("code");

ALTER TABLE "playlist_dish" ADD FOREIGN KEY ("dish_id") REFERENCES "dish" ("id");

ALTER TABLE "playlist_dish" ADD FOREIGN KEY ("playlist_id") REFERENCES "playlist" ("id");

ALTER TABLE "dish" ADD FOREIGN KEY ("restaurant_id") REFERENCES "restaurant" ("id");
