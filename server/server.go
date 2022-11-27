package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/praveenpkg8/dattebayo-go/controllers"
	"github.com/praveenpkg8/dattebayo-go/models"
)

func postAlbums(c *gin.Context) {
	var newBrothers models.Brothers
	var botherPersonalDetails models.PersonalDetails

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.ShouldBindBodyWith(&newBrothers, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}
	if err := c.ShouldBindBodyWith(&botherPersonalDetails, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}
	brotherId, brPerDetId := controllers.CreateBrothers(newBrothers, botherPersonalDetails)
	botherPersonalDetails.Id = brPerDetId
	botherPersonalDetails.BrotherId = brotherId

	// Add the new album to the slice.
	c.IndentedJSON(http.StatusCreated, botherPersonalDetails)
}


func updatePersonalContact(c *gin.Context) {
	var botherPersonalDetails models.PersonalDetails

	if err := c.ShouldBindBodyWith(&botherPersonalDetails, binding.JSON); err != nil {
		log.Printf("%+v", err)
	
	}
	brotherId := c.Param("brothersId")
	perDetId := c.Param("perDetId")
	botherPersonalDetails.Id = perDetId
	botherPersonalDetails.BrotherId = brotherId
	controllers.UpdatePersonalContact(perDetId, botherPersonalDetails)
	c.IndentedJSON(http.StatusCreated, botherPersonalDetails)
}



func Init() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/api/v1/brothers", postAlbums)
	r.PATCH("/api/v1/brothers/:brothersId/personalDetail/:perDetId/contact", updatePersonalContact)


	r.Run() // listen and serve on 0.0.0.0:8080
}
