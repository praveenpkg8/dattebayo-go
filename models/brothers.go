package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/praveenpkg8/dattebayo-go/database"
	// "github.com/google/uuid"
)

type Address struct {
	PhoneNo    string `json:"phone_no"`
	DoorNo     string `json:"door_no"`
	StreetName string `json:"street_name"`
	Area       string `json:"area"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PinCode    string `json:"pin_code"`
}

type PersonalDetails struct {
	Id                     string  `json:"id,omitempty"`
	BrotherId              string  `json:"brother_id"`
	DoB                    string  `json:"DoB"`
	Email                  string  `json:"email"`
	PhoneNumber            string  `json:"phone_number"`
	AltPhoneNumber         string  `json:"alt_phone_number"`
	AltContactName         string  `json:"alt_contact_name"`
	AltContactRelationship string  `json:"alt_contact_relationship"`
	CurrentAddress         Address `json:"current_address"`
	CurrentFileType	string `json:"current_file_type"`
	CurrentContentType	string `json:"current_content_type"`
	PermanentAddress       Address `json:"permanent_address"`
	PermanentFileType	string `json:"permanent_file_type"`
	PermanentContentType	string `json:"permanent_content_type"`
}

type Brothers struct {
	Id                      string `json:"id,omitempty"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	SudoName                string `json:"sudo_name"`
	ReferralCode            string `json:"referral_code"`
	ReferredBy              string `json:"referred_by"`
	EmailVerificationStatus string `json:"email_verification_status"`
	ApprovalStatus          string `json:"approval_status"`
	ApprovedBy              string `json:"approvedBy"`
}

type BrotherResponse struct {
	Id               string `json:"id"`
	PersonalDetailId string `json:"personal_detail_id"`
}

func AddBrothers(newBrothers Brothers) (string, error) {
	log.Println("Adding brother")
	var db *sql.DB = database.GetDB()
	brotherId := uuid.New().String()
	stmt, err := db.Prepare("INSERT INTO brothers(id, firstName, lastName, sudoName, referralCode) VALUES(?, ?, ?, ?)")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		brotherId,
		newBrothers.FirstName,
		newBrothers.LastName,
		newBrothers.ReferralCode,
	)
	if err != nil {
		return "", err
	}

	log.Printf("Added Brother: %v | %v | %v\n", brotherId, newBrothers.FirstName, newBrothers.LastName)
	return brotherId, nil
}

func AddPersonalDetails(brotherId string, brotherPersonalDetails PersonalDetails) string {
	log.Println("Adding Email")
	var err error
	var db *sql.DB = database.GetDB()

	id := uuid.New()
	stmt, err := db.Prepare("INSERT INTO personalDetails(id, brotherId, email) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(
		id.String(),
		brotherId,
		brotherPersonalDetails.Email,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	log.Printf("Added PersonalDetails : %v | %v | %v \n", id.String(), brotherId, brotherPersonalDetails.Email)
	return id.String()
}

func VerifyBrother(brotherId string, approvalStatus string, approvedBy string) {
	log.Printf("Adding update brothers contact")
	var err error
	var db *sql.DB = database.GetDB()

	stmt, err := db.Prepare(
		"UPDATE brothers SET " +
			"approvalStatus = ?, approvedBy = ?" +
			"WHERE id = ?",
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(
		approvalStatus,
		approvedBy,
		brotherId,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
}

func GetBrotherById(brotherId string) (Brothers, error) {
	var db *sql.DB = database.GetDB()
	query := "SELECT * FROM brothers WHERE id = ?"
	rows, err := db.Query(query, brotherId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		var brothers Brothers // Create a Brothers struct to store the retrieved data
		if err := rows.Scan(&brothers.Id, &brothers.FirstName, &brothers.LastName, &brothers.SudoName, &brothers.ReferralCode, &brothers.ReferredBy, &brothers.EmailVerificationStatus, &brothers.ApprovalStatus, &brothers.ApprovedBy); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("ID: %s, First Name: %s, Last Name: %s, ...\n", brothers.Id, brothers.FirstName, brothers.LastName)
		return brothers, nil
	} else {
		return Brothers{}, errors.New("no results found") // Return an empty Brothers struct
	}
}

func GetAllBrothers() ([]Brothers, error) {
	var db *sql.DB = database.GetDB()
	query := "SELECT * FROM brothers"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var brothersList []Brothers

	for rows.Next() {
		var brothers Brothers
		if err := rows.Scan(
			&brothers.Id,
			&brothers.FirstName,
			&brothers.LastName,
			&brothers.SudoName, // Scan as NullString.String
			&brothers.ReferralCode,
			&brothers.ReferredBy,
			&brothers.EmailVerificationStatus,
			&brothers.ApprovalStatus,
			&brothers.ApprovedBy,
		); err != nil {
			log.Fatal(err)
		}
		brothersList = append(brothersList, brothers)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return brothersList, nil
}
