package main

import (
	connectdb "github.com/Asad2730/RDS-DB-Auth/connectDB"
	"github.com/Asad2730/RDS-DB-Auth/hadler"
	"github.com/gin-gonic/gin"
)

func init() {
	connectdb.DBConnect()
}

func main() {

	r := gin.Default()
	hadler.UserRoutes(r)
	r.Run()

}
