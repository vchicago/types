package database

import (
	"encoding/json"
	"time"
)

type OAuthClient struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name" gorm:"primaryKey"`
	ClientId     string `json:"client_id" gorm:"type:varchar(128)"`
	ClientSecret string `json:"-" gorm:"type:varchar(255)"`
	RedirectURIs string `json:"return_uris" gorm:"type:text"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (c *OAuthClient) ValidURI(uri string) (bool, error) {
	uris := []string{}
	err := json.Unmarshal([]byte(c.RedirectURIs), &uris)
	if err != nil {
		return false, err
	}
	for _, v := range uris {
		if uri == v {
			return true, nil
		}
	}

	return false, nil
}

func (OAuthClient) TableName() string {
	return "oauth_clients"
}
