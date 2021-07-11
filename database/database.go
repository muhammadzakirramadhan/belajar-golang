package database

var connection string

func init() {
	connection = "Mysql"
}

func GetDatabase() string {
	return connection
}
