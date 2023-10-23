package server

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/praveenpkg8/dattebayo-go/controllers"
	"github.com/praveenpkg8/dattebayo-go/models"
	"github.com/praveenpkg8/dattebayo-go/scripts"
)

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

func GetPersonalDetByBrotherId(c *gin.Context) {
	brotherId := c.Param("brothersId")
	personalDetails, err := controllers.GetPersonalDetByBrotherID(brotherId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, personalDetails)
}

func populateDbScript(c *gin.Context) {
	scripts.PopulateDB()
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "DB populated"})
}

func handleCurrentFileUpload(c *gin.Context) {
	brotherId := c.Param("brothersId")
	fileSeries := "current"
	file, err := c.FormFile("file")
	fileExtension := filepath.Ext(file.Filename)
	contentType := file.Header.Get("Content-Type")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	outputFilePath := fmt.Sprintf("./uploads/%s", brotherId)
	perDetId, err := controllers.GetPerIdByBrotherId(brotherId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	controllers.UpdateFileType(perDetId, fileSeries, fileExtension, contentType)
	fileName := fmt.Sprintf("%s%s.enc", fileSeries, fileExtension)
	err = controllers.EncryptAndSave(file, outputFilePath, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded and encrypted successfully!"})
}

func handlePermanentFileUpload(c *gin.Context) {
	brotherId := c.Param("brothersId")
	fileSeries := "permanent"
	file, err := c.FormFile("file")
	fileExtension := filepath.Ext(file.Filename)
	contentType := file.Header.Get("Content-Type")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	outputFilePath := fmt.Sprintf("./uploads/%s", brotherId)
	perDetId, err := controllers.GetPerIdByBrotherId(brotherId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	controllers.UpdateFileType(perDetId, fileSeries, fileExtension, contentType)
	fileName := fmt.Sprintf("%s%s.enc", fileSeries, fileExtension)
	err = controllers.EncryptAndSave(file, outputFilePath, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded and encrypted successfully!"})
}

func handleCurrentFileDownload(c *gin.Context) {
	brotherId := c.Param("brothersId")
	fileSeries := "current"
	fileType, fileContent, err := controllers.GetFileExtAndContentType(brotherId, fileSeries)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileName := fmt.Sprintf("%s%s.enc", fileSeries, fileType)
	EncFilePath := fmt.Sprintf("./uploads/%s/%s", brotherId, fileName)
	decryptFile, err := controllers.GetDecryptFile(EncFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", fileContent)
	c.Writer.Write(decryptFile)
}

func handlePermanentFileDownload(c *gin.Context) {
	brotherId := c.Param("brothersId")
	fileSeries := "permanent"
	fileType, fileContent, err := controllers.GetFileExtAndContentType(brotherId, fileSeries)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileName := fmt.Sprintf("%s%s.enc", fileSeries, fileType)
	EncFilePath := fmt.Sprintf("./uploads/%s/%s", brotherId, fileName)
	decryptFile, err := controllers.GetDecryptFile(EncFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", fileContent)
	c.Writer.Write(decryptFile)
}

// func handleFileDecrypt(c *gin.Context) {
// 	brotherId := c.Param("brothersId")
// 	fileSeries := c.Param("fileSeries")
// 	fileName := fmt.Sprintf("%s%s.enc", fileSeries, fileExtension)
// 	decrytFileName := fmt.Sprintf("./uploads/%s/%s", brotherId, fileName)
// 	err := controllers.DecryptFile(
// 		decrytFileName,
// 		"test-decrypt.jpeg",
// 	)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "File Decrypted successfully!"})
// }

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Init() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pang",
		})
	})
	r.GET("/api/v1/brothers", GetAllBrothers)
	r.GET("/api/v1/brothers/:brothersId", GetBrothersById)
	r.POST("/api/v1/brothers", CreateBrothers)
	r.POST("/api/v1/brothers/:brothersId/verify", VerifyBrother)

	r.GET("/api/v1/brothers/:brothersId/personalDetails", GetPersonalDetByBrotherId)
	r.PATCH("/api/v1/brothers/:brothersId/personalDetails/:perDetId/contacts", updatePersonalContact)
	r.POST("/api/v1/upload/:brothersId/current", handleCurrentFileUpload)
	r.POST("/api/v1/upload/:brothersId/permanent", handlePermanentFileUpload)
	r.POST("/api/v1/download/:brothersId/current", handleCurrentFileDownload)
	r.POST("/api/v1/download/:brothersId/permanent", handlePermanentFileDownload)
	r.GET("/api/v1/populate-db", populateDbScript)

	// r.GET("/api/v1/decrypt", handleFileDecrypt)

	r.Run() // listen and serve on 0.0.0.0:8080
}
