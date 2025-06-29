package persistence

type PostgresConfigRepository interface {
	GetDatabaseName() string
	GetDatabaseHost() string
	GetDatabasePort() string
	GetDatabaseUser() string
	GetDatabasePassword() string
}
