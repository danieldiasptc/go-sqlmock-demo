package campaigns

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gosqlmockdemo/definitions/campaigns"
	mocks "gosqlmockdemo/mocks/definitions/campaigns"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setupController(t *testing.T) (Controller, *mocks.Service) {

	srv := mocks.NewService(t)

	return NewController(srv), srv
}

func TestController_GetCampaigns(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		r := gin.Default()

		ctrl, m := setupController(t)
		r.GET("/", ctrl.GetCampaigns)

		// mocks
		mReq := campaigns.ListFilter{
			Query:  "asd",
			Status: "enabled",
		}
		mRes := []campaigns.Campaign{
			{
				ID:          1,
				Name:        "c1",
				Description: "d1",
				EndDate:     time.Date(2022, time.September, 12, 12, 12, 0, 0, time.UTC),
				Tags: []campaigns.Tag{{
					ID:   1,
					Name: "tag1",
				}},
			},
		}

		m.On("GetCampaigns", mReq).Return(mRes, nil)

		// request
		url := fmt.Sprintf("/?query=%s&status=%s", mReq.Query, mReq.Status)
		req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
		if err != nil {
			t.Errorf("Error requesting test controller: %v\n", err)
		}
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		// asserts
		assert.Equal(t, http.StatusOK, rr.Code)
		expRes := `{"data":[{"id":1,"Name":"c1","Description":"d1","EndDate":"2022-09-12T12:12:00Z","Tags":[{"id":1,"Name":"tag1"}]}],"links":null,"meta":null,"jsonapi":null}`
		assert.Equal(t, expRes, rr.Body.String())
	})
}
