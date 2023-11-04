package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-bookstore-users-api/domain"
	"github.com/go-bookstore-users-api/service"
	"github.com/go-bookstore-users-api/utils"
)

func GetUser(c *gin.Context) {
	// Validate user_id
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		errResponse := utils.BadRequest("Invalid user id.")
		c.JSON(errResponse.Status, errResponse)
		return
	}

	usr, errService := service.GetUser(userId)
	if errService != nil {
		c.JSON(errService.Status, errService)
		return
	}

	c.JSON(http.StatusOK, usr)
	return
}

func CreateUser(c *gin.Context) {
	var user domain.User

	// Validate Json
	err := c.ShouldBindJSON(&user)
	if err != nil {
		errResponse := utils.BadRequest("invalid json request aca.")
		c.JSON(errResponse.Status, errResponse)
		return
	}

	// Create user
	userCreated, errService := service.CreateUser(user)
	if errService != nil {
		c.JSON(errService.Status, errService)
		return
	}

	c.JSON(http.StatusCreated, userCreated)
}

func UpdateUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest("Invalid user_id"))
		return
	}

	var user domain.User
	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest("Invalid json request"))
		return
	}

	user.Id = userId
	result, srvErr := service.UpdateUser(user)
	if srvErr != nil {
		c.JSON(srvErr.Status, srvErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func SearchUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Endpoint not implemented")
}

func DeleteUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest("Invalid user id"))
		return
	}

	if err := service.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
