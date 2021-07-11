package database

import "time"

type OAuthLogin struct {
	ID                  uint        `json:"id" gorm:"primaryKey"`
	Token               string      `json:"token" gorm:"type:varchar(128)"`
	Code                string      `json:"code" gorm:"varchar(48)"`
	UserAgent           string      `json:"ua" gorm:"type:varchar(255)"`
	RedirectURI         string      `json:"url" gorm:"type:varchar(255)"`
	ClientID            uint        `json:"-"`
	Client              OAuthClient `json:"-"`
	State               string      `json:"state"`
	CodeChallenge       string      `json:"-"`
	CodeChallengeMethod string      `json:"-"`
	Scope               string      `json:"-"`
	CID                 uint        `json:"cid"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (OAuthLogin) TableName() string {
	return "oauth_logins"
}
