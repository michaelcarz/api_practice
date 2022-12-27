package service

import (
	"golangAPI/objj"
	"log"
	"net/http"

	//"regexp"
	//"strconv"

	"github.com/gin-gonic/gin"
)

//var userList = []objj.User{}

// Get User
func GetUser(c *gin.Context) {
	//c.JSON(http.StatusOK, userList)
	users := objj.FindAllUser()
	c.JSON(http.StatusOK, users)
}

// get user by id
func FindByUserId(c *gin.Context) {
	user := objj.FindByUserId(c.Param("id"))
	if user.UUID == 0 {
		c.JSON(http.StatusNotFound, "error")
		return
	}
	log.Println("User ->", user)
	c.JSON(http.StatusOK, user) //show the json format on postman
}

// Post User
func PostUser(c *gin.Context) {
	user := objj.User{}
	err := c.BindJSON(&user) //bindjson get the parameter from the post requst
	if err != nil {          //Determine whether the json request data structure and the defined structure are equal
		c.JSON(http.StatusNotFound, "error")
		return
	}
	newUser := objj.CreateUser(user)
	c.JSON(http.StatusOK, newUser) //return newUser with Json
}

// delete User
func DeleteUser(c *gin.Context) {
	user := objj.DeleteUser(c.Param("id"))
	if !user { //UUID exists or not
		c.JSON(http.StatusNotFound, "error")
		return
	}
	c.JSON(http.StatusOK, "Successfully") // UUId not zero, then exists
}

// update User
func PutUser(c *gin.Context) {
	user := objj.User{}
	err := c.BindJSON(&user) //the modified data should show in Json format
	if err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}
	user = objj.UpdateUser(c.Param("id"), user) // return id and user which are modified data
	if user.UUID == 0 {
		c.JSON(http.StatusNotFound, "error")
		return
	}
	c.JSON(http.StatusOK, user)
}
