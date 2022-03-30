package database

type Table struct {
	Down string
	Up   string
}

var (
	UsersTable = Table{
		Down: `
DROP TABLE IF EXISTS users CASCADE
`,
		Up: `
CREATE TABLE users (
	user_id serial PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	created_on TIMESTAMP NOT NULL,
    last_login TIMESTAMP 
);`,
	}
)

var SeedTables = []Table{
	UsersTable,
}
