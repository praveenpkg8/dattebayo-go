package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/praveenpkg8/dattebayo-go/controllers"
	"github.com/praveenpkg8/dattebayo-go/models"
)

func GetAllBrothers(c *gin.Context) {
	brotherList, err := controllers.GetAllBrothers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, brotherList)
}

func GetBrothersById(c *gin.Context) {
	brotherId := c.Param("brothersId")
	brotherDetail, err := controllers.GetBrotherById(brotherId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, brotherDetail)
}

func VerifyBrother(c *gin.Context) {
	brotherId := c.Param("brothersId")
	var brotherDetails models.Brothers
	if err := c.ShouldBindBodyWith(&brotherDetails, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}
	controllers.VerifyBrother(brotherId, brotherDetails.ApprovalStatus)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Status Updated"})
}

func CreateBrothers(c *gin.Context) {
	var newBrothers models.Brothers
	var botherPersonalDetails models.PersonalDetails
	var brotherResp models.BrotherResponse

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.ShouldBindBodyWith(&newBrothers, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}
	if err := c.ShouldBindBodyWith(&botherPersonalDetails, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}
	brotherId, brPerDetId := controllers.CreateBrothers(newBrothers, botherPersonalDetails)
	brotherResp.Id = brotherId
	brotherResp.PersonalDetailId = brPerDetId

	// Add the new album to the slice.
	c.IndentedJSON(http.StatusCreated, brotherResp)

}
