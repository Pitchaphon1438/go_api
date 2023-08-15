package user

import (
	"fmt"
	"memodule/jwt-api/orm"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserAll(c *gin.Context) {
	var users []orm.User
	orm.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "get user success", "users": users})
}
func GetProfile(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	var user orm.User
	orm.Db.First(&user, userId)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "get user success", "user": user})
}
func UpdateUser(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	var user orm.User
	var updateUser orm.User
	orm.Db.First(&user, userId)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orm.Db.First(&updateUser, userId)
	updateUser.Username = user.Username
	updateUser.Password = user.Password
	updateUser.Fullname = user.Fullname
	updateUser.PersonalID = user.PersonalID
	updateUser.Email = user.Email
	updateUser.UserTypeID = user.UserTypeID
	orm.Db.Save(updateUser)
	c.JSON(http.StatusOK, gin.H{"message": "update user success", "update": updateUser})
}
func UploadAvatar(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	var user orm.User
	orm.Db.First(&user, userId)
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(404, gin.H{"message": "not able upload your file",
			"error": err.Error(),
		})
		return
	}
	randomUuidString := uuid.NewString()
	c.SaveUploadedFile(file, fmt.Sprintf("./upload/%s_%s", randomUuidString, file.Filename))
	user.Avatar = fmt.Sprintf("/upload/%s_%s", randomUuidString, file.Filename)
	orm.Db.Save(&user)
	c.JSON(http.StatusOK, gin.H{"message": "upload success", "update": user})
}
