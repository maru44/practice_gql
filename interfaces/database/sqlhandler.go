package database

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"practice_gql/config"
	"time"
)

const slugLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

func ConnectSql() *sql.DB {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s dbname=%s user=%s password=%s sslmode=%s",
			config.PSQL_HOST, config.PSQL_NAME, config.PSQL_USER, config.PSQL_PASS, config.PSQL_SSL,
		),
	)
	if err != nil {
		log.Print(err)
	}
	return db
}

func GenRandSlug(n int) string {
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = slugLetters[rand.Intn(len(slugLetters))]
	}
	return string(b)
}
