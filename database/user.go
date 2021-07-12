package database

import "time"

type User struct {
	CID               uint      `json:"cid" gorm:"primaryKey"`
	FirstName         string    `json:"firstname" gorm:"type:varchar(128)"`
	LastName          string    `json:"lastname" gorm:"type:varchar(128)"`
	Email             string    `json:"-" gorm:"type:varchar(128);index"`
	ControllerType    string    `json:"controllerType" gorm:"type:enum('none', 'visitor', 'home')"`
	OperatingInitials string    `json:"operatingInitials" gorm:"type:char(2);index"`
	RatingID          int       `json:"-"`
	Rating            Rating    `json:"rating"`
	Status            string    `json:"status" gorm:"type:enum('none', 'active', 'inactive', 'leave')"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
