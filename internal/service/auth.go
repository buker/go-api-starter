package service

import (
	"github.com/buker/go-api-starter/internal/database"
	"github.com/buker/go-api-starter/internal/database/model"
)

// GetUserByEmail ...
func GetUserByEmail(email string) (*model.Auth, error) {
	db := database.GetDB()

	var auth model.Auth

	if err := db.Where("email = ? ", email).Find(&auth).Error; err != nil {
		return nil, err
	}

	return &auth, nil
}
