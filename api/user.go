package api

import (
	"contacts-server/db"
	"contacts-server/models"
	"github.com/gin-gonic/gin"
	"github.com/imdario/mergo"
	"net/http"
	"strconv"
)

func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := db.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "msg": "Error fetching user from DB"})
	}
	ctx.JSON(http.StatusOK, user)
}

func GetUsers(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")

	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page integer"})
		return
	}

	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page integer"})
		return
	}

	users, err := db.GetUsers(pageInt, limitInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "msg": "Error fetching users from DB"})
	}

	ctx.JSON(http.StatusOK, users)
}

func CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := db.CreateUser(&models.User{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		ContactInfo: user.ContactInfo,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := db.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "msg": "Error fetching users from DB"})
		return
	}

	var updatedUser models.User
	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User Payload"})
		return
	}

	if err := mergo.Merge(&user, updatedUser, mergo.WithOverride); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to merge user data"})
		return
	}

	if err = db.UpdateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err, "msg": "Failed to update User Model"})
	}

	ctx.JSON(http.StatusOK, user)
}

func DeleteUser(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Not Available"})
}
