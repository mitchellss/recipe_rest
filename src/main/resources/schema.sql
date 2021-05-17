CREATE SEQUENCE hibernate_sequence START 1;

CREATE TABLE ingredient(
id INTEGER NOT NULL,
ingredient_name varchar(100) NOT NULL,
measurement varchar(100) NOT NULL,
grams varchar(10) NOT NULL,
PRIMARY KEY (id)
);

CREATE TABLE material
(
id INTEGER NOT NULL,
recipe_id INTEGER NOT NULL,
material_number varchar(100) NOT NULL,
measurement varchar(100) NOT NULL,
ingredient_id INTEGER NOT NULL,
note varchar(100) NOT NULL,
PRIMARY KEY (id),
CONSTRAINT testconstraint FOREIGN KEY(recipe_id) REFERENCES recipe(id),
CONSTRAINT testconstraint2 FOREIGN KEY(ingredient_id) REFERENCES ingredient(id)
);

CREATE TABLE recipe
(
id INTEGER NOT NULL,
title varchar(100) NOT NULL,
recipe_datetime varchar(100) NOT NULL,
PRIMARY KEY (id)
);

CREATE TABLE step
(
id INTEGER NOT NULL,
step_number varchar(100) NOT NULL,
step_text varchar(100) NOT NULL,
recipe_id INTEGER NOT NULL,
PRIMARY KEY (id),
CONSTRAINT testconstraint3 FOREIGN KEY(recipe_id) REFERENCES recipe(id)
);
