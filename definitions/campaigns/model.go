package campaigns

import (
	"time"
)

type Campaign struct {
	ID          int64  `gorm:"primary_key" json:"id"`
	Name        string `gorm:"type:varchar; column:name"`
	Description string
	EndDate     time.Time `gorm:"type:date; column:end_date"`
	Tags        []Tag     `gorm:"many2many:campaign_tag"`
}

type Tag struct {
	ID   int64  `gorm:"primary_key" json:"id"`
	Name string `gorm:"type:varchar; column:name"`
}
