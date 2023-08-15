package room

import (
	"memodule/jwt-api/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoomAll(c *gin.Context) {
	var rooms []orm.Room
	orm.Db.Find(&rooms)
	c.JSON(http.StatusOK, gin.H{"message": "get rooms success", "rooms": rooms})
}
func GetRoomById(c *gin.Context) {
	var room orm.Room
	id := c.Params.ByName("id")
	orm.Db.First(&room, id)
	c.JSON(http.StatusOK, gin.H{"message": "get room success", "room": room})
}
func CreateRoom(c *gin.Context) {
	var newRoom orm.Room
	if err := c.ShouldBindJSON(&newRoom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orm.Db.Create(&newRoom)
	c.JSON(http.StatusOK, gin.H{"message": "create room success", "newRoom": newRoom})
}
func UpdateRoom(c *gin.Context) {
	id := c.Params.ByName("id")
	var room orm.Room
	var updateRoom orm.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orm.Db.First(&updateRoom, id)
	updateRoom.RoomNumber = room.RoomNumber
	orm.Db.Save(updateRoom)
	c.JSON(http.StatusOK, gin.H{"message": "update room success", "update": updateRoom})
}
func DeleteRoom(c *gin.Context) {
	id := c.Params.ByName("id")
	var room orm.Room
	orm.Db.First(&room, id)
	orm.Db.Delete(&room)
	c.JSON(http.StatusOK, gin.H{"message": "delete room success", "update": room})
}
