package campaigns

import "gosqlmockdemo/definitions/campaigns"

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) GetCampaigns(f campaigns.ListFilter) ([]campaigns.Campaign, error) {
	return nil, nil
}
