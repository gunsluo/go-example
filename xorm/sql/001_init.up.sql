
CREATE TABLE
IF
	NOT EXISTS inspector (
		ID SERIAL PRIMARY KEY NOT NULL,
		username VARCHAR ( 256 ) NOT NULL DEFAULT '',
		password VARCHAR ( 256 ) NOT NULL DEFAULT '', 
		created TIMESTAMP NOT NULL DEFAULT NOW()
	);

CREATE TABLE
IF
	NOT EXISTS credential (
		ID SERIAL PRIMARY KEY NOT NULL,
		phone_number VARCHAR ( 32 ) NOT NULL UNIQUE DEFAULT '',
		verified bool NOT NULL DEFAULT FALSE,
		changed_pwd bool NOT NULL DEFAULT FALSE,
		PASSWORD VARCHAR ( 256 ) NOT NULL DEFAULT '',
		created TIMESTAMP NOT NULL DEFAULT NOW()
	);
CREATE UNIQUE INDEX credential_phone_number_uindex ON credential ( phone_number );
