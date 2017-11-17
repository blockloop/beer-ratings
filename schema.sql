CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid VARCHAR(36) NOT NULL,
	username VARCHAR(128) NOT NULL,
	email VARCHAR(128) NOT NULL,
	password_hash VARCHAR(128) NOT NULL,
	address_line_1 VARCHAR(128) NOT NULL DEFAULT '',
	address_line_2 VARCHAR(128) NOT NULL DEFAULT '',
	city VARCHAR(128) NOT NULL DEFAULT '',
	state VARCHAR(128) NOT NULL DEFAULT '',
	postal_code INTEGER NOT NULL DEFAULT '',
	country VARCHAR(3) NOT NULL DEFAULT '',
	verified BOOLEAN NOT NULL DEFAULT 0,
	created DATETIME NOT NULL,
	modified DATETIME NOT NULL
);
CREATE UNIQUE INDEX idx_users_username ON users(username);
CREATE UNIQUE INDEX idx_users_email ON users(email);
CREATE UNIQUE INDEX idx_users_uuid ON users(uuid);

CREATE TABLE IF NOT EXISTS breweries (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid VARCHAR(36) NOT NULL,
	owner_id INTEGER NOT NULL,
	name VARCHAR(128) NOT NULL,
	address_line_1 VARCHAR(128) NOT NULL DEFAULT '',
	address_line_2 VARCHAR(128) NOT NULL DEFAULT '',
	city VARCHAR(128) NOT NULL DEFAULT '',
	state VARCHAR(128) NOT NULL DEFAULT '',
	postal_code INTEGER NOT NULL DEFAULT '',
	country VARCHAR(3) NOT NULL DEFAULT '',
	verified BOOLEAN NOT NULL DEFAULT 0,
	created DATETIME NOT NULL,
	modified DATETIME NOT NULL,
		FOREIGN KEY (owner_id) REFERENCES users (id)
);
CREATE UNIQUE INDEX idx_breweries_uuid ON breweries(uuid);
CREATE INDEX idx_breweries_name ON breweries(name);
CREATE INDEX idx_breweries_state ON breweries(state);

CREATE TABLE IF NOT EXISTS beers (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid VARCHAR(36) NOT NULL,
	brewery_id INTEGER NOT NULL,
	name VARCHAR(128) NOT NULL,
	description VARCHAR(1024) NOT NULL DEFAULT '',
	created DATETIME NOT NULL,
	modified DATETIME NOT NULL,
		FOREIGN KEY (brewery_id) REFERENCES breweries (id)
);
CREATE UNIQUE INDEX idx_beers_uuid ON beers(uuid);
CREATE UNIQUE INDEX idx_beers_name_brewery ON beers(name,brewery_id);

CREATE TABLE IF NOT EXISTS ratings (
	uuid VARCHAR(36) NOT NULL,
	user_id INTEGER NOT NULL,
	beer_id INTEGER NOT NULL,
	rating TINYINT NOT NULL,
	created DATETIME NOT NULL,
	modified DATETIME NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users (id),
		FOREIGN KEY (beer_id) REFERENCES beers (id)
);
CREATE UNIQUE INDEX idx_ratings_uuid ON ratings(uuid);
CREATE INDEX idx_ratings_userid ON ratings(user_id);
CREATE INDEX idx_ratings_beerid ON ratings(beer_id);
