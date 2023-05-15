package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	var (
    host     = getEnv("POSTGRES_HOST", "localhost")
    port     = getEnv("POSTGRES_PORT", "5432")
    user     = getEnv("POSTGRES_USER", "postgres")
    password = getEnv("POSTGRES_PASS", "postgres")
    dbname   = getEnv("POSTGRES_DB_NAME", "demo")
  )
  var err error

//   psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
  psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

  db, err := sql.Open("postgres", psqlconn)
  if err != nil {
    panic(err)
  }

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("[PSQL] Connected!")

  return db
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}