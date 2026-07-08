package project

// Database identifies the database engine a project is configured to use.
type Database string

// Supported database engines.
const (
	DatabaseMariaDB  Database = "mariadb"
	DatabaseMySQL    Database = "mysql"
	DatabasePostgres Database = "postgres"
)

// String returns the human-readable label for d (e.g. "MariaDB").
func (d Database) String() string {
	switch d {
	case DatabaseMariaDB:
		return "MariaDB"
	case DatabaseMySQL:
		return "MySQL"
	case DatabasePostgres:
		return "PostgreSQL"
	default:
		return string(d)
	}
}
