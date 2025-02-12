package db

const InitDB = `
	CREATE type status AS ENUM ('completed', 'incomplete');
	
	CREATE TABLE IF NOT EXISTS tasks
	(
		id         UUID PRIMARY KEY,
		title      varchar(255) NOT NULL,
		status     status       NOT NULL,
		CREATED_AT TIMESTAMP    NOT NULL,
		UPDATED_AT TIMESTAMP    NOT NULL,
		DELETED_AT timestamp
	);
`
