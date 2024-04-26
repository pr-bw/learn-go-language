package database

var connection = "postgresql"

func init() {
	connection = "postgresql"
}

func GetDatabase() string {
	return connection
}
