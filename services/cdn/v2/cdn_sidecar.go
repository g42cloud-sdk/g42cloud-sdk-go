package v2

import "github.com/g42cloud-sdk/g42cloud-sdk-go/services/cdn/v2/model"

func (c *CdnClient) GetShowDomainLocationStatsRequest() *model.ShowDomainLocationStatsRequest {
	return new(model.ShowDomainLocationStatsRequest)
}

func (c *CdnClient) GetShowDomainStatsRequest() *model.ShowDomainStatsRequest {
	return new(model.ShowDomainStatsRequest)
}

func (c *CdnClient) GetShowTopUrlRequest() *model.ShowTopUrlRequest {
	return new(model.ShowTopUrlRequest)
}
