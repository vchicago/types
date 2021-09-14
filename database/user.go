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

type User struct {
	CID               uint      `json:"cid" gorm:"primaryKey;type:int(11);column:cid"`
	FirstName         string    `json:"firstname" gorm:"type:varchar(128)"`
	LastName          string    `json:"lastname" gorm:"type:varchar(128)"`
	Email             string    `json:"-" yaml:"-" xml:"-" gorm:"type:varchar(128);index"`
	ControllerType    string    `json:"controllerType" gorm:"type:enum('none', 'visitor', 'home');not null;default:'none'"`
	OperatingInitials string    `json:"operatingInitials" gorm:"type:char(2);index"`
	RatingID          int       `json:"-" yaml:"-" xml:"-"`
	Rating            Rating    `json:"rating"`
	Roles             []*Role   `json:"roles" gorm:"many2many:user_roles"`
	Status            string    `json:"status" gorm:"type:enum('none', 'active', 'inactive', 'leave');not null;default:'none'"`
	Delivery          string    `json:"delivery" gorm:"type:enum('none','minor','major-solo','certified');not null;default:'none'"`
	Ground            string    `json:"ground" gorm:"type:enum('none','minor','major-solo','certified');not null;default:'none'"`
	Local             string    `json:"local" gorm:"type:enum('none','minor','major-solo','certified');not null;default:'none'"`
	Approach          string    `json:"approach" gorm:"type:enum('none','minor','major-solo','certified');not null;default:'none'"`
	Enroute           string    `json:"enroute" gorm:"type:enum('none','major-solo','certified');not null;default:'none'"`
	UpdateId          string    `json:"updateId" gorm:"type:varchar(21)"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
