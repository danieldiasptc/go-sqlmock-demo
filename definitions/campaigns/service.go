package campaigns

type Service interface {
	GetCampaigns(f ListFilter) ([]Campaign, error)
}
