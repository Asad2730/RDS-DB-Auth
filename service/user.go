package service

import (
	"github.com/Asad2730/RDS-DB-Auth/model"
	"github.com/Asad2730/RDS-DB-Auth/repository"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(401, err.Error())
		return
	}
	rs, err := repository.CreateUser(&user)
	if err != nil {
		c.JSON(500, err.Error())
	}
	c.JSON(201, rs)
}

func UpdateUser(c *gin.Context) {

	var user model.User
	email := c.Param("email")

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(401, err.Error())
		return
	}
	rs, err := repository.UpdateUser(email, &user)
	if err != nil {
		c.JSON(500, err.Error())
	}
	c.JSON(200, rs)
}

func GetUsers(c *gin.Context) {

	var user []model.User
	rs, err := repository.GetUsers(user)
	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, rs)
}
