package postals

import (
	"memodule/jwt-api/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePostalStatus(c *gin.Context) {
	var newPostalStatus orm.PostalStatus
	if err := c.ShouldBindJSON(&newPostalStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orm.Db.Create(&newPostalStatus)
	c.JSON(http.StatusOK, gin.H{"message": "create postal status success", "newPostalStatus": newPostalStatus})
}
func GetPostalStatuses(c *gin.Context) {
	var postalStatuses []orm.PostalStatus
	orm.Db.Find(&postalStatuses)
	c.JSON(http.StatusOK, gin.H{"message": "get postal status success", "postalStatus": postalStatuses})
}
func GetPostalStatusById(c *gin.Context) {
	var postalStatus orm.PostalStatus
	id := c.Params.ByName("id")
	orm.Db.First(&postalStatus, id)
	c.JSON(http.StatusOK, gin.H{"message": "get postal status success", "postalStatus": postalStatus})
}
func DeletePostalStatus(c *gin.Context) {
	id := c.Params.ByName("id")
	var PostalStatus orm.PostalStatus
	orm.Db.First(&PostalStatus, id)
	orm.Db.Delete(&PostalStatus)
	c.JSON(http.StatusOK, gin.H{"message": "delete postal status success", "update": PostalStatus})
}
func UpdatePostalStatus(c *gin.Context) {
	id := c.Params.ByName("id")
	var postalStatus orm.PostalStatus
	var updatePostalStatus orm.PostalStatus
	if err := c.ShouldBindJSON(&postalStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orm.Db.First(&updatePostalStatus, id)
	updatePostalStatus.Name = postalStatus.Name
	updatePostalStatus.Description = postalStatus.Description
	orm.Db.Save(updatePostalStatus)
	c.JSON(http.StatusOK, gin.H{"message": "update postal status success", "update": updatePostalStatus})
}
