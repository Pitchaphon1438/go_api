package postals

import (
	"fmt"
	"memodule/jwt-api/orm"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreatePostal(c *gin.Context) {
	var newPostal orm.Postal
	if err := c.ShouldBindJSON(&newPostal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orm.Db.Create(&newPostal)
	c.JSON(http.StatusOK, gin.H{"message": "create postal success", "newPostal": newPostal})
}
func GetPostals(c *gin.Context) {
	var postals []orm.Postal
	orm.Db.Find(&postals)
	c.JSON(http.StatusOK, gin.H{"message": "get postals success", "Postals": postals})
}
func GetPostalById(c *gin.Context) {
	var postal orm.Room
	id := c.Params.ByName("id")
	orm.Db.First(&postal, id)
	c.JSON(http.StatusOK, gin.H{"message": "get postals success", "postal": postal})
}
func UpdatePostal(c *gin.Context) {
	id := c.Params.ByName("id")
	var postal orm.Postal
	var updatePostal orm.Postal
	err := c.ShouldBindJSON(&postal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if postal.ID == 0 {
		c.JSON(404, gin.H{"message": "no postal found",
			"error": err.Error()})
		return
	}
	orm.Db.First(&updatePostal, id)
	updatePostal.Name = postal.Name
	updatePostal.ImageUrl = postal.ImageUrl
	updatePostal.PostalStatusID = postal.PostalStatusID
	updatePostal.UserID = postal.PostalStatusID
	updatePostal.ReceivedDate = time.Now()
	orm.Db.Save(updatePostal)
	c.JSON(http.StatusOK, gin.H{"message": "update postal success", "update": updatePostal})
}
func UploadPostal(c *gin.Context) {
	id := c.Params.ByName("id")
	var postal orm.Postal
	orm.Db.First(&postal, id)
	file, err := c.FormFile("file")
	if postal.ID == 0 {
		c.JSON(404, gin.H{"message": "no postal found",
			"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(404, gin.H{"message": "not able upload your file",
			"error": err.Error(),
		})
		return
	}
	randomUuidString := uuid.NewString()
	c.SaveUploadedFile(file, fmt.Sprintf("./upload/%s_%s", randomUuidString, file.Filename))
	postal.ImageUrl = fmt.Sprintf("/upload/%s_%s", randomUuidString, file.Filename)
	fmt.Print(postal.ImageUrl)
	orm.Db.Save(&postal)
	c.JSON(http.StatusOK, gin.H{"message": "upload success", "update": postal})
}
