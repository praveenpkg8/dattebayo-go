package controllers

import "github.com/praveenpkg8/dattebayo-go/models"

func GetPerIdByBrotherId(brotherId string) (string, error) {
	perId, err := models.GetPerIdByBrotherId(brotherId)
	if err != nil {
		return "", err
	}
	return perId, nil
}

func GetPersonalDetByBrotherID(brotherId string) (models.PersonalDetails, error) {
	return models.GetPersonalDetByBrotherID(brotherId)
}

func UpdatePersonalContact(perDetId string, brPersonalDetail models.PersonalDetails) {
	models.UpdateContactDetails(perDetId, brPersonalDetail)
}

func UpdateFileType(perDetId string, fileSeries string, fileExtension string, contentType string) {
	models.UpdateFileType(perDetId, fileSeries, fileExtension, contentType)
}

func GetFileExtAndContentType(brotherId string, fileSeries string) (string, string, error) {
	return models.GetFileExtAndContentType(brotherId, fileSeries)
}
