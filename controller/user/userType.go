package user

import (
	"memodule/jwt-api/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserType(c *gin.Context) {
	var newUserType orm.UserType
	if err := c.ShouldBindJSON(&newUserType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orm.Db.Create(&newUserType)
	c.JSON(http.StatusOK, gin.H{"message": "create user_type success", "newUserType": newUserType})
}
func GetUserTypeAll(c *gin.Context) {
	var userType []orm.UserType
	orm.Db.Find(&userType)
	c.JSON(http.StatusOK, gin.H{"message": "get user type success", "userType": userType})
}
func GetUserTypeById(c *gin.Context) {
	id := c.Params.ByName("id")
	var userType orm.UserType
	orm.Db.First(&userType, id)
	c.JSON(http.StatusOK, gin.H{"message": "get user type success", "userType": userType})
}
func UpdateUserType(c *gin.Context) {
	id := c.Params.ByName("id")
	var userType orm.UserType
	var updateuserType orm.UserType
	if err := c.ShouldBindJSON(&userType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orm.Db.First(&updateuserType, id)
	updateuserType.Name = userType.Name
	updateuserType.Description = userType.Description
	orm.Db.Save(updateuserType)
	c.JSON(http.StatusOK, gin.H{"message": "update user type success", "update": updateuserType})
}
func DeleteUserType(c *gin.Context) {
	id := c.Params.ByName("id")
	var userType orm.UserType
	orm.Db.First(&userType, id)
	orm.Db.Delete(&userType)
	c.JSON(http.StatusOK, gin.H{"message": "delete user type success", "update": userType})
}
