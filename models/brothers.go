package models

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/praveenpkg8/dattebayo-go/database"
	// "github.com/google/uuid"
)

type Address struct {
	DoorNo     string `json:"door_no"`
	StreetName string `json:"street_name"`
	Area       string `json:"area"`
	District   string `json:"district"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PinCode    string `json:"pin_code"`
}

type PersonalDetails struct {
	Id               string  `json:"id,omitempty"`
	UserId           string  `json:"user_id"`
	DoB              string  `json:"DoB"`
	Email            string  `json:"email"`
	PhoneNumber      string  `json:"phone_number"`
	AltPhoneNumber   string  `json:"alt_phone_number"`
	CurrentAddress   Address `json:"current_address"`
	PermanentAddress Address `json:"permanent_address"`
}

type Brothers struct {
	Id             string `json:"id,omitempty"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	SudoName       string `json:"sudo_name"`
	ReferralCode   string `json:"referral_code"`
	ReferredBy     string `json:"referred_by"`
	ApprovalStatus string `json:"approval_status"`
}

func AddBrothers(brotherId string, newBrothers Brothers) {
	log.Println("Adding brother")
	var err error
	var db *sql.DB = database.GetDB()

	id := uuid.New()
	stmt, err := db.Prepare("INSERT INTO brothers(id, firstName, lastName, referralCode) values(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(
		brotherId,
		newBrothers.FirstName,
		newBrothers.LastName,
		newBrothers.ReferralCode,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	log.Printf("Added Brother: %v | %v | %v \n", id.String(), newBrothers.FirstName, newBrothers.LastName)
}

func AddPersonalDetails(brotherId string, brotherPersonalDetails PersonalDetails) {
	log.Println("Adding Email")
	var err error
	var db *sql.DB = database.GetDB()

	id := uuid.New()
	stmt, err := db.Prepare("INSERT INTO personalDetails(id, userId, email) values(?, ?, ?)")
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
}
