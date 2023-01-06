package v3

import (
	http_client "github.com/g42cloud-sdk/g42cloud-sdk-go/core"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/invoker"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/services/cts/v3/model"
)

type CtsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCtsClient(hcClient *http_client.HcHttpClient) *CtsClient {
	return &CtsClient{HcClient: hcClient}
}

func CtsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *CtsClient) CreateNotification(request *model.CreateNotificationRequest) (*model.CreateNotificationResponse, error) {
	requestDef := GenReqDefForCreateNotification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNotificationResponse), nil
	}
}

func (c *CtsClient) CreateNotificationInvoker(request *model.CreateNotificationRequest) *CreateNotificationInvoker {
	requestDef := GenReqDefForCreateNotification()
	return &CreateNotificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CtsClient) CreateTracker(request *model.CreateTrackerRequest) (*model.CreateTrackerResponse, error) {
	requestDef := GenReqDefForCreateTracker()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTrackerResponse), nil
	}
}

func (c *CtsClient) CreateTrackerInvoker(request *model.CreateTrackerRequest) *CreateTrackerInvoker {
	requestDef := GenReqDefForCreateTracker()
	return &CreateTrackerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CtsClient) DeleteNotification(request *model.DeleteNotificationRequest) (*model.DeleteNotificationResponse, error) {
	requestDef := GenReqDefForDeleteNotification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNotificationResponse), nil
	}
}

func (c *CtsClient) DeleteNotificationInvoker(request *model.DeleteNotificationRequest) *DeleteNotificationInvoker {
	requestDef := GenReqDefForDeleteNotification()
	return &DeleteNotificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CtsClient) DeleteTracker(request *model.DeleteTrackerRequest) (*model.DeleteTrackerResponse, error) {
	requestDef := GenReqDefForDeleteTracker()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTrackerResponse), nil
	}
}

func (c *CtsClient) DeleteTrackerInvoker(request *model.DeleteTrackerRequest) *DeleteTrackerInvoker {
	requestDef := GenReqDefForDeleteTracker()
	return &DeleteTrackerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CtsClient) ListNotifications(request *model.ListNotificationsRequest) (*model.ListNotificationsResponse, error) {
	requestDef := GenReqDefForListNotifications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNotificationsResponse), nil
	}
}

func (c *CtsClient) ListNotificationsInvoker(request *model.ListNotificationsRequest) *ListNotificationsInvoker {
	requestDef := GenReqDefForListNotifications()
	return &ListNotificationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CtsClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasResponse), nil
	}
}

func (c *CtsClient) ListQuotasInvoker(request *model.ListQuotasRequest) *ListQuotasInvoker {
	requestDef := GenReqDefForListQuotas()
	return &ListQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CtsClient) ListTraces(request *model.ListTracesRequest) (*model.ListTracesResponse, error) {
	requestDef := GenReqDefForListTraces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTracesResponse), nil
	}
}

func (c *CtsClient) ListTracesInvoker(request *model.ListTracesRequest) *ListTracesInvoker {
	requestDef := GenReqDefForListTraces()
	return &ListTracesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CtsClient) ListTrackers(request *model.ListTrackersRequest) (*model.ListTrackersResponse, error) {
	requestDef := GenReqDefForListTrackers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTrackersResponse), nil
	}
}

func (c *CtsClient) ListTrackersInvoker(request *model.ListTrackersRequest) *ListTrackersInvoker {
	requestDef := GenReqDefForListTrackers()
	return &ListTrackersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CtsClient) UpdateNotification(request *model.UpdateNotificationRequest) (*model.UpdateNotificationResponse, error) {
	requestDef := GenReqDefForUpdateNotification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNotificationResponse), nil
	}
}

func (c *CtsClient) UpdateNotificationInvoker(request *model.UpdateNotificationRequest) *UpdateNotificationInvoker {
	requestDef := GenReqDefForUpdateNotification()
	return &UpdateNotificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CtsClient) UpdateTracker(request *model.UpdateTrackerRequest) (*model.UpdateTrackerResponse, error) {
	requestDef := GenReqDefForUpdateTracker()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTrackerResponse), nil
	}
}

func (c *CtsClient) UpdateTrackerInvoker(request *model.UpdateTrackerRequest) *UpdateTrackerInvoker {
	requestDef := GenReqDefForUpdateTracker()
	return &UpdateTrackerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
