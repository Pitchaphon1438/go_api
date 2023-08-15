package orm

import (
	"time"

	"gorm.io/gorm"
)

//กำหนด field table User
type User struct {
	gorm.Model
	Username   string
	Password   string
	Fullname   string
	Avatar     string
	PersonalID string
	Email      string

	Rooms      []Room `gorm:"many2many:user_rooms"`
	UserTypeID int
	UserType   *UserType
}
type Room struct {
	gorm.Model
	RoomNumber string `gorm:"not null;"`
}

type UserType struct {
	gorm.Model
	Name        string `validate:"required"`
	Description string
}

type PostalStatus struct {
	gorm.Model
	Name        string
	Description string
}
type Postal struct {
	gorm.Model
	ImageUrl string
	Name     string

	PostalStatusID int
	RoomID         int
	UserID         int

	PostalStatus *PostalStatus
	Room         *Room
	User         *User

	ReceivedById int
	ReceivedBy   *User `gorm:"foreignKey:ReceivedById"`
	ReceivedDate time.Time
}
