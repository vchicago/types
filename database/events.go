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

type Event struct {
	ID          int64           `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description" gorm:"type:text"`
	Start       time.Time       `json:"start"`
	End         time.Time       `json:"end"`
	Banner      File            `json:"banner"`
	Positions   []EventPosition `json:"positions"`
	SignUps     []EventSignUp   `json:"signups"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type EventPosition struct {
	ID        int64     `json:"id"`
	Event     Event     `json:"event"`
	Position  int64     `json:"position"`
	TakenBy   User      `json:"taken_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EventSignUp struct {
	ID        int64     `json:"id"`
	Event     Event     `json:"event"`
	User      User      `json:"user"`
	First     string    `json:"first"`
	Second    string    `json:"second"`
	Third     string    `json:"third"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
