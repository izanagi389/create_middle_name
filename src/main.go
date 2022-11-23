package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv"
	"izanagi-portfolio-site.com/database"
	"izanagi-portfolio-site.com/server"
)

func main() {
	database.Init()
	server.Init()
}
