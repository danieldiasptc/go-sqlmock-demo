package campaigns

import (
	"github.com/JumiaMDS/common4go/v2/pkg/db"
	"gosqlmockdemo/definitions/campaigns"
)

type Repository struct {
	conn db.Database
}

func NewRepository(conn db.Database) Repository {
	return Repository{
		conn: conn,
	}
}

func (r Repository) GetByID(id int64) (campaigns.Campaign, error) {
	var c campaigns.Campaign
	//err := r.conn.
	//	Preload("Tags").
	//	Where(`"campaigns"."id" = ?`, id).
	//	Find(&c).
	//	Error()

	c = campaigns.Campaign{
		ID: id,
	}
	err := r.conn.
		Find(&c).
		Error()

	return c, err
}
