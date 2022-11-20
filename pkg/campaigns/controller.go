package campaigns

import (
	"github.com/JumiaMDS/common4go/v2/pkg/api"
	"github.com/gin-gonic/gin"
	"gosqlmockdemo/definitions/campaigns"
	"net/http"
)

type Controller struct {
	srv campaigns.Service
}

func NewController(srv campaigns.Service) Controller {
	return Controller{
		srv: srv,
	}
}

func (ctrl Controller) GetCampaigns(c *gin.Context) {
	var req campaigns.ListFilter
	err := c.ShouldBindQuery(&req)
	if err != nil {
		api.HandleResponseError(c, err)
	}

	res, err := ctrl.srv.GetCampaigns(req)
	if err != nil {
		api.HandleResponseError(c, err)
		return
	}

	api.HandleResponse(c, http.StatusOK, api.Response{
		Data: res,
	})
}
