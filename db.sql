CREATE TABLE "users" {
	id SERIAL PRIMARY KEY,
	username varchar(50) NOT NULL,
	password text NOT NULL,
	hash bytea NOT NULL
-- FIX THIS LATER
}

CREATE TABLE "files" {
	id SERIAL PRIMARY KEY,
	fileid UUID NOT NULL,
	uploader INT references users(id)
}