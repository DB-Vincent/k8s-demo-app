package main

import (
  "net/http"
  "database/sql"
//   "encoding/json"

  "github.com/gin-gonic/gin"

  "github.com/DB-Vincent/k8s-demo-app/db"
)

func main() {
	psql := db.Connect()

	defer psql.Close()

  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  r.GET("/", func(c *gin.Context) {
    quoteList, err := getAllQuotes(psql)
    if err != nil {
      panic(err)
    }

		c.JSON(http.StatusOK, gin.H{
			"quotes": quoteList,
		})
	})

  r.Run()
}

type Quote struct {
	ID int
	Message string
}

func getAllQuotes(db *sql.DB) ([]string, error) {
    rows, err := db.Query("SELECT * FROM QUOTE")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var quotes []string

    for rows.Next() {
        var quote Quote
        if err := rows.Scan(&quote.ID, &quote.Message); err != nil {
          return quotes, err
        }
        quotes = append(quotes, quote.Message)
    }
    if err = rows.Err(); err != nil {
        return quotes, err
    }
    return quotes, nil
}