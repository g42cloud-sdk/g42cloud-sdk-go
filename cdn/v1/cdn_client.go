package v1

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/cdn/v1/model"
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
)

type CdnClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCdnClient(hcClient *http_client.HcHttpClient) *CdnClient {
	return &CdnClient{HcClient: hcClient}
}

func CdnClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

func (c *CdnClient) CreateDomain(request *model.CreateDomainRequest) (*model.CreateDomainResponse, error) {
	requestDef := GenReqDefForCreateDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDomainResponse), nil
	}
}

func (c *CdnClient) CreateDomainInvoker(request *model.CreateDomainRequest) *CreateDomainInvoker {
	requestDef := GenReqDefForCreateDomain()
	return &CreateDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) CreatePreheatingTasks(request *model.CreatePreheatingTasksRequest) (*model.CreatePreheatingTasksResponse, error) {
	requestDef := GenReqDefForCreatePreheatingTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePreheatingTasksResponse), nil
	}
}

func (c *CdnClient) CreatePreheatingTasksInvoker(request *model.CreatePreheatingTasksRequest) *CreatePreheatingTasksInvoker {
	requestDef := GenReqDefForCreatePreheatingTasks()
	return &CreatePreheatingTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) CreateRefreshTasks(request *model.CreateRefreshTasksRequest) (*model.CreateRefreshTasksResponse, error) {
	requestDef := GenReqDefForCreateRefreshTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRefreshTasksResponse), nil
	}
}

func (c *CdnClient) CreateRefreshTasksInvoker(request *model.CreateRefreshTasksRequest) *CreateRefreshTasksInvoker {
	requestDef := GenReqDefForCreateRefreshTasks()
	return &CreateRefreshTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) DeleteDomain(request *model.DeleteDomainRequest) (*model.DeleteDomainResponse, error) {
	requestDef := GenReqDefForDeleteDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDomainResponse), nil
	}
}

func (c *CdnClient) DeleteDomainInvoker(request *model.DeleteDomainRequest) *DeleteDomainInvoker {
	requestDef := GenReqDefForDeleteDomain()
	return &DeleteDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) DisableDomain(request *model.DisableDomainRequest) (*model.DisableDomainResponse, error) {
	requestDef := GenReqDefForDisableDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableDomainResponse), nil
	}
}

func (c *CdnClient) DisableDomainInvoker(request *model.DisableDomainRequest) *DisableDomainInvoker {
	requestDef := GenReqDefForDisableDomain()
	return &DisableDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) EnableDomain(request *model.EnableDomainRequest) (*model.EnableDomainResponse, error) {
	requestDef := GenReqDefForEnableDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableDomainResponse), nil
	}
}

func (c *CdnClient) EnableDomainInvoker(request *model.EnableDomainRequest) *EnableDomainInvoker {
	requestDef := GenReqDefForEnableDomain()
	return &EnableDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) ListDomains(request *model.ListDomainsRequest) (*model.ListDomainsResponse, error) {
	requestDef := GenReqDefForListDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainsResponse), nil
	}
}

func (c *CdnClient) ListDomainsInvoker(request *model.ListDomainsRequest) *ListDomainsInvoker {
	requestDef := GenReqDefForListDomains()
	return &ListDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) ShowBlackWhiteList(request *model.ShowBlackWhiteListRequest) (*model.ShowBlackWhiteListResponse, error) {
	requestDef := GenReqDefForShowBlackWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBlackWhiteListResponse), nil
	}
}

func (c *CdnClient) ShowBlackWhiteListInvoker(request *model.ShowBlackWhiteListRequest) *ShowBlackWhiteListInvoker {
	requestDef := GenReqDefForShowBlackWhiteList()
	return &ShowBlackWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) ShowCacheRules(request *model.ShowCacheRulesRequest) (*model.ShowCacheRulesResponse, error) {
	requestDef := GenReqDefForShowCacheRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCacheRulesResponse), nil
	}
}

func (c *CdnClient) ShowCacheRulesInvoker(request *model.ShowCacheRulesRequest) *ShowCacheRulesInvoker {
	requestDef := GenReqDefForShowCacheRules()
	return &ShowCacheRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) ShowCertificatesHttpsInfo(request *model.ShowCertificatesHttpsInfoRequest) (*model.ShowCertificatesHttpsInfoResponse, error) {
	requestDef := GenReqDefForShowCertificatesHttpsInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertificatesHttpsInfoResponse), nil
	}
}

func (c *CdnClient) ShowCertificatesHttpsInfoInvoker(request *model.ShowCertificatesHttpsInfoRequest) *ShowCertificatesHttpsInfoInvoker {
	requestDef := GenReqDefForShowCertificatesHttpsInfo()
	return &ShowCertificatesHttpsInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) ShowDomainDetail(request *model.ShowDomainDetailRequest) (*model.ShowDomainDetailResponse, error) {
	requestDef := GenReqDefForShowDomainDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainDetailResponse), nil
	}
}

func (c *CdnClient) ShowDomainDetailInvoker(request *model.ShowDomainDetailRequest) *ShowDomainDetailInvoker {
	requestDef := GenReqDefForShowDomainDetail()
	return &ShowDomainDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) ShowDomainFullConfig(request *model.ShowDomainFullConfigRequest) (*model.ShowDomainFullConfigResponse, error) {
	requestDef := GenReqDefForShowDomainFullConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainFullConfigResponse), nil
	}
}

func (c *CdnClient) ShowDomainFullConfigInvoker(request *model.ShowDomainFullConfigRequest) *ShowDomainFullConfigInvoker {
	requestDef := GenReqDefForShowDomainFullConfig()
	return &ShowDomainFullConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) ShowDomainItemDetails(request *model.ShowDomainItemDetailsRequest) (*model.ShowDomainItemDetailsResponse, error) {
	requestDef := GenReqDefForShowDomainItemDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainItemDetailsResponse), nil
	}
}

func (c *CdnClient) ShowDomainItemDetailsInvoker(request *model.ShowDomainItemDetailsRequest) *ShowDomainItemDetailsInvoker {
	requestDef := GenReqDefForShowDomainItemDetails()
	return &ShowDomainItemDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) ShowDomainItemLocationDetails(request *model.ShowDomainItemLocationDetailsRequest) (*model.ShowDomainItemLocationDetailsResponse, error) {
	requestDef := GenReqDefForShowDomainItemLocationDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainItemLocationDetailsResponse), nil
	}
}

func (c *CdnClient) ShowDomainItemLocationDetailsInvoker(request *model.ShowDomainItemLocationDetailsRequest) *ShowDomainItemLocationDetailsInvoker {
	requestDef := GenReqDefForShowDomainItemLocationDetails()
	return &ShowDomainItemLocationDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) ShowDomainLocationStats(request *model.ShowDomainLocationStatsRequest) (*model.ShowDomainLocationStatsResponse, error) {
	requestDef := GenReqDefForShowDomainLocationStats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainLocationStatsResponse), nil
	}
}

func (c *CdnClient) ShowDomainLocationStatsInvoker(request *model.ShowDomainLocationStatsRequest) *ShowDomainLocationStatsInvoker {
	requestDef := GenReqDefForShowDomainLocationStats()
	return &ShowDomainLocationStatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) ShowDomainStats(request *model.ShowDomainStatsRequest) (*model.ShowDomainStatsResponse, error) {
	requestDef := GenReqDefForShowDomainStats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainStatsResponse), nil
	}
}

func (c *CdnClient) ShowDomainStatsInvoker(request *model.ShowDomainStatsRequest) *ShowDomainStatsInvoker {
	requestDef := GenReqDefForShowDomainStats()
	return &ShowDomainStatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) ShowHistoryTaskDetails(request *model.ShowHistoryTaskDetailsRequest) (*model.ShowHistoryTaskDetailsResponse, error) {
	requestDef := GenReqDefForShowHistoryTaskDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHistoryTaskDetailsResponse), nil
	}
}

func (c *CdnClient) ShowHistoryTaskDetailsInvoker(request *model.ShowHistoryTaskDetailsRequest) *ShowHistoryTaskDetailsInvoker {
	requestDef := GenReqDefForShowHistoryTaskDetails()
	return &ShowHistoryTaskDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) ShowHistoryTasks(request *model.ShowHistoryTasksRequest) (*model.ShowHistoryTasksResponse, error) {
	requestDef := GenReqDefForShowHistoryTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHistoryTasksResponse), nil
	}
}

func (c *CdnClient) ShowHistoryTasksInvoker(request *model.ShowHistoryTasksRequest) *ShowHistoryTasksInvoker {
	requestDef := GenReqDefForShowHistoryTasks()
	return &ShowHistoryTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) ShowHttpInfo(request *model.ShowHttpInfoRequest) (*model.ShowHttpInfoResponse, error) {
	requestDef := GenReqDefForShowHttpInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpInfoResponse), nil
	}
}

func (c *CdnClient) ShowHttpInfoInvoker(request *model.ShowHttpInfoRequest) *ShowHttpInfoInvoker {
	requestDef := GenReqDefForShowHttpInfo()
	return &ShowHttpInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) ShowIpInfo(request *model.ShowIpInfoRequest) (*model.ShowIpInfoResponse, error) {
	requestDef := GenReqDefForShowIpInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIpInfoResponse), nil
	}
}

func (c *CdnClient) ShowIpInfoInvoker(request *model.ShowIpInfoRequest) *ShowIpInfoInvoker {
	requestDef := GenReqDefForShowIpInfo()
	return &ShowIpInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) ShowLogs(request *model.ShowLogsRequest) (*model.ShowLogsResponse, error) {
	requestDef := GenReqDefForShowLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLogsResponse), nil
	}
}

func (c *CdnClient) ShowLogsInvoker(request *model.ShowLogsRequest) *ShowLogsInvoker {
	requestDef := GenReqDefForShowLogs()
	return &ShowLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) ShowOriginHost(request *model.ShowOriginHostRequest) (*model.ShowOriginHostResponse, error) {
	requestDef := GenReqDefForShowOriginHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOriginHostResponse), nil
	}
}

func (c *CdnClient) ShowOriginHostInvoker(request *model.ShowOriginHostRequest) *ShowOriginHostInvoker {
	requestDef := GenReqDefForShowOriginHost()
	return &ShowOriginHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) ShowQuota(request *model.ShowQuotaRequest) (*model.ShowQuotaResponse, error) {
	requestDef := GenReqDefForShowQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotaResponse), nil
	}
}

func (c *CdnClient) ShowQuotaInvoker(request *model.ShowQuotaRequest) *ShowQuotaInvoker {
	requestDef := GenReqDefForShowQuota()
	return &ShowQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) ShowRefer(request *model.ShowReferRequest) (*model.ShowReferResponse, error) {
	requestDef := GenReqDefForShowRefer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReferResponse), nil
	}
}

func (c *CdnClient) ShowReferInvoker(request *model.ShowReferRequest) *ShowReferInvoker {
	requestDef := GenReqDefForShowRefer()
	return &ShowReferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) ShowResponseHeader(request *model.ShowResponseHeaderRequest) (*model.ShowResponseHeaderResponse, error) {
	requestDef := GenReqDefForShowResponseHeader()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResponseHeaderResponse), nil
	}
}

func (c *CdnClient) ShowResponseHeaderInvoker(request *model.ShowResponseHeaderRequest) *ShowResponseHeaderInvoker {
	requestDef := GenReqDefForShowResponseHeader()
	return &ShowResponseHeaderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) ShowTopUrl(request *model.ShowTopUrlRequest) (*model.ShowTopUrlResponse, error) {
	requestDef := GenReqDefForShowTopUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTopUrlResponse), nil
	}
}

func (c *CdnClient) ShowTopUrlInvoker(request *model.ShowTopUrlRequest) *ShowTopUrlInvoker {
	requestDef := GenReqDefForShowTopUrl()
	return &ShowTopUrlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) UpdateBlackWhiteList(request *model.UpdateBlackWhiteListRequest) (*model.UpdateBlackWhiteListResponse, error) {
	requestDef := GenReqDefForUpdateBlackWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBlackWhiteListResponse), nil
	}
}

func (c *CdnClient) UpdateBlackWhiteListInvoker(request *model.UpdateBlackWhiteListRequest) *UpdateBlackWhiteListInvoker {
	requestDef := GenReqDefForUpdateBlackWhiteList()
	return &UpdateBlackWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) UpdateCacheRules(request *model.UpdateCacheRulesRequest) (*model.UpdateCacheRulesResponse, error) {
	requestDef := GenReqDefForUpdateCacheRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCacheRulesResponse), nil
	}
}

func (c *CdnClient) UpdateCacheRulesInvoker(request *model.UpdateCacheRulesRequest) *UpdateCacheRulesInvoker {
	requestDef := GenReqDefForUpdateCacheRules()
	return &UpdateCacheRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) UpdateDomainFullConfig(request *model.UpdateDomainFullConfigRequest) (*model.UpdateDomainFullConfigResponse, error) {
	requestDef := GenReqDefForUpdateDomainFullConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainFullConfigResponse), nil
	}
}

func (c *CdnClient) UpdateDomainFullConfigInvoker(request *model.UpdateDomainFullConfigRequest) *UpdateDomainFullConfigInvoker {
	requestDef := GenReqDefForUpdateDomainFullConfig()
	return &UpdateDomainFullConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) UpdateDomainMultiCertificates(request *model.UpdateDomainMultiCertificatesRequest) (*model.UpdateDomainMultiCertificatesResponse, error) {
	requestDef := GenReqDefForUpdateDomainMultiCertificates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainMultiCertificatesResponse), nil
	}
}

func (c *CdnClient) UpdateDomainMultiCertificatesInvoker(request *model.UpdateDomainMultiCertificatesRequest) *UpdateDomainMultiCertificatesInvoker {
	requestDef := GenReqDefForUpdateDomainMultiCertificates()
	return &UpdateDomainMultiCertificatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) UpdateDomainOrigin(request *model.UpdateDomainOriginRequest) (*model.UpdateDomainOriginResponse, error) {
	requestDef := GenReqDefForUpdateDomainOrigin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainOriginResponse), nil
	}
}

func (c *CdnClient) UpdateDomainOriginInvoker(request *model.UpdateDomainOriginRequest) *UpdateDomainOriginInvoker {
	requestDef := GenReqDefForUpdateDomainOrigin()
	return &UpdateDomainOriginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) UpdateFollow302Switch(request *model.UpdateFollow302SwitchRequest) (*model.UpdateFollow302SwitchResponse, error) {
	requestDef := GenReqDefForUpdateFollow302Switch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFollow302SwitchResponse), nil
	}
}

func (c *CdnClient) UpdateFollow302SwitchInvoker(request *model.UpdateFollow302SwitchRequest) *UpdateFollow302SwitchInvoker {
	requestDef := GenReqDefForUpdateFollow302Switch()
	return &UpdateFollow302SwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) UpdateHttpsInfo(request *model.UpdateHttpsInfoRequest) (*model.UpdateHttpsInfoResponse, error) {
	requestDef := GenReqDefForUpdateHttpsInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHttpsInfoResponse), nil
	}
}

func (c *CdnClient) UpdateHttpsInfoInvoker(request *model.UpdateHttpsInfoRequest) *UpdateHttpsInfoInvoker {
	requestDef := GenReqDefForUpdateHttpsInfo()
	return &UpdateHttpsInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) UpdateOriginHost(request *model.UpdateOriginHostRequest) (*model.UpdateOriginHostResponse, error) {
	requestDef := GenReqDefForUpdateOriginHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateOriginHostResponse), nil
	}
}

func (c *CdnClient) UpdateOriginHostInvoker(request *model.UpdateOriginHostRequest) *UpdateOriginHostInvoker {
	requestDef := GenReqDefForUpdateOriginHost()
	return &UpdateOriginHostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) UpdatePrivateBucketAccess(request *model.UpdatePrivateBucketAccessRequest) (*model.UpdatePrivateBucketAccessResponse, error) {
	requestDef := GenReqDefForUpdatePrivateBucketAccess()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePrivateBucketAccessResponse), nil
	}
}

func (c *CdnClient) UpdatePrivateBucketAccessInvoker(request *model.UpdatePrivateBucketAccessRequest) *UpdatePrivateBucketAccessInvoker {
	requestDef := GenReqDefForUpdatePrivateBucketAccess()
	return &UpdatePrivateBucketAccessInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) UpdateRangeSwitch(request *model.UpdateRangeSwitchRequest) (*model.UpdateRangeSwitchResponse, error) {
	requestDef := GenReqDefForUpdateRangeSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRangeSwitchResponse), nil
	}
}

func (c *CdnClient) UpdateRangeSwitchInvoker(request *model.UpdateRangeSwitchRequest) *UpdateRangeSwitchInvoker {
	requestDef := GenReqDefForUpdateRangeSwitch()
	return &UpdateRangeSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) UpdateRefer(request *model.UpdateReferRequest) (*model.UpdateReferResponse, error) {
	requestDef := GenReqDefForUpdateRefer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateReferResponse), nil
	}
}

func (c *CdnClient) UpdateReferInvoker(request *model.UpdateReferRequest) *UpdateReferInvoker {
	requestDef := GenReqDefForUpdateRefer()
	return &UpdateReferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CdnClient) UpdateResponseHeader(request *model.UpdateResponseHeaderRequest) (*model.UpdateResponseHeaderResponse, error) {
	requestDef := GenReqDefForUpdateResponseHeader()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResponseHeaderResponse), nil
	}
}

func (c *CdnClient) UpdateResponseHeaderInvoker(request *model.UpdateResponseHeaderRequest) *UpdateResponseHeaderInvoker {
	requestDef := GenReqDefForUpdateResponseHeader()
	return &UpdateResponseHeaderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
