CREATE TABLE ingredient
{
id varchar(11) NOT NULL,
name varchar(100) NOT NULL,
measurement varchar(100) NOT NULL,
grams varchar(10) NOT NULL,
PRIMARY KEY (id)
};

CREATE TABLE material
{
id varchar(11) NOT NULL,
recipe_id varchar(100) NOT NULL,
material_number varchar(100) NOT NULL,
measurement varchar(100) NOT NULL,
ingredient varchar(100) NOT NULL,
note varchar(100) NOT NULL,
PRIMARY KEY (id)
};

CREATE TABLE recipe
{
id varchar(11) NOT NULL,
name varchar(100) NOT NULL,
title varchar(100) NOT NULL,
datetime varchar(100) NOT NULL,
PRIMARY KEY (id)
};

CREATE TABLE step
{
id varchar(11) NOT NULL,
step_number varchar(100) NOT NULL,
text varchar(100) NOT NULL,
recipe varchar(100) NOT NULL,
PRIMARY KEY (id)
};
