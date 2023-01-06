package v2

import (
	http_client "github.com/g42cloud-sdk/g42cloud-sdk-go/core"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/invoker"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/services/ces/v2/model"
)

type CesClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCesClient(hcClient *http_client.HcHttpClient) *CesClient {
	return &CesClient{HcClient: hcClient}
}

func CesClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *CesClient) AddAlarmRuleResources(request *model.AddAlarmRuleResourcesRequest) (*model.AddAlarmRuleResourcesResponse, error) {
	requestDef := GenReqDefForAddAlarmRuleResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAlarmRuleResourcesResponse), nil
	}
}

func (c *CesClient) AddAlarmRuleResourcesInvoker(request *model.AddAlarmRuleResourcesRequest) *AddAlarmRuleResourcesInvoker {
	requestDef := GenReqDefForAddAlarmRuleResources()
	return &AddAlarmRuleResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) BatchDeleteAlarmRules(request *model.BatchDeleteAlarmRulesRequest) (*model.BatchDeleteAlarmRulesResponse, error) {
	requestDef := GenReqDefForBatchDeleteAlarmRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAlarmRulesResponse), nil
	}
}

func (c *CesClient) BatchDeleteAlarmRulesInvoker(request *model.BatchDeleteAlarmRulesRequest) *BatchDeleteAlarmRulesInvoker {
	requestDef := GenReqDefForBatchDeleteAlarmRules()
	return &BatchDeleteAlarmRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) BatchEnableAlarmRules(request *model.BatchEnableAlarmRulesRequest) (*model.BatchEnableAlarmRulesResponse, error) {
	requestDef := GenReqDefForBatchEnableAlarmRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchEnableAlarmRulesResponse), nil
	}
}

func (c *CesClient) BatchEnableAlarmRulesInvoker(request *model.BatchEnableAlarmRulesRequest) *BatchEnableAlarmRulesInvoker {
	requestDef := GenReqDefForBatchEnableAlarmRules()
	return &BatchEnableAlarmRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) CreateAlarmRules(request *model.CreateAlarmRulesRequest) (*model.CreateAlarmRulesResponse, error) {
	requestDef := GenReqDefForCreateAlarmRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlarmRulesResponse), nil
	}
}

func (c *CesClient) CreateAlarmRulesInvoker(request *model.CreateAlarmRulesRequest) *CreateAlarmRulesInvoker {
	requestDef := GenReqDefForCreateAlarmRules()
	return &CreateAlarmRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) DeleteAlarmRuleResources(request *model.DeleteAlarmRuleResourcesRequest) (*model.DeleteAlarmRuleResourcesResponse, error) {
	requestDef := GenReqDefForDeleteAlarmRuleResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlarmRuleResourcesResponse), nil
	}
}

func (c *CesClient) DeleteAlarmRuleResourcesInvoker(request *model.DeleteAlarmRuleResourcesRequest) *DeleteAlarmRuleResourcesInvoker {
	requestDef := GenReqDefForDeleteAlarmRuleResources()
	return &DeleteAlarmRuleResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) ListAgentDimensionInfo(request *model.ListAgentDimensionInfoRequest) (*model.ListAgentDimensionInfoResponse, error) {
	requestDef := GenReqDefForListAgentDimensionInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAgentDimensionInfoResponse), nil
	}
}

func (c *CesClient) ListAgentDimensionInfoInvoker(request *model.ListAgentDimensionInfoRequest) *ListAgentDimensionInfoInvoker {
	requestDef := GenReqDefForListAgentDimensionInfo()
	return &ListAgentDimensionInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) ListAlarmHistories(request *model.ListAlarmHistoriesRequest) (*model.ListAlarmHistoriesResponse, error) {
	requestDef := GenReqDefForListAlarmHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmHistoriesResponse), nil
	}
}

func (c *CesClient) ListAlarmHistoriesInvoker(request *model.ListAlarmHistoriesRequest) *ListAlarmHistoriesInvoker {
	requestDef := GenReqDefForListAlarmHistories()
	return &ListAlarmHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) ListAlarmRulePolicies(request *model.ListAlarmRulePoliciesRequest) (*model.ListAlarmRulePoliciesResponse, error) {
	requestDef := GenReqDefForListAlarmRulePolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmRulePoliciesResponse), nil
	}
}

func (c *CesClient) ListAlarmRulePoliciesInvoker(request *model.ListAlarmRulePoliciesRequest) *ListAlarmRulePoliciesInvoker {
	requestDef := GenReqDefForListAlarmRulePolicies()
	return &ListAlarmRulePoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) ListAlarmRuleResources(request *model.ListAlarmRuleResourcesRequest) (*model.ListAlarmRuleResourcesResponse, error) {
	requestDef := GenReqDefForListAlarmRuleResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmRuleResourcesResponse), nil
	}
}

func (c *CesClient) ListAlarmRuleResourcesInvoker(request *model.ListAlarmRuleResourcesRequest) *ListAlarmRuleResourcesInvoker {
	requestDef := GenReqDefForListAlarmRuleResources()
	return &ListAlarmRuleResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) ListAlarmRules(request *model.ListAlarmRulesRequest) (*model.ListAlarmRulesResponse, error) {
	requestDef := GenReqDefForListAlarmRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmRulesResponse), nil
	}
}

func (c *CesClient) ListAlarmRulesInvoker(request *model.ListAlarmRulesRequest) *ListAlarmRulesInvoker {
	requestDef := GenReqDefForListAlarmRules()
	return &ListAlarmRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) UpdateAlarmRulePolicies(request *model.UpdateAlarmRulePoliciesRequest) (*model.UpdateAlarmRulePoliciesResponse, error) {
	requestDef := GenReqDefForUpdateAlarmRulePolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlarmRulePoliciesResponse), nil
	}
}

func (c *CesClient) UpdateAlarmRulePoliciesInvoker(request *model.UpdateAlarmRulePoliciesRequest) *UpdateAlarmRulePoliciesInvoker {
	requestDef := GenReqDefForUpdateAlarmRulePolicies()
	return &UpdateAlarmRulePoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
