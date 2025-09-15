package link

import (
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	S3id string `json:"s3id"`
}

func NewLink(url string) *Link {
	link := &Link{
		Url: url,
	}
	return link
}
