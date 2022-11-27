package models

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/praveenpkg8/dattebayo-go/database"
)

func UpdateContactDetails(perDetId string, brotherContactDetails PersonalDetails) {
	log.Println("Adding brother")
	var err error
	var db *sql.DB = database.GetDB()

	id := uuid.New()
	stmt, err := db.Prepare("UPDATE personalDetails SET phoneNumber = ?, altPhoneNumber = ?, altContactName = ?, altContactRelationship = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(
		brotherContactDetails.PhoneNumber,
		brotherContactDetails.AltPhoneNumber,
		brotherContactDetails.AltContactName,
		brotherContactDetails.AltContactRelationship,
		perDetId,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	log.Printf("Added Contact Details: %v | %v \n", id.String(), perDetId)
}
