package config

import "fmt"

const (
	host     = "localhost"
	port     = 5432
	user     = "p_admin"
	password = "thenullmind"
	dbname   = "login_credit"
)

func ProvideDBConfig() (string, string) {
	return "postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}
