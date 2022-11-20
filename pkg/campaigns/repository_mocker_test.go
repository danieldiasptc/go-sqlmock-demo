package campaigns_test

import (
	mocks "github.com/JumiaMDS/common4go/v2/mocks/pkg/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	campaigns2 "gosqlmockdemo/definitions/campaigns"
	"gosqlmockdemo/pkg/campaigns"
	"testing"
	"time"
)

func setupRepository(t *testing.T) (campaigns.Repository, *mocks.Database) {
	dbMock := mocks.NewDatabase(t)
	return campaigns.NewRepository(dbMock), dbMock
}

func TestRepository_GetByID_mockery(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r, m := setupRepository(t)

		id := int64(123)
		ec := campaigns2.Campaign{
			ID:          123,
			Name:        "Campaign 1",
			Description: " description of the campaign 1",
			EndDate:     time.Date(2022, 12, 01, 22, 0, 0, 0, time.UTC),
		}

		m.On("Preload", "Tags").Return(m)
		m.On("Where", `"campaigns"."id" = ?`, id).Return(m)
		m.On("Find", mock.AnythingOfType("*campaigns.Campaign")).Return(m).
			Run(func(args mock.Arguments) {
				c := args.Get(0).(*campaigns2.Campaign)
				*c = ec
			})
		m.On("Error").Return(nil)

		res, err := r.GetByID(id)

		assert.Nil(t, err)
		assert.Equal(t, ec, res)
	})
}

func TestRepository_Create_mockery(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r, m := setupRepository(t)

		id := int64(123)
		cr := campaigns2.Campaign{
			ID:          0,
			Name:        "name",
			Description: "desc",
			EndDate:     time.Date(2022, time.December, 12, 12, 12, 0, 0, time.UTC),
			Tags: []campaigns2.Tag{
				{
					ID:   0,
					Name: "tag1",
				},
			},
		}

		m.On("Begin").Return(m)
		m.On("Create", mock.AnythingOfType("*campaigns.Campaign")).Return(m).
			Run(func(args mock.Arguments) {
				c := args.Get(0).(*campaigns2.Campaign)
				c.ID = id
			})
		m.On("Error").Return(nil)
		m.On("Commit").Return(m)

		res, err := r.Create(cr)

		cx := cr
		cx.ID = id

		assert.Nil(t, err)
		assert.Equal(t, cx, res)
	})
}
