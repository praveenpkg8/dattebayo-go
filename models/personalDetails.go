package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/praveenpkg8/dattebayo-go/database"
)

func GetPerIdByBrotherId(brotherId string) (string, error) {
	var db *sql.DB = database.GetDB()

	query := "SELECT id FROM personalDetails WHERE brotherId = ?"

	rows, err := db.Query(query, brotherId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		var id string

		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}

		// Print or use the retrieved data
		fmt.Printf("ID: %s \n", id)
		return id, nil
	} else {
		return "", errors.New("no results found")
	}
}

func GetPersonalDetByBrotherID(brotherId string) (PersonalDetails, error) {
	var db *sql.DB = database.GetDB()
	var CurrentAddress string
	var PermanentAddress string
	query := "SELECT * FROM personalDetails WHERE brotherId = ?"

	rows, err := db.Query(query, brotherId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		var personalDetails PersonalDetails

		if err := rows.Scan(
			&personalDetails.Id,
			&personalDetails.BrotherId,
			&personalDetails.DoB,
			&personalDetails.Email,
			&personalDetails.PhoneNumber,
			&personalDetails.AltPhoneNumber,
			&personalDetails.AltContactName,
			&personalDetails.AltContactRelationship,
			&CurrentAddress,
			&personalDetails.CurrentFileType,
			&personalDetails.CurrentContentType,
			&PermanentAddress,
			&personalDetails.PermanentFileType,
			&personalDetails.PermanentContentType,
		); err != nil {
			log.Fatal(err)
		}
		json.Unmarshal([]byte(CurrentAddress), &personalDetails.CurrentAddress)
		json.Unmarshal([]byte(PermanentAddress), &personalDetails.PermanentAddress)

		fmt.Printf("ID: %s \n", brotherId)
		return personalDetails, nil
	} else {
		return PersonalDetails{}, errors.New("no results found")
	}
}

func UpdateContactDetails(perDetId string, brotherContactDetails PersonalDetails) {
	log.Printf("Adding update brothers contact")
	var err error
	var db *sql.DB = database.GetDB()

	CurrentAddressString, _ := json.Marshal(brotherContactDetails.CurrentAddress)
	PermanentAddressString, _ := json.Marshal(brotherContactDetails.PermanentAddress)

	stmt, err := db.Prepare(
		"UPDATE personalDetails SET " +
			"dob = ?, email = ?," +
			"phoneNumber = ?, altPhoneNumber = ?, " +
			"altContactName = ?, altContactRelationship = ?," +
			"currentAddress = ?, permanentAddress = ? " +
			"WHERE id = ?",
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(
		brotherContactDetails.DoB,
		brotherContactDetails.Email,
		brotherContactDetails.PhoneNumber,
		brotherContactDetails.AltPhoneNumber,
		brotherContactDetails.AltContactName,
		brotherContactDetails.AltContactRelationship,
		CurrentAddressString,
		PermanentAddressString,
		perDetId,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	log.Printf("Added Contact Details: %v \n", perDetId)
}

func UpdateFileType(perDetId string, fileSeries string, fileExtension string, contentType string) {
	log.Printf("Updating file Extension %+v %+v", fileSeries, fileExtension)
	var db *sql.DB = database.GetDB()

	updateSqlStatement := fmt.Sprintf("UPDATE personalDetails SET %sFileType = ?, %sContentType = ? WHERE id = ?", fileSeries, fileSeries)
	stmt, err := db.Prepare(updateSqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(
		fileExtension,
		contentType,
		perDetId,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	log.Printf("Added Address Details: %v \n", perDetId)
}

func GetFileExtAndContentType(brotherId string, fileSeries string) (string, string, error) {
	var db *sql.DB = database.GetDB()
	var fileType string
	var contentType string
	query := fmt.Sprintf("SELECT %sFileType, %sContentType FROM personalDetails WHERE brotherId = ?", fileSeries, fileSeries)

	rows, err := db.Query(query, brotherId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {

		if err := rows.Scan(
			&fileType,
			&contentType,
		); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("ID: %s \n", brotherId)
		return fileType, contentType, nil
	} else {
		return fileType, contentType, errors.New("no results found")
	}
}
