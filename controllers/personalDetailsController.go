package controllers

import "github.com/praveenpkg8/dattebayo-go/models"

func UpdatePersonalContact(perDetId string, brPersonalDetail models.PersonalDetails) {
	models.UpdateContactDetails(perDetId, brPersonalDetail)
}
