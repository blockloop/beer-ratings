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
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_username ON users(username);
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_uuid ON users(uuid);

INSERT OR IGNORE INTO users (
        uuid, username, email, password_hash, verified, created, modified
)
VALUES (
        'a3edabe3-d4c3-4d70-930a-a29760442852', 'root', 'root@root.com',
        -- password_hash => letmein
        '$2a$10$gEygKg52dEkk1uekIHyz5.zqaRD8skzmi.Ma.9MTxaCPa0KdYANou',
        1, DATETIME(), DATETIME()
);

CREATE TABLE IF NOT EXISTS breweries (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid VARCHAR(36) NOT NULL,
	owner_id INTEGER NOT NULL,
	created_by_user_id INTEGER NOT NULL,
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
		FOREIGN KEY (owner_id) REFERENCES users (id),
		FOREIGN KEY (created_by_user_id) REFERENCES users (id)
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_breweries_uuid ON breweries(uuid);
CREATE INDEX IF NOT EXISTS idx_breweries_name ON breweries(name);
CREATE INDEX IF NOT EXISTS idx_breweries_state ON breweries(state);

CREATE TABLE IF NOT EXISTS beers (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid VARCHAR(36) NOT NULL,
	brewery_id INTEGER NOT NULL,
	name VARCHAR(128) NOT NULL,
        avg_rating FLOAT NOT NULL DEFAULT 0.0,
	description VARCHAR(1024) NOT NULL DEFAULT '',
	created DATETIME NOT NULL,
	modified DATETIME NOT NULL,
		FOREIGN KEY (brewery_id) REFERENCES breweries (id)
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_beers_uuid ON beers(uuid);
CREATE UNIQUE INDEX IF NOT EXISTS idx_beers_name_brewery ON beers(name,brewery_id);

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
CREATE UNIQUE INDEX IF NOT EXISTS idx_ratings_uuid ON ratings(uuid);
CREATE INDEX IF NOT EXISTS idx_ratings_userid ON ratings(user_id);
CREATE INDEX IF NOT EXISTS idx_ratings_beerid ON ratings(beer_id);
