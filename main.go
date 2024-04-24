package main

import (
	"product.com/m/database"
	"product.com/m/redisModule"
	"product.com/m/server"
)

func init() {
	database.InitDatabase()
	redisModule.InitRedis()
}

func main() {
	r := server.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	go r.Run(":8080")
	database.MigrateDB()
	select {}
}
