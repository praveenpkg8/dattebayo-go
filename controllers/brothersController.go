package controllers

import (
	"fmt"
	"net/smtp"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/praveenpkg8/dattebayo-go/models"
)

var sampleSecretKey = []byte("GoLinuxCloudKey")

func CreateBrothers(
	newBrothers models.Brothers,
	brotherPersonalDetails models.PersonalDetails,
) (string, string) {
	brotherId, _ := models.AddBrothers(newBrothers)
	brPerDetId := models.AddPersonalDetails(brotherId, brotherPersonalDetails)
	return brotherId, brPerDetId
}


func SendVerificationEmail(brotherId string, email string) {
	verificationToken, _ := GenerateJWT(brotherId, email)

	// Sender data.
	from := "from@gmail.com"
	password := "<Email Password>"

	// Receiver email address.
	to := []string{
		"sender@example.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(verificationToken)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")

}

func GenerateJWT(brotherId string, email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["brother_id"] = brotherId
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(sampleSecretKey)

	if err != nil {
		fmt.Printf("Something Went Wrong: %v", err.Error())
		return "", err
	}
	return tokenString, nil
}

func VerifyBrother(brotherId string, approvalStatus string) {
	models.VerifyBrother(brotherId, approvalStatus, "")
}

func GetBrotherById(brotherId string) (models.Brothers, error) {
	brotherDetails, err := models.GetBrotherById(brotherId)
	return brotherDetails, err
}

func GetAllBrothers() ([]models.Brothers, error) {
	brotherList, err := models.GetAllBrothers()
	return brotherList, err
}
