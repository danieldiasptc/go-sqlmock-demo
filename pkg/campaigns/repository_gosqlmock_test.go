package campaigns_test

import (
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	cGorm "github.com/JumiaMDS/common4go/v2/pkg/db/gorm"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	campaigns2 "gosqlmockdemo/definitions/campaigns"
	"gosqlmockdemo/pkg/campaigns"
	"regexp"
	"testing"
	"time"
)

func setupIntegrationRepository(t *testing.T) (campaigns.Repository, sqlmock.Sqlmock) {
	db, m, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	pgc := postgres.New(postgres.Config{
		Conn: db,
	})
	cfg := gorm.Config{
		SkipDefaultTransaction: true,
	}
	gDB, err := cGorm.New(pgc, cfg, 1, 1, 300)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating grom database connection", err)
	}
	gDB = gDB.Session(&gorm.Session{})

	m.MatchExpectationsInOrder(true)

	repo := campaigns.NewRepository(gDB)
	return repo, m
}

func TestRepository_GetByID_sqlmock(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r, m := setupIntegrationRepository(t)

		id := int64(123)
		tagId := int64(321)
		ec := campaigns2.Campaign{
			ID:          123,
			Name:        "Campaign 1",
			Description: " description of the campaign 1",
			EndDate:     time.Date(2022, 12, 01, 22, 0, 0, 0, time.UTC),
			Tags: []campaigns2.Tag{
				{ID: tagId, Name: "tag1"},
			},
		}

		cols := []string{"id", "name", "description", "end_date"}
		values := []driver.Value{ec.ID, ec.Name, ec.Description, ec.EndDate}

		q := `SELECT * FROM "campaigns" WHERE "campaigns"."id" = $1`
		m.ExpectQuery(regexp.QuoteMeta(q)).
			WithArgs(id).
			WillReturnRows(sqlmock.NewRows(cols).AddRow(values...))

		cols2 := []string{"campaign_id", "tag_id"}
		values2 := []driver.Value{id, tagId}
		q2 := `SELECT * FROM "campaign_tag" WHERE "campaign_tag"."campaign_id" = $1`
		m.ExpectQuery(regexp.QuoteMeta(q2)).
			WithArgs(id).
			WillReturnRows(sqlmock.NewRows(cols2).AddRow(values2...))

		cols3 := []string{"id", "name"}
		values3 := []driver.Value{tagId, ec.Tags[0].Name}
		q3 := `SELECT * FROM "tags" WHERE "tags"."id" = $1`
		m.ExpectQuery(regexp.QuoteMeta(q3)).
			WithArgs(ec.Tags[0].ID).
			WillReturnRows(sqlmock.NewRows(cols3).AddRow(values3...))

		c, err := r.GetByID(id)

		err = m.ExpectationsWereMet()
		if err != nil {
			t.Errorf("Failed to meet expectations, got error: %v", err)
		}
		assert.Nil(t, err)
		assert.Equal(t, ec, c)
	})
}
