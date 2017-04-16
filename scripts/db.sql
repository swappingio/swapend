CREATE TABLE "users" (
	id SERIAL PRIMARY KEY,
	username varchar(50) NOT NULL,
	password text NOT NULL,
	email text NOT NULL,
	salt char(40) NOT NULL,
	reset char(100),
	verified boolean NOT NULL
);

CREATE TABLE "songs" (
	id SERIAL PRIMARY KEY,
	songid UUID NOT NULL,
	created timestamp(6)  NOT NULL,
	name varchar(300) NOT NULL,
	uploader INT references users(id),
	versions varchar(36)[]
);

CREATE TABLE "versions" (
	id SERIAL PRIMARY KEY,
	created timestamp(6) NOT NULL,
	fileid UUID NOT NULL,
	uploader INT references users(id)
);