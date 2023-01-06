package v2

import (
	http_client "github.com/g42cloud-sdk/g42cloud-sdk-go/core"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/invoker"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/services/smn/v2/model"
)

type SmnClient struct {
	HcClient *http_client.HcHttpClient
}

func NewSmnClient(hcClient *http_client.HcHttpClient) *SmnClient {
	return &SmnClient{HcClient: hcClient}
}

func SmnClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *SmnClient) AddSubscription(request *model.AddSubscriptionRequest) (*model.AddSubscriptionResponse, error) {
	requestDef := GenReqDefForAddSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddSubscriptionResponse), nil
	}
}

func (c *SmnClient) AddSubscriptionInvoker(request *model.AddSubscriptionRequest) *AddSubscriptionInvoker {
	requestDef := GenReqDefForAddSubscription()
	return &AddSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) BatchCreateOrDeleteResourceTags(request *model.BatchCreateOrDeleteResourceTagsRequest) (*model.BatchCreateOrDeleteResourceTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateOrDeleteResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateOrDeleteResourceTagsResponse), nil
	}
}

func (c *SmnClient) BatchCreateOrDeleteResourceTagsInvoker(request *model.BatchCreateOrDeleteResourceTagsRequest) *BatchCreateOrDeleteResourceTagsInvoker {
	requestDef := GenReqDefForBatchCreateOrDeleteResourceTags()
	return &BatchCreateOrDeleteResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) CancelSubscription(request *model.CancelSubscriptionRequest) (*model.CancelSubscriptionResponse, error) {
	requestDef := GenReqDefForCancelSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelSubscriptionResponse), nil
	}
}

func (c *SmnClient) CancelSubscriptionInvoker(request *model.CancelSubscriptionRequest) *CancelSubscriptionInvoker {
	requestDef := GenReqDefForCancelSubscription()
	return &CancelSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) CreateMessageTemplate(request *model.CreateMessageTemplateRequest) (*model.CreateMessageTemplateResponse, error) {
	requestDef := GenReqDefForCreateMessageTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMessageTemplateResponse), nil
	}
}

func (c *SmnClient) CreateMessageTemplateInvoker(request *model.CreateMessageTemplateRequest) *CreateMessageTemplateInvoker {
	requestDef := GenReqDefForCreateMessageTemplate()
	return &CreateMessageTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) CreateResourceTag(request *model.CreateResourceTagRequest) (*model.CreateResourceTagResponse, error) {
	requestDef := GenReqDefForCreateResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResourceTagResponse), nil
	}
}

func (c *SmnClient) CreateResourceTagInvoker(request *model.CreateResourceTagRequest) *CreateResourceTagInvoker {
	requestDef := GenReqDefForCreateResourceTag()
	return &CreateResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) CreateTopic(request *model.CreateTopicRequest) (*model.CreateTopicResponse, error) {
	requestDef := GenReqDefForCreateTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTopicResponse), nil
	}
}

func (c *SmnClient) CreateTopicInvoker(request *model.CreateTopicRequest) *CreateTopicInvoker {
	requestDef := GenReqDefForCreateTopic()
	return &CreateTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) DeleteMessageTemplate(request *model.DeleteMessageTemplateRequest) (*model.DeleteMessageTemplateResponse, error) {
	requestDef := GenReqDefForDeleteMessageTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMessageTemplateResponse), nil
	}
}

func (c *SmnClient) DeleteMessageTemplateInvoker(request *model.DeleteMessageTemplateRequest) *DeleteMessageTemplateInvoker {
	requestDef := GenReqDefForDeleteMessageTemplate()
	return &DeleteMessageTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) DeleteResourceTag(request *model.DeleteResourceTagRequest) (*model.DeleteResourceTagResponse, error) {
	requestDef := GenReqDefForDeleteResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceTagResponse), nil
	}
}

func (c *SmnClient) DeleteResourceTagInvoker(request *model.DeleteResourceTagRequest) *DeleteResourceTagInvoker {
	requestDef := GenReqDefForDeleteResourceTag()
	return &DeleteResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) DeleteTopic(request *model.DeleteTopicRequest) (*model.DeleteTopicResponse, error) {
	requestDef := GenReqDefForDeleteTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTopicResponse), nil
	}
}

func (c *SmnClient) DeleteTopicInvoker(request *model.DeleteTopicRequest) *DeleteTopicInvoker {
	requestDef := GenReqDefForDeleteTopic()
	return &DeleteTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) DeleteTopicAttributeByName(request *model.DeleteTopicAttributeByNameRequest) (*model.DeleteTopicAttributeByNameResponse, error) {
	requestDef := GenReqDefForDeleteTopicAttributeByName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTopicAttributeByNameResponse), nil
	}
}

func (c *SmnClient) DeleteTopicAttributeByNameInvoker(request *model.DeleteTopicAttributeByNameRequest) *DeleteTopicAttributeByNameInvoker {
	requestDef := GenReqDefForDeleteTopicAttributeByName()
	return &DeleteTopicAttributeByNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) DeleteTopicAttributes(request *model.DeleteTopicAttributesRequest) (*model.DeleteTopicAttributesResponse, error) {
	requestDef := GenReqDefForDeleteTopicAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTopicAttributesResponse), nil
	}
}

func (c *SmnClient) DeleteTopicAttributesInvoker(request *model.DeleteTopicAttributesRequest) *DeleteTopicAttributesInvoker {
	requestDef := GenReqDefForDeleteTopicAttributes()
	return &DeleteTopicAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) ListMessageTemplateDetails(request *model.ListMessageTemplateDetailsRequest) (*model.ListMessageTemplateDetailsResponse, error) {
	requestDef := GenReqDefForListMessageTemplateDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMessageTemplateDetailsResponse), nil
	}
}

func (c *SmnClient) ListMessageTemplateDetailsInvoker(request *model.ListMessageTemplateDetailsRequest) *ListMessageTemplateDetailsInvoker {
	requestDef := GenReqDefForListMessageTemplateDetails()
	return &ListMessageTemplateDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) ListMessageTemplates(request *model.ListMessageTemplatesRequest) (*model.ListMessageTemplatesResponse, error) {
	requestDef := GenReqDefForListMessageTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMessageTemplatesResponse), nil
	}
}

func (c *SmnClient) ListMessageTemplatesInvoker(request *model.ListMessageTemplatesRequest) *ListMessageTemplatesInvoker {
	requestDef := GenReqDefForListMessageTemplates()
	return &ListMessageTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) ListProjectTags(request *model.ListProjectTagsRequest) (*model.ListProjectTagsResponse, error) {
	requestDef := GenReqDefForListProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTagsResponse), nil
	}
}

func (c *SmnClient) ListProjectTagsInvoker(request *model.ListProjectTagsRequest) *ListProjectTagsInvoker {
	requestDef := GenReqDefForListProjectTags()
	return &ListProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) ListResourceInstances(request *model.ListResourceInstancesRequest) (*model.ListResourceInstancesResponse, error) {
	requestDef := GenReqDefForListResourceInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceInstancesResponse), nil
	}
}

func (c *SmnClient) ListResourceInstancesInvoker(request *model.ListResourceInstancesRequest) *ListResourceInstancesInvoker {
	requestDef := GenReqDefForListResourceInstances()
	return &ListResourceInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) ListResourceTags(request *model.ListResourceTagsRequest) (*model.ListResourceTagsResponse, error) {
	requestDef := GenReqDefForListResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceTagsResponse), nil
	}
}

func (c *SmnClient) ListResourceTagsInvoker(request *model.ListResourceTagsRequest) *ListResourceTagsInvoker {
	requestDef := GenReqDefForListResourceTags()
	return &ListResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) ListSubscriptions(request *model.ListSubscriptionsRequest) (*model.ListSubscriptionsResponse, error) {
	requestDef := GenReqDefForListSubscriptions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubscriptionsResponse), nil
	}
}

func (c *SmnClient) ListSubscriptionsInvoker(request *model.ListSubscriptionsRequest) *ListSubscriptionsInvoker {
	requestDef := GenReqDefForListSubscriptions()
	return &ListSubscriptionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) ListSubscriptionsByTopic(request *model.ListSubscriptionsByTopicRequest) (*model.ListSubscriptionsByTopicResponse, error) {
	requestDef := GenReqDefForListSubscriptionsByTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubscriptionsByTopicResponse), nil
	}
}

func (c *SmnClient) ListSubscriptionsByTopicInvoker(request *model.ListSubscriptionsByTopicRequest) *ListSubscriptionsByTopicInvoker {
	requestDef := GenReqDefForListSubscriptionsByTopic()
	return &ListSubscriptionsByTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) ListTopicAttributes(request *model.ListTopicAttributesRequest) (*model.ListTopicAttributesResponse, error) {
	requestDef := GenReqDefForListTopicAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTopicAttributesResponse), nil
	}
}

func (c *SmnClient) ListTopicAttributesInvoker(request *model.ListTopicAttributesRequest) *ListTopicAttributesInvoker {
	requestDef := GenReqDefForListTopicAttributes()
	return &ListTopicAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) ListTopicDetails(request *model.ListTopicDetailsRequest) (*model.ListTopicDetailsResponse, error) {
	requestDef := GenReqDefForListTopicDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTopicDetailsResponse), nil
	}
}

func (c *SmnClient) ListTopicDetailsInvoker(request *model.ListTopicDetailsRequest) *ListTopicDetailsInvoker {
	requestDef := GenReqDefForListTopicDetails()
	return &ListTopicDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) ListTopics(request *model.ListTopicsRequest) (*model.ListTopicsResponse, error) {
	requestDef := GenReqDefForListTopics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTopicsResponse), nil
	}
}

func (c *SmnClient) ListTopicsInvoker(request *model.ListTopicsRequest) *ListTopicsInvoker {
	requestDef := GenReqDefForListTopics()
	return &ListTopicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) ListVersion(request *model.ListVersionRequest) (*model.ListVersionResponse, error) {
	requestDef := GenReqDefForListVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVersionResponse), nil
	}
}

func (c *SmnClient) ListVersionInvoker(request *model.ListVersionRequest) *ListVersionInvoker {
	requestDef := GenReqDefForListVersion()
	return &ListVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) ListVersions(request *model.ListVersionsRequest) (*model.ListVersionsResponse, error) {
	requestDef := GenReqDefForListVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVersionsResponse), nil
	}
}

func (c *SmnClient) ListVersionsInvoker(request *model.ListVersionsRequest) *ListVersionsInvoker {
	requestDef := GenReqDefForListVersions()
	return &ListVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) PublishMessage(request *model.PublishMessageRequest) (*model.PublishMessageResponse, error) {
	requestDef := GenReqDefForPublishMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PublishMessageResponse), nil
	}
}

func (c *SmnClient) PublishMessageInvoker(request *model.PublishMessageRequest) *PublishMessageInvoker {
	requestDef := GenReqDefForPublishMessage()
	return &PublishMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) UpdateMessageTemplate(request *model.UpdateMessageTemplateRequest) (*model.UpdateMessageTemplateResponse, error) {
	requestDef := GenReqDefForUpdateMessageTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMessageTemplateResponse), nil
	}
}

func (c *SmnClient) UpdateMessageTemplateInvoker(request *model.UpdateMessageTemplateRequest) *UpdateMessageTemplateInvoker {
	requestDef := GenReqDefForUpdateMessageTemplate()
	return &UpdateMessageTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) UpdateTopic(request *model.UpdateTopicRequest) (*model.UpdateTopicResponse, error) {
	requestDef := GenReqDefForUpdateTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTopicResponse), nil
	}
}

func (c *SmnClient) UpdateTopicInvoker(request *model.UpdateTopicRequest) *UpdateTopicInvoker {
	requestDef := GenReqDefForUpdateTopic()
	return &UpdateTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) UpdateTopicAttribute(request *model.UpdateTopicAttributeRequest) (*model.UpdateTopicAttributeResponse, error) {
	requestDef := GenReqDefForUpdateTopicAttribute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTopicAttributeResponse), nil
	}
}

func (c *SmnClient) UpdateTopicAttributeInvoker(request *model.UpdateTopicAttributeRequest) *UpdateTopicAttributeInvoker {
	requestDef := GenReqDefForUpdateTopicAttribute()
	return &UpdateTopicAttributeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) CreateApplication(request *model.CreateApplicationRequest) (*model.CreateApplicationResponse, error) {
	requestDef := GenReqDefForCreateApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApplicationResponse), nil
	}
}

func (c *SmnClient) CreateApplicationInvoker(request *model.CreateApplicationRequest) *CreateApplicationInvoker {
	requestDef := GenReqDefForCreateApplication()
	return &CreateApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) DeleteApplication(request *model.DeleteApplicationRequest) (*model.DeleteApplicationResponse, error) {
	requestDef := GenReqDefForDeleteApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApplicationResponse), nil
	}
}

func (c *SmnClient) DeleteApplicationInvoker(request *model.DeleteApplicationRequest) *DeleteApplicationInvoker {
	requestDef := GenReqDefForDeleteApplication()
	return &DeleteApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) ListApplicationAttributes(request *model.ListApplicationAttributesRequest) (*model.ListApplicationAttributesResponse, error) {
	requestDef := GenReqDefForListApplicationAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationAttributesResponse), nil
	}
}

func (c *SmnClient) ListApplicationAttributesInvoker(request *model.ListApplicationAttributesRequest) *ListApplicationAttributesInvoker {
	requestDef := GenReqDefForListApplicationAttributes()
	return &ListApplicationAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) ListApplications(request *model.ListApplicationsRequest) (*model.ListApplicationsResponse, error) {
	requestDef := GenReqDefForListApplications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationsResponse), nil
	}
}

func (c *SmnClient) ListApplicationsInvoker(request *model.ListApplicationsRequest) *ListApplicationsInvoker {
	requestDef := GenReqDefForListApplications()
	return &ListApplicationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) PublishAppMessage(request *model.PublishAppMessageRequest) (*model.PublishAppMessageResponse, error) {
	requestDef := GenReqDefForPublishAppMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PublishAppMessageResponse), nil
	}
}

func (c *SmnClient) PublishAppMessageInvoker(request *model.PublishAppMessageRequest) *PublishAppMessageInvoker {
	requestDef := GenReqDefForPublishAppMessage()
	return &PublishAppMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) UpdateApplication(request *model.UpdateApplicationRequest) (*model.UpdateApplicationResponse, error) {
	requestDef := GenReqDefForUpdateApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApplicationResponse), nil
	}
}

func (c *SmnClient) UpdateApplicationInvoker(request *model.UpdateApplicationRequest) *UpdateApplicationInvoker {
	requestDef := GenReqDefForUpdateApplication()
	return &UpdateApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) CreateApplicationEndpoint(request *model.CreateApplicationEndpointRequest) (*model.CreateApplicationEndpointResponse, error) {
	requestDef := GenReqDefForCreateApplicationEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApplicationEndpointResponse), nil
	}
}

func (c *SmnClient) CreateApplicationEndpointInvoker(request *model.CreateApplicationEndpointRequest) *CreateApplicationEndpointInvoker {
	requestDef := GenReqDefForCreateApplicationEndpoint()
	return &CreateApplicationEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) DeleteApplicationEndpoint(request *model.DeleteApplicationEndpointRequest) (*model.DeleteApplicationEndpointResponse, error) {
	requestDef := GenReqDefForDeleteApplicationEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApplicationEndpointResponse), nil
	}
}

func (c *SmnClient) DeleteApplicationEndpointInvoker(request *model.DeleteApplicationEndpointRequest) *DeleteApplicationEndpointInvoker {
	requestDef := GenReqDefForDeleteApplicationEndpoint()
	return &DeleteApplicationEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) ListApplicationEndpointAttributes(request *model.ListApplicationEndpointAttributesRequest) (*model.ListApplicationEndpointAttributesResponse, error) {
	requestDef := GenReqDefForListApplicationEndpointAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationEndpointAttributesResponse), nil
	}
}

func (c *SmnClient) ListApplicationEndpointAttributesInvoker(request *model.ListApplicationEndpointAttributesRequest) *ListApplicationEndpointAttributesInvoker {
	requestDef := GenReqDefForListApplicationEndpointAttributes()
	return &ListApplicationEndpointAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) ListApplicationEndpoints(request *model.ListApplicationEndpointsRequest) (*model.ListApplicationEndpointsResponse, error) {
	requestDef := GenReqDefForListApplicationEndpoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationEndpointsResponse), nil
	}
}

func (c *SmnClient) ListApplicationEndpointsInvoker(request *model.ListApplicationEndpointsRequest) *ListApplicationEndpointsInvoker {
	requestDef := GenReqDefForListApplicationEndpoints()
	return &ListApplicationEndpointsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmnClient) UpdateApplicationEndpoint(request *model.UpdateApplicationEndpointRequest) (*model.UpdateApplicationEndpointResponse, error) {
	requestDef := GenReqDefForUpdateApplicationEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApplicationEndpointResponse), nil
	}
}

func (c *SmnClient) UpdateApplicationEndpointInvoker(request *model.UpdateApplicationEndpointRequest) *UpdateApplicationEndpointInvoker {
	requestDef := GenReqDefForUpdateApplicationEndpoint()
	return &UpdateApplicationEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
