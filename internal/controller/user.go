package controller

import (
	"net/http"

	"github.com/buker/go-api-starter/internal/database"
	"github.com/buker/go-api-starter/internal/database/model"
	"github.com/buker/go-api-starter/internal/middleware"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// GetUsers - GET /users
func GetUsers(c *gin.Context) {
	db := database.GetDB()
	users := []model.User{}

	if err := db.Find(&users).Error; err != nil {
		render(c, gin.H{"msg": "not found"}, http.StatusNotFound)
	} else {
		render(c, users, http.StatusOK)
	}
}

// GetUser - GET /users/:id
func GetUser(c *gin.Context) {
	db := database.GetDB()
	id := c.Params.ByName("id")
	user := model.User{}

	if err := db.Where("user_id = ? ", id).First(&user).Error; err != nil {
		render(c, gin.H{"msg": "not found"}, http.StatusNotFound)
	} else {
		render(c, user, http.StatusOK)
	}
}

// CreateUser - POST /users
func CreateUser(c *gin.Context) {
	db := database.GetDB()
	user := model.User{}
	createUser := 0 // default

	user.IDAuth = middleware.AuthID

	if err := db.Where("id_auth = ?", user.IDAuth).First(&user).Error; err == nil {
		createUser = 1
		log.WithError(err).Info()
		render(c, gin.H{"msg": "bad request"}, http.StatusBadRequest)
		return
	}

	if createUser == 0 {
		c.ShouldBindJSON(&user)

		tx := db.Begin()
		if err := tx.Create(&user).Error; err != nil {
			tx.Rollback()
			log.WithError(err).Error()
			render(c, gin.H{"msg": "internal server error"}, http.StatusInternalServerError)
		} else {
			tx.Commit()
			render(c, user, http.StatusCreated)
		}
	}
}

// UpdateUser - PUT /users
func UpdateUser(c *gin.Context) {
	db := database.GetDB()
	user := model.User{}
	updateUser := 0 // default

	user.IDAuth = middleware.AuthID

	if err := db.Where("id_auth = ?", user.IDAuth).First(&user).Error; err != nil {
		updateUser = 1 // user data is not registered, nothing can be updated
		render(c, gin.H{"msg": "not found"}, http.StatusNotFound)
		log.Info()
		return
	}

	if updateUser == 0 {
		c.ShouldBindJSON(&user)

		tx := db.Begin()
		if err := tx.Save(&user).Error; err != nil {
			tx.Rollback()
			log.WithError(err).Error()
			render(c, gin.H{"msg": "internal server error"}, http.StatusInternalServerError)
		} else {
			tx.Commit()
			render(c, user, http.StatusOK)
		}
	}
}
