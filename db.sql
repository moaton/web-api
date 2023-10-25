CREATE TABLE users (
	id 				SERIAL 					PRIMARY KEY,
	email 		VARCHAR (255) 	UNIQUE,
	name		 	VARCHAR (100),
  password 	VARCHAR (255),
	amount 		INTEGER,
	deposit 	INTEGER
);

CREATE TABLE revenues (
	id 						SERIAL 				PRIMARY KEY,
	userId 				INTEGER,
	title 				VARCHAR(255),
	description 	TEXT,
	type 					VARCHAR(100),
	amount 				INTEGER,
	createdAt 		TIMESTAMP,
	updatedAt 		TIMESTAMP
);