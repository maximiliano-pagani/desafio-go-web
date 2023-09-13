package main

import (
	"os"

	"github.com/bootcamp-go/desafio-go-web/cmd/server/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("./.env")
	ticketsCsvPath := os.Getenv("TICKETS_CSV_PATH")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	router := router.NewRouterTicket(r, ticketsCsvPath)
	router.Run()
}
