package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/subosito/gotenv"
)

func main()  {
	_ = gotenv.Load()

	g := gin.Default()
	ConnectToDatabase()
	jobs()
	_ = g.Run(":8080")
}
