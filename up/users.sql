--CREATE USERS--
--Do not rename the table or modify the existing columns unless you know what you're doing.
CREATE TABLE users (
    id serial primary KEY,
    email varchar(50) not null,
    name varchar(25) not null,
    surname varchar(25) not null,
    password text not null
)