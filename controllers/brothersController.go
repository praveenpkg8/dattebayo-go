package controllers

import (
	"github.com/praveenpkg8/dattebayo-go/models"
)

func CreateBrothers(
	newBrothers models.Brothers,
	brotherPersonalDetails models.PersonalDetails,
) (string, string) {
	brotherId := models.AddBrothers(newBrothers)
	brPerDetId := models.AddPersonalDetails(brotherId, brotherPersonalDetails)
	return brotherId, brPerDetId

}
