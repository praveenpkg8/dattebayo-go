package scripts

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/praveenpkg8/dattebayo-go/controllers"
	"github.com/praveenpkg8/dattebayo-go/models"
)

func PopulateDB() {
	fileName := "./scripts/mockData.json"
	plan, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(err)
	}
	var brothersData models.Brothers
	var personalDetailsData models.PersonalDetails

	if err := json.Unmarshal(plan, &brothersData); err != nil {
		log.Println(err)
	}
	if err := json.Unmarshal(plan, &personalDetailsData); err != nil {
		log.Println(err)
	}

	brotherId, perDetId := controllers.CreateBrothers(brothersData, personalDetailsData)
	brothersData.Id = brotherId
	personalDetailsData.Id = perDetId
	controllers.UpdatePersonalContact(perDetId, personalDetailsData)

}
