package main

import (
	"fmt"
	AuthController "memodule/jwt-api/controller/auth"
	PostalsController "memodule/jwt-api/controller/postals"
	RoomController "memodule/jwt-api/controller/room"
	UserController "memodule/jwt-api/controller/user"
	"memodule/jwt-api/middelware"

	"memodule/jwt-api/orm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "golang.org/x/crypto/bcrypt"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}
	orm.InitDB()
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)
	authorized := r.Group("/users", middelware.JWTAuthen())
	authorized.GET("/getUsersAll", UserController.GetUserAll)
	authorized.GET("/getProfile", UserController.GetProfile)
	authorized.PUT("/update", UserController.UpdateUser)
	authorized.POST("/uploadAvatar", UserController.UploadAvatar)

	userType := r.Group("/userType")
	userType.POST("create", UserController.CreateUserType)
	userType.GET("/get", UserController.GetUserTypeAll)
	userType.GET("/get/:id", UserController.GetUserTypeById)
	userType.PUT("/update/:id,", UserController.UpdateUserType)
	userType.DELETE("/delete/:id", UserController.DeleteUserType)

	room := r.Group("/room")
	room.GET("/get", RoomController.GetRoomAll)
	room.GET("/get/:id", RoomController.GetRoomById)
	room.POST("/create", RoomController.CreateRoom)
	room.PUT("/update/:id", RoomController.UpdateRoom)
	room.DELETE("/delete/:id", RoomController.DeleteRoom)

	postals := r.Group("/postals")
	postals.POST("/create", PostalsController.CreatePostal)
	postals.GET("/get", PostalsController.GetPostals)
	postals.PUT("/update/:id", PostalsController.UpdatePostal)
	postals.POST("/upload/:id", PostalsController.UploadPostal)

	postalStatuses := r.Group("/postalStatuses")
	postalStatuses.POST("/create", PostalsController.CreatePostalStatus)
	postalStatuses.GET("/get", PostalsController.GetPostalStatuses)
	postalStatuses.GET("/get/:id", PostalsController.GetPostalStatusById)
	postalStatuses.PUT("/update/:id", PostalsController.UpdatePostalStatus)
	postalStatuses.DELETE("/delete/:id", PostalsController.DeletePostalStatus)

	r.Run("localhost:8080")
}
