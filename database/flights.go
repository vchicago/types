/*
   ZAU Database Typings
   Copyright (C) 2021  Daniel A. Hawton <daniel@hawton.org>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package database

import "time"

type Flights struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Callsign    string    `json:"callsign" gorm:"index;type:varchar(10)"`
	CID         int       `json:"cid" gorm:"index;column:cid"`
	Name        string    `json:"name" gorm:"type:varchar(191)"`
	Facility    string    `json:"facility" gorm:"type:varchar(4)"`
	Latitude    float32   `json:"latitude" gorm:"type:float(10,8)"`
	Longitude   float32   `json:"longitude" gorm:"type:float(11,8)"`
	Groundspeed int       `json:"groundspeed"`
	Heading     int       `json:"heading"`
	Altitude    int       `json:"altitude"`
	Aircraft    string    `json:"aircraft" gorm:"type:varchar(10)"`
	Departure   string    `json:"departure" gorm:"type:varchar(4)"`
	Arrival     string    `json:"arrival" gorm:"type:varchar(4)"`
	Route       string    `json:"route" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
