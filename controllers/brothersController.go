package controllers

import (
	"github.com/google/uuid"
	"github.com/praveenpkg8/dattebayo-go/models"
)

func CreateBrothers(newBrothers models.Brothers, brotherPersonalDetails models.PersonalDetails) {
	brotherId := uuid.New().String()
	models.AddBrothers(brotherId, newBrothers)
	models.AddPersonalDetails(brotherId, brotherPersonalDetails)

}
