CREATE TABLE kids_users (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  email text NOT NULL,
  password text NOT NULL,
  name text,
  gender text,
  family integer
);
CREATE TABLE kids_schools (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  prefecture text,
  city text,
  type text,
  name text NOT NULL
);
CREATE TABLE kids_families (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text
);
CREATE TABLE kids_kids (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text,
  birth date,
  gender text,
  family integer,
  school integer
);
CREATE TABLE kids_task_types (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text,
  family integer
);
CREATE TABLE kids_tasks (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text,
  detail text,
  types integer [],
  status text,
  update date,
  due date,
  items integer [],
  kid integer,
  userId integer,
  family integer
);
CREATE TABLE kids_items (
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text,
  detail text,
  type text,
  image text,
  kid integer,
  family integer
);