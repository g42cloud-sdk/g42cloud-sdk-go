package v1

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/ces/v1/model"
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
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

func (c *CesClient) BatchListMetricData(request *model.BatchListMetricDataRequest) (*model.BatchListMetricDataResponse, error) {
	requestDef := GenReqDefForBatchListMetricData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListMetricDataResponse), nil
	}
}

func (c *CesClient) BatchListMetricDataInvoker(request *model.BatchListMetricDataRequest) *BatchListMetricDataInvoker {
	requestDef := GenReqDefForBatchListMetricData()
	return &BatchListMetricDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) CreateAlarm(request *model.CreateAlarmRequest) (*model.CreateAlarmResponse, error) {
	requestDef := GenReqDefForCreateAlarm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlarmResponse), nil
	}
}

func (c *CesClient) CreateAlarmInvoker(request *model.CreateAlarmRequest) *CreateAlarmInvoker {
	requestDef := GenReqDefForCreateAlarm()
	return &CreateAlarmInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) CreateAlarmTemplate(request *model.CreateAlarmTemplateRequest) (*model.CreateAlarmTemplateResponse, error) {
	requestDef := GenReqDefForCreateAlarmTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlarmTemplateResponse), nil
	}
}

func (c *CesClient) CreateAlarmTemplateInvoker(request *model.CreateAlarmTemplateRequest) *CreateAlarmTemplateInvoker {
	requestDef := GenReqDefForCreateAlarmTemplate()
	return &CreateAlarmTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) CreateEvents(request *model.CreateEventsRequest) (*model.CreateEventsResponse, error) {
	requestDef := GenReqDefForCreateEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEventsResponse), nil
	}
}

func (c *CesClient) CreateEventsInvoker(request *model.CreateEventsRequest) *CreateEventsInvoker {
	requestDef := GenReqDefForCreateEvents()
	return &CreateEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) CreateMetricData(request *model.CreateMetricDataRequest) (*model.CreateMetricDataResponse, error) {
	requestDef := GenReqDefForCreateMetricData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMetricDataResponse), nil
	}
}

func (c *CesClient) CreateMetricDataInvoker(request *model.CreateMetricDataRequest) *CreateMetricDataInvoker {
	requestDef := GenReqDefForCreateMetricData()
	return &CreateMetricDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) CreateResourceGroup(request *model.CreateResourceGroupRequest) (*model.CreateResourceGroupResponse, error) {
	requestDef := GenReqDefForCreateResourceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResourceGroupResponse), nil
	}
}

func (c *CesClient) CreateResourceGroupInvoker(request *model.CreateResourceGroupRequest) *CreateResourceGroupInvoker {
	requestDef := GenReqDefForCreateResourceGroup()
	return &CreateResourceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) DeleteAlarm(request *model.DeleteAlarmRequest) (*model.DeleteAlarmResponse, error) {
	requestDef := GenReqDefForDeleteAlarm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlarmResponse), nil
	}
}

func (c *CesClient) DeleteAlarmInvoker(request *model.DeleteAlarmRequest) *DeleteAlarmInvoker {
	requestDef := GenReqDefForDeleteAlarm()
	return &DeleteAlarmInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) DeleteAlarmTemplate(request *model.DeleteAlarmTemplateRequest) (*model.DeleteAlarmTemplateResponse, error) {
	requestDef := GenReqDefForDeleteAlarmTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlarmTemplateResponse), nil
	}
}

func (c *CesClient) DeleteAlarmTemplateInvoker(request *model.DeleteAlarmTemplateRequest) *DeleteAlarmTemplateInvoker {
	requestDef := GenReqDefForDeleteAlarmTemplate()
	return &DeleteAlarmTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) DeleteResourceGroup(request *model.DeleteResourceGroupRequest) (*model.DeleteResourceGroupResponse, error) {
	requestDef := GenReqDefForDeleteResourceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceGroupResponse), nil
	}
}

func (c *CesClient) DeleteResourceGroupInvoker(request *model.DeleteResourceGroupRequest) *DeleteResourceGroupInvoker {
	requestDef := GenReqDefForDeleteResourceGroup()
	return &DeleteResourceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

func (c *CesClient) ListAlarmTemplates(request *model.ListAlarmTemplatesRequest) (*model.ListAlarmTemplatesResponse, error) {
	requestDef := GenReqDefForListAlarmTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmTemplatesResponse), nil
	}
}

func (c *CesClient) ListAlarmTemplatesInvoker(request *model.ListAlarmTemplatesRequest) *ListAlarmTemplatesInvoker {
	requestDef := GenReqDefForListAlarmTemplates()
	return &ListAlarmTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) ListAlarms(request *model.ListAlarmsRequest) (*model.ListAlarmsResponse, error) {
	requestDef := GenReqDefForListAlarms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmsResponse), nil
	}
}

func (c *CesClient) ListAlarmsInvoker(request *model.ListAlarmsRequest) *ListAlarmsInvoker {
	requestDef := GenReqDefForListAlarms()
	return &ListAlarmsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) ListEventDetail(request *model.ListEventDetailRequest) (*model.ListEventDetailResponse, error) {
	requestDef := GenReqDefForListEventDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventDetailResponse), nil
	}
}

func (c *CesClient) ListEventDetailInvoker(request *model.ListEventDetailRequest) *ListEventDetailInvoker {
	requestDef := GenReqDefForListEventDetail()
	return &ListEventDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) ListEvents(request *model.ListEventsRequest) (*model.ListEventsResponse, error) {
	requestDef := GenReqDefForListEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventsResponse), nil
	}
}

func (c *CesClient) ListEventsInvoker(request *model.ListEventsRequest) *ListEventsInvoker {
	requestDef := GenReqDefForListEvents()
	return &ListEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) ListMetrics(request *model.ListMetricsRequest) (*model.ListMetricsResponse, error) {
	requestDef := GenReqDefForListMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMetricsResponse), nil
	}
}

func (c *CesClient) ListMetricsInvoker(request *model.ListMetricsRequest) *ListMetricsInvoker {
	requestDef := GenReqDefForListMetrics()
	return &ListMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) ListResourceGroup(request *model.ListResourceGroupRequest) (*model.ListResourceGroupResponse, error) {
	requestDef := GenReqDefForListResourceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceGroupResponse), nil
	}
}

func (c *CesClient) ListResourceGroupInvoker(request *model.ListResourceGroupRequest) *ListResourceGroupInvoker {
	requestDef := GenReqDefForListResourceGroup()
	return &ListResourceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) ShowAlarm(request *model.ShowAlarmRequest) (*model.ShowAlarmResponse, error) {
	requestDef := GenReqDefForShowAlarm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlarmResponse), nil
	}
}

func (c *CesClient) ShowAlarmInvoker(request *model.ShowAlarmRequest) *ShowAlarmInvoker {
	requestDef := GenReqDefForShowAlarm()
	return &ShowAlarmInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) ShowEventData(request *model.ShowEventDataRequest) (*model.ShowEventDataResponse, error) {
	requestDef := GenReqDefForShowEventData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEventDataResponse), nil
	}
}

func (c *CesClient) ShowEventDataInvoker(request *model.ShowEventDataRequest) *ShowEventDataInvoker {
	requestDef := GenReqDefForShowEventData()
	return &ShowEventDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) ShowMetricData(request *model.ShowMetricDataRequest) (*model.ShowMetricDataResponse, error) {
	requestDef := GenReqDefForShowMetricData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMetricDataResponse), nil
	}
}

func (c *CesClient) ShowMetricDataInvoker(request *model.ShowMetricDataRequest) *ShowMetricDataInvoker {
	requestDef := GenReqDefForShowMetricData()
	return &ShowMetricDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) ShowQuotas(request *model.ShowQuotasRequest) (*model.ShowQuotasResponse, error) {
	requestDef := GenReqDefForShowQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotasResponse), nil
	}
}

func (c *CesClient) ShowQuotasInvoker(request *model.ShowQuotasRequest) *ShowQuotasInvoker {
	requestDef := GenReqDefForShowQuotas()
	return &ShowQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) ShowResourceGroup(request *model.ShowResourceGroupRequest) (*model.ShowResourceGroupResponse, error) {
	requestDef := GenReqDefForShowResourceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceGroupResponse), nil
	}
}

func (c *CesClient) ShowResourceGroupInvoker(request *model.ShowResourceGroupRequest) *ShowResourceGroupInvoker {
	requestDef := GenReqDefForShowResourceGroup()
	return &ShowResourceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) UpdateAlarm(request *model.UpdateAlarmRequest) (*model.UpdateAlarmResponse, error) {
	requestDef := GenReqDefForUpdateAlarm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlarmResponse), nil
	}
}

func (c *CesClient) UpdateAlarmInvoker(request *model.UpdateAlarmRequest) *UpdateAlarmInvoker {
	requestDef := GenReqDefForUpdateAlarm()
	return &UpdateAlarmInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) UpdateAlarmAction(request *model.UpdateAlarmActionRequest) (*model.UpdateAlarmActionResponse, error) {
	requestDef := GenReqDefForUpdateAlarmAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlarmActionResponse), nil
	}
}

func (c *CesClient) UpdateAlarmActionInvoker(request *model.UpdateAlarmActionRequest) *UpdateAlarmActionInvoker {
	requestDef := GenReqDefForUpdateAlarmAction()
	return &UpdateAlarmActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) UpdateAlarmTemplate(request *model.UpdateAlarmTemplateRequest) (*model.UpdateAlarmTemplateResponse, error) {
	requestDef := GenReqDefForUpdateAlarmTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlarmTemplateResponse), nil
	}
}

func (c *CesClient) UpdateAlarmTemplateInvoker(request *model.UpdateAlarmTemplateRequest) *UpdateAlarmTemplateInvoker {
	requestDef := GenReqDefForUpdateAlarmTemplate()
	return &UpdateAlarmTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CesClient) UpdateResourceGroup(request *model.UpdateResourceGroupRequest) (*model.UpdateResourceGroupResponse, error) {
	requestDef := GenReqDefForUpdateResourceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResourceGroupResponse), nil
	}
}

func (c *CesClient) UpdateResourceGroupInvoker(request *model.UpdateResourceGroupRequest) *UpdateResourceGroupInvoker {
	requestDef := GenReqDefForUpdateResourceGroup()
	return &UpdateResourceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
