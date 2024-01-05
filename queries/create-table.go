package queries

const CreateTableQuery = `CREATE TABLE IF NOT EXISTS link (
		id BIGSERIAL NOT NULL PRIMARY KEY,
		long_url VARCHAR(1500) NOT NULL ,
		short_key VARCHAR(7) NOT NULL,
		short_url VARCHAR(30) NOT NULL,
		UNIQUE(short_key)
	)`
