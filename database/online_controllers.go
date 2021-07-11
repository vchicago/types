package database

import "time"

type OnlineControllers struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Callsign  string    `json:"callsign" gorm:"index;type:varchar(10)"`
	CID       int       `json:"cid" gorm:"index;column:cid"`
	Name      string    `json:"name" gorm:"type:varchar(191)"`
	Facility  string    `json:"facility" gorm:"type:varchar(4)"`
	Position  string    `json:"position" gorm:"type:varchar(10)"`
	Frequency string    `json:"frequency" gorm:"type:varchar(7)"`
	LogonTime time.Time `json:"logon_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
