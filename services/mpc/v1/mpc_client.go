package v1

import (
	http_client "github.com/g42cloud-sdk/g42cloud-sdk-go/core"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/invoker"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/services/mpc/v1/model"
)

type MpcClient struct {
	HcClient *http_client.HcHttpClient
}

func NewMpcClient(hcClient *http_client.HcHttpClient) *MpcClient {
	return &MpcClient{HcClient: hcClient}
}

func MpcClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *MpcClient) CreateAnimatedGraphicsTask(request *model.CreateAnimatedGraphicsTaskRequest) (*model.CreateAnimatedGraphicsTaskResponse, error) {
	requestDef := GenReqDefForCreateAnimatedGraphicsTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAnimatedGraphicsTaskResponse), nil
	}
}

func (c *MpcClient) CreateAnimatedGraphicsTaskInvoker(request *model.CreateAnimatedGraphicsTaskRequest) *CreateAnimatedGraphicsTaskInvoker {
	requestDef := GenReqDefForCreateAnimatedGraphicsTask()
	return &CreateAnimatedGraphicsTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) DeleteAnimatedGraphicsTask(request *model.DeleteAnimatedGraphicsTaskRequest) (*model.DeleteAnimatedGraphicsTaskResponse, error) {
	requestDef := GenReqDefForDeleteAnimatedGraphicsTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAnimatedGraphicsTaskResponse), nil
	}
}

func (c *MpcClient) DeleteAnimatedGraphicsTaskInvoker(request *model.DeleteAnimatedGraphicsTaskRequest) *DeleteAnimatedGraphicsTaskInvoker {
	requestDef := GenReqDefForDeleteAnimatedGraphicsTask()
	return &DeleteAnimatedGraphicsTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ListAnimatedGraphicsTask(request *model.ListAnimatedGraphicsTaskRequest) (*model.ListAnimatedGraphicsTaskResponse, error) {
	requestDef := GenReqDefForListAnimatedGraphicsTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAnimatedGraphicsTaskResponse), nil
	}
}

func (c *MpcClient) ListAnimatedGraphicsTaskInvoker(request *model.ListAnimatedGraphicsTaskRequest) *ListAnimatedGraphicsTaskInvoker {
	requestDef := GenReqDefForListAnimatedGraphicsTask()
	return &ListAnimatedGraphicsTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) CreateAgenciesTask(request *model.CreateAgenciesTaskRequest) (*model.CreateAgenciesTaskResponse, error) {
	requestDef := GenReqDefForCreateAgenciesTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAgenciesTaskResponse), nil
	}
}

func (c *MpcClient) CreateAgenciesTaskInvoker(request *model.CreateAgenciesTaskRequest) *CreateAgenciesTaskInvoker {
	requestDef := GenReqDefForCreateAgenciesTask()
	return &CreateAgenciesTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ListAllBuckets(request *model.ListAllBucketsRequest) (*model.ListAllBucketsResponse, error) {
	requestDef := GenReqDefForListAllBuckets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllBucketsResponse), nil
	}
}

func (c *MpcClient) ListAllBucketsInvoker(request *model.ListAllBucketsRequest) *ListAllBucketsInvoker {
	requestDef := GenReqDefForListAllBuckets()
	return &ListAllBucketsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ListAllObsObjList(request *model.ListAllObsObjListRequest) (*model.ListAllObsObjListResponse, error) {
	requestDef := GenReqDefForListAllObsObjList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllObsObjListResponse), nil
	}
}

func (c *MpcClient) ListAllObsObjListInvoker(request *model.ListAllObsObjListRequest) *ListAllObsObjListInvoker {
	requestDef := GenReqDefForListAllObsObjList()
	return &ListAllObsObjListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ListNotifyEvent(request *model.ListNotifyEventRequest) (*model.ListNotifyEventResponse, error) {
	requestDef := GenReqDefForListNotifyEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNotifyEventResponse), nil
	}
}

func (c *MpcClient) ListNotifyEventInvoker(request *model.ListNotifyEventRequest) *ListNotifyEventInvoker {
	requestDef := GenReqDefForListNotifyEvent()
	return &ListNotifyEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ListNotifySmnTopicConfig(request *model.ListNotifySmnTopicConfigRequest) (*model.ListNotifySmnTopicConfigResponse, error) {
	requestDef := GenReqDefForListNotifySmnTopicConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNotifySmnTopicConfigResponse), nil
	}
}

func (c *MpcClient) ListNotifySmnTopicConfigInvoker(request *model.ListNotifySmnTopicConfigRequest) *ListNotifySmnTopicConfigInvoker {
	requestDef := GenReqDefForListNotifySmnTopicConfig()
	return &ListNotifySmnTopicConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) NotifySmnTopicConfig(request *model.NotifySmnTopicConfigRequest) (*model.NotifySmnTopicConfigResponse, error) {
	requestDef := GenReqDefForNotifySmnTopicConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NotifySmnTopicConfigResponse), nil
	}
}

func (c *MpcClient) NotifySmnTopicConfigInvoker(request *model.NotifySmnTopicConfigRequest) *NotifySmnTopicConfigInvoker {
	requestDef := GenReqDefForNotifySmnTopicConfig()
	return &NotifySmnTopicConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ShowAgenciesTask(request *model.ShowAgenciesTaskRequest) (*model.ShowAgenciesTaskResponse, error) {
	requestDef := GenReqDefForShowAgenciesTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAgenciesTaskResponse), nil
	}
}

func (c *MpcClient) ShowAgenciesTaskInvoker(request *model.ShowAgenciesTaskRequest) *ShowAgenciesTaskInvoker {
	requestDef := GenReqDefForShowAgenciesTask()
	return &ShowAgenciesTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) UpdateBucketAuthorized(request *model.UpdateBucketAuthorizedRequest) (*model.UpdateBucketAuthorizedResponse, error) {
	requestDef := GenReqDefForUpdateBucketAuthorized()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBucketAuthorizedResponse), nil
	}
}

func (c *MpcClient) UpdateBucketAuthorizedInvoker(request *model.UpdateBucketAuthorizedRequest) *UpdateBucketAuthorizedInvoker {
	requestDef := GenReqDefForUpdateBucketAuthorized()
	return &UpdateBucketAuthorizedInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) CreateEditingJob(request *model.CreateEditingJobRequest) (*model.CreateEditingJobResponse, error) {
	requestDef := GenReqDefForCreateEditingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEditingJobResponse), nil
	}
}

func (c *MpcClient) CreateEditingJobInvoker(request *model.CreateEditingJobRequest) *CreateEditingJobInvoker {
	requestDef := GenReqDefForCreateEditingJob()
	return &CreateEditingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) DeleteEditingJob(request *model.DeleteEditingJobRequest) (*model.DeleteEditingJobResponse, error) {
	requestDef := GenReqDefForDeleteEditingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEditingJobResponse), nil
	}
}

func (c *MpcClient) DeleteEditingJobInvoker(request *model.DeleteEditingJobRequest) *DeleteEditingJobInvoker {
	requestDef := GenReqDefForDeleteEditingJob()
	return &DeleteEditingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ListEditingJob(request *model.ListEditingJobRequest) (*model.ListEditingJobResponse, error) {
	requestDef := GenReqDefForListEditingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEditingJobResponse), nil
	}
}

func (c *MpcClient) ListEditingJobInvoker(request *model.ListEditingJobRequest) *ListEditingJobInvoker {
	requestDef := GenReqDefForListEditingJob()
	return &ListEditingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) CreateEncryptTask(request *model.CreateEncryptTaskRequest) (*model.CreateEncryptTaskResponse, error) {
	requestDef := GenReqDefForCreateEncryptTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEncryptTaskResponse), nil
	}
}

func (c *MpcClient) CreateEncryptTaskInvoker(request *model.CreateEncryptTaskRequest) *CreateEncryptTaskInvoker {
	requestDef := GenReqDefForCreateEncryptTask()
	return &CreateEncryptTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) DeleteEncryptTask(request *model.DeleteEncryptTaskRequest) (*model.DeleteEncryptTaskResponse, error) {
	requestDef := GenReqDefForDeleteEncryptTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEncryptTaskResponse), nil
	}
}

func (c *MpcClient) DeleteEncryptTaskInvoker(request *model.DeleteEncryptTaskRequest) *DeleteEncryptTaskInvoker {
	requestDef := GenReqDefForDeleteEncryptTask()
	return &DeleteEncryptTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ListEncryptTask(request *model.ListEncryptTaskRequest) (*model.ListEncryptTaskResponse, error) {
	requestDef := GenReqDefForListEncryptTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEncryptTaskResponse), nil
	}
}

func (c *MpcClient) ListEncryptTaskInvoker(request *model.ListEncryptTaskRequest) *ListEncryptTaskInvoker {
	requestDef := GenReqDefForListEncryptTask()
	return &ListEncryptTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) CreateExtractTask(request *model.CreateExtractTaskRequest) (*model.CreateExtractTaskResponse, error) {
	requestDef := GenReqDefForCreateExtractTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateExtractTaskResponse), nil
	}
}

func (c *MpcClient) CreateExtractTaskInvoker(request *model.CreateExtractTaskRequest) *CreateExtractTaskInvoker {
	requestDef := GenReqDefForCreateExtractTask()
	return &CreateExtractTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) DeleteExtractTask(request *model.DeleteExtractTaskRequest) (*model.DeleteExtractTaskResponse, error) {
	requestDef := GenReqDefForDeleteExtractTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteExtractTaskResponse), nil
	}
}

func (c *MpcClient) DeleteExtractTaskInvoker(request *model.DeleteExtractTaskRequest) *DeleteExtractTaskInvoker {
	requestDef := GenReqDefForDeleteExtractTask()
	return &DeleteExtractTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ListExtractTask(request *model.ListExtractTaskRequest) (*model.ListExtractTaskResponse, error) {
	requestDef := GenReqDefForListExtractTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListExtractTaskResponse), nil
	}
}

func (c *MpcClient) ListExtractTaskInvoker(request *model.ListExtractTaskRequest) *ListExtractTaskInvoker {
	requestDef := GenReqDefForListExtractTask()
	return &ListExtractTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) CreateMbTasksReport(request *model.CreateMbTasksReportRequest) (*model.CreateMbTasksReportResponse, error) {
	requestDef := GenReqDefForCreateMbTasksReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMbTasksReportResponse), nil
	}
}

func (c *MpcClient) CreateMbTasksReportInvoker(request *model.CreateMbTasksReportRequest) *CreateMbTasksReportInvoker {
	requestDef := GenReqDefForCreateMbTasksReport()
	return &CreateMbTasksReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) CreateMergeChannelsTask(request *model.CreateMergeChannelsTaskRequest) (*model.CreateMergeChannelsTaskResponse, error) {
	requestDef := GenReqDefForCreateMergeChannelsTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMergeChannelsTaskResponse), nil
	}
}

func (c *MpcClient) CreateMergeChannelsTaskInvoker(request *model.CreateMergeChannelsTaskRequest) *CreateMergeChannelsTaskInvoker {
	requestDef := GenReqDefForCreateMergeChannelsTask()
	return &CreateMergeChannelsTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) CreateResetTracksTask(request *model.CreateResetTracksTaskRequest) (*model.CreateResetTracksTaskResponse, error) {
	requestDef := GenReqDefForCreateResetTracksTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResetTracksTaskResponse), nil
	}
}

func (c *MpcClient) CreateResetTracksTaskInvoker(request *model.CreateResetTracksTaskRequest) *CreateResetTracksTaskInvoker {
	requestDef := GenReqDefForCreateResetTracksTask()
	return &CreateResetTracksTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) DeleteMergeChannelsTask(request *model.DeleteMergeChannelsTaskRequest) (*model.DeleteMergeChannelsTaskResponse, error) {
	requestDef := GenReqDefForDeleteMergeChannelsTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMergeChannelsTaskResponse), nil
	}
}

func (c *MpcClient) DeleteMergeChannelsTaskInvoker(request *model.DeleteMergeChannelsTaskRequest) *DeleteMergeChannelsTaskInvoker {
	requestDef := GenReqDefForDeleteMergeChannelsTask()
	return &DeleteMergeChannelsTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) DeleteResetTracksTask(request *model.DeleteResetTracksTaskRequest) (*model.DeleteResetTracksTaskResponse, error) {
	requestDef := GenReqDefForDeleteResetTracksTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResetTracksTaskResponse), nil
	}
}

func (c *MpcClient) DeleteResetTracksTaskInvoker(request *model.DeleteResetTracksTaskRequest) *DeleteResetTracksTaskInvoker {
	requestDef := GenReqDefForDeleteResetTracksTask()
	return &DeleteResetTracksTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ListMergeChannelsTask(request *model.ListMergeChannelsTaskRequest) (*model.ListMergeChannelsTaskResponse, error) {
	requestDef := GenReqDefForListMergeChannelsTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMergeChannelsTaskResponse), nil
	}
}

func (c *MpcClient) ListMergeChannelsTaskInvoker(request *model.ListMergeChannelsTaskRequest) *ListMergeChannelsTaskInvoker {
	requestDef := GenReqDefForListMergeChannelsTask()
	return &ListMergeChannelsTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ListResetTracksTask(request *model.ListResetTracksTaskRequest) (*model.ListResetTracksTaskResponse, error) {
	requestDef := GenReqDefForListResetTracksTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResetTracksTaskResponse), nil
	}
}

func (c *MpcClient) ListResetTracksTaskInvoker(request *model.ListResetTracksTaskRequest) *ListResetTracksTaskInvoker {
	requestDef := GenReqDefForListResetTracksTask()
	return &ListResetTracksTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) CreateMediaProcessTask(request *model.CreateMediaProcessTaskRequest) (*model.CreateMediaProcessTaskResponse, error) {
	requestDef := GenReqDefForCreateMediaProcessTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMediaProcessTaskResponse), nil
	}
}

func (c *MpcClient) CreateMediaProcessTaskInvoker(request *model.CreateMediaProcessTaskRequest) *CreateMediaProcessTaskInvoker {
	requestDef := GenReqDefForCreateMediaProcessTask()
	return &CreateMediaProcessTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) DeleteMediaProcessTask(request *model.DeleteMediaProcessTaskRequest) (*model.DeleteMediaProcessTaskResponse, error) {
	requestDef := GenReqDefForDeleteMediaProcessTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMediaProcessTaskResponse), nil
	}
}

func (c *MpcClient) DeleteMediaProcessTaskInvoker(request *model.DeleteMediaProcessTaskRequest) *DeleteMediaProcessTaskInvoker {
	requestDef := GenReqDefForDeleteMediaProcessTask()
	return &DeleteMediaProcessTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ListMediaProcessTask(request *model.ListMediaProcessTaskRequest) (*model.ListMediaProcessTaskResponse, error) {
	requestDef := GenReqDefForListMediaProcessTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMediaProcessTaskResponse), nil
	}
}

func (c *MpcClient) ListMediaProcessTaskInvoker(request *model.ListMediaProcessTaskRequest) *ListMediaProcessTaskInvoker {
	requestDef := GenReqDefForListMediaProcessTask()
	return &ListMediaProcessTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) CreateMpeCallBack(request *model.CreateMpeCallBackRequest) (*model.CreateMpeCallBackResponse, error) {
	requestDef := GenReqDefForCreateMpeCallBack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMpeCallBackResponse), nil
	}
}

func (c *MpcClient) CreateMpeCallBackInvoker(request *model.CreateMpeCallBackRequest) *CreateMpeCallBackInvoker {
	requestDef := GenReqDefForCreateMpeCallBack()
	return &CreateMpeCallBackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) CreateQualityEnhanceTemplate(request *model.CreateQualityEnhanceTemplateRequest) (*model.CreateQualityEnhanceTemplateResponse, error) {
	requestDef := GenReqDefForCreateQualityEnhanceTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateQualityEnhanceTemplateResponse), nil
	}
}

func (c *MpcClient) CreateQualityEnhanceTemplateInvoker(request *model.CreateQualityEnhanceTemplateRequest) *CreateQualityEnhanceTemplateInvoker {
	requestDef := GenReqDefForCreateQualityEnhanceTemplate()
	return &CreateQualityEnhanceTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) DeleteQualityEnhanceTemplate(request *model.DeleteQualityEnhanceTemplateRequest) (*model.DeleteQualityEnhanceTemplateResponse, error) {
	requestDef := GenReqDefForDeleteQualityEnhanceTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteQualityEnhanceTemplateResponse), nil
	}
}

func (c *MpcClient) DeleteQualityEnhanceTemplateInvoker(request *model.DeleteQualityEnhanceTemplateRequest) *DeleteQualityEnhanceTemplateInvoker {
	requestDef := GenReqDefForDeleteQualityEnhanceTemplate()
	return &DeleteQualityEnhanceTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ListQualityEnhanceDefaultTemplate(request *model.ListQualityEnhanceDefaultTemplateRequest) (*model.ListQualityEnhanceDefaultTemplateResponse, error) {
	requestDef := GenReqDefForListQualityEnhanceDefaultTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQualityEnhanceDefaultTemplateResponse), nil
	}
}

func (c *MpcClient) ListQualityEnhanceDefaultTemplateInvoker(request *model.ListQualityEnhanceDefaultTemplateRequest) *ListQualityEnhanceDefaultTemplateInvoker {
	requestDef := GenReqDefForListQualityEnhanceDefaultTemplate()
	return &ListQualityEnhanceDefaultTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) UpdateQualityEnhanceTemplate(request *model.UpdateQualityEnhanceTemplateRequest) (*model.UpdateQualityEnhanceTemplateResponse, error) {
	requestDef := GenReqDefForUpdateQualityEnhanceTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateQualityEnhanceTemplateResponse), nil
	}
}

func (c *MpcClient) UpdateQualityEnhanceTemplateInvoker(request *model.UpdateQualityEnhanceTemplateRequest) *UpdateQualityEnhanceTemplateInvoker {
	requestDef := GenReqDefForUpdateQualityEnhanceTemplate()
	return &UpdateQualityEnhanceTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ListTranscodeDetail(request *model.ListTranscodeDetailRequest) (*model.ListTranscodeDetailResponse, error) {
	requestDef := GenReqDefForListTranscodeDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTranscodeDetailResponse), nil
	}
}

func (c *MpcClient) ListTranscodeDetailInvoker(request *model.ListTranscodeDetailRequest) *ListTranscodeDetailInvoker {
	requestDef := GenReqDefForListTranscodeDetail()
	return &ListTranscodeDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) CancelRemuxTask(request *model.CancelRemuxTaskRequest) (*model.CancelRemuxTaskResponse, error) {
	requestDef := GenReqDefForCancelRemuxTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelRemuxTaskResponse), nil
	}
}

func (c *MpcClient) CancelRemuxTaskInvoker(request *model.CancelRemuxTaskRequest) *CancelRemuxTaskInvoker {
	requestDef := GenReqDefForCancelRemuxTask()
	return &CancelRemuxTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) CreateRemuxTask(request *model.CreateRemuxTaskRequest) (*model.CreateRemuxTaskResponse, error) {
	requestDef := GenReqDefForCreateRemuxTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRemuxTaskResponse), nil
	}
}

func (c *MpcClient) CreateRemuxTaskInvoker(request *model.CreateRemuxTaskRequest) *CreateRemuxTaskInvoker {
	requestDef := GenReqDefForCreateRemuxTask()
	return &CreateRemuxTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) CreateRetryRemuxTask(request *model.CreateRetryRemuxTaskRequest) (*model.CreateRetryRemuxTaskResponse, error) {
	requestDef := GenReqDefForCreateRetryRemuxTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRetryRemuxTaskResponse), nil
	}
}

func (c *MpcClient) CreateRetryRemuxTaskInvoker(request *model.CreateRetryRemuxTaskRequest) *CreateRetryRemuxTaskInvoker {
	requestDef := GenReqDefForCreateRetryRemuxTask()
	return &CreateRetryRemuxTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) DeleteRemuxTask(request *model.DeleteRemuxTaskRequest) (*model.DeleteRemuxTaskResponse, error) {
	requestDef := GenReqDefForDeleteRemuxTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRemuxTaskResponse), nil
	}
}

func (c *MpcClient) DeleteRemuxTaskInvoker(request *model.DeleteRemuxTaskRequest) *DeleteRemuxTaskInvoker {
	requestDef := GenReqDefForDeleteRemuxTask()
	return &DeleteRemuxTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ListRemuxTask(request *model.ListRemuxTaskRequest) (*model.ListRemuxTaskResponse, error) {
	requestDef := GenReqDefForListRemuxTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRemuxTaskResponse), nil
	}
}

func (c *MpcClient) ListRemuxTaskInvoker(request *model.ListRemuxTaskRequest) *ListRemuxTaskInvoker {
	requestDef := GenReqDefForListRemuxTask()
	return &ListRemuxTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) CreateTemplateGroup(request *model.CreateTemplateGroupRequest) (*model.CreateTemplateGroupResponse, error) {
	requestDef := GenReqDefForCreateTemplateGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTemplateGroupResponse), nil
	}
}

func (c *MpcClient) CreateTemplateGroupInvoker(request *model.CreateTemplateGroupRequest) *CreateTemplateGroupInvoker {
	requestDef := GenReqDefForCreateTemplateGroup()
	return &CreateTemplateGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) DeleteTemplateGroup(request *model.DeleteTemplateGroupRequest) (*model.DeleteTemplateGroupResponse, error) {
	requestDef := GenReqDefForDeleteTemplateGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTemplateGroupResponse), nil
	}
}

func (c *MpcClient) DeleteTemplateGroupInvoker(request *model.DeleteTemplateGroupRequest) *DeleteTemplateGroupInvoker {
	requestDef := GenReqDefForDeleteTemplateGroup()
	return &DeleteTemplateGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ListTemplateGroup(request *model.ListTemplateGroupRequest) (*model.ListTemplateGroupResponse, error) {
	requestDef := GenReqDefForListTemplateGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplateGroupResponse), nil
	}
}

func (c *MpcClient) ListTemplateGroupInvoker(request *model.ListTemplateGroupRequest) *ListTemplateGroupInvoker {
	requestDef := GenReqDefForListTemplateGroup()
	return &ListTemplateGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) UpdateTemplateGroup(request *model.UpdateTemplateGroupRequest) (*model.UpdateTemplateGroupResponse, error) {
	requestDef := GenReqDefForUpdateTemplateGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTemplateGroupResponse), nil
	}
}

func (c *MpcClient) UpdateTemplateGroupInvoker(request *model.UpdateTemplateGroupRequest) *UpdateTemplateGroupInvoker {
	requestDef := GenReqDefForUpdateTemplateGroup()
	return &UpdateTemplateGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) CreateThumbnailsTask(request *model.CreateThumbnailsTaskRequest) (*model.CreateThumbnailsTaskResponse, error) {
	requestDef := GenReqDefForCreateThumbnailsTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateThumbnailsTaskResponse), nil
	}
}

func (c *MpcClient) CreateThumbnailsTaskInvoker(request *model.CreateThumbnailsTaskRequest) *CreateThumbnailsTaskInvoker {
	requestDef := GenReqDefForCreateThumbnailsTask()
	return &CreateThumbnailsTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) DeleteThumbnailsTask(request *model.DeleteThumbnailsTaskRequest) (*model.DeleteThumbnailsTaskResponse, error) {
	requestDef := GenReqDefForDeleteThumbnailsTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteThumbnailsTaskResponse), nil
	}
}

func (c *MpcClient) DeleteThumbnailsTaskInvoker(request *model.DeleteThumbnailsTaskRequest) *DeleteThumbnailsTaskInvoker {
	requestDef := GenReqDefForDeleteThumbnailsTask()
	return &DeleteThumbnailsTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ListThumbnailsTask(request *model.ListThumbnailsTaskRequest) (*model.ListThumbnailsTaskResponse, error) {
	requestDef := GenReqDefForListThumbnailsTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListThumbnailsTaskResponse), nil
	}
}

func (c *MpcClient) ListThumbnailsTaskInvoker(request *model.ListThumbnailsTaskRequest) *ListThumbnailsTaskInvoker {
	requestDef := GenReqDefForListThumbnailsTask()
	return &ListThumbnailsTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) CreateTranscodingTask(request *model.CreateTranscodingTaskRequest) (*model.CreateTranscodingTaskResponse, error) {
	requestDef := GenReqDefForCreateTranscodingTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTranscodingTaskResponse), nil
	}
}

func (c *MpcClient) CreateTranscodingTaskInvoker(request *model.CreateTranscodingTaskRequest) *CreateTranscodingTaskInvoker {
	requestDef := GenReqDefForCreateTranscodingTask()
	return &CreateTranscodingTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) DeleteTranscodingTask(request *model.DeleteTranscodingTaskRequest) (*model.DeleteTranscodingTaskResponse, error) {
	requestDef := GenReqDefForDeleteTranscodingTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTranscodingTaskResponse), nil
	}
}

func (c *MpcClient) DeleteTranscodingTaskInvoker(request *model.DeleteTranscodingTaskRequest) *DeleteTranscodingTaskInvoker {
	requestDef := GenReqDefForDeleteTranscodingTask()
	return &DeleteTranscodingTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) DeleteTranscodingTaskByConsole(request *model.DeleteTranscodingTaskByConsoleRequest) (*model.DeleteTranscodingTaskByConsoleResponse, error) {
	requestDef := GenReqDefForDeleteTranscodingTaskByConsole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTranscodingTaskByConsoleResponse), nil
	}
}

func (c *MpcClient) DeleteTranscodingTaskByConsoleInvoker(request *model.DeleteTranscodingTaskByConsoleRequest) *DeleteTranscodingTaskByConsoleInvoker {
	requestDef := GenReqDefForDeleteTranscodingTaskByConsole()
	return &DeleteTranscodingTaskByConsoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ListStatSummary(request *model.ListStatSummaryRequest) (*model.ListStatSummaryResponse, error) {
	requestDef := GenReqDefForListStatSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStatSummaryResponse), nil
	}
}

func (c *MpcClient) ListStatSummaryInvoker(request *model.ListStatSummaryRequest) *ListStatSummaryInvoker {
	requestDef := GenReqDefForListStatSummary()
	return &ListStatSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ListTranscodingTask(request *model.ListTranscodingTaskRequest) (*model.ListTranscodingTaskResponse, error) {
	requestDef := GenReqDefForListTranscodingTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTranscodingTaskResponse), nil
	}
}

func (c *MpcClient) ListTranscodingTaskInvoker(request *model.ListTranscodingTaskRequest) *ListTranscodingTaskInvoker {
	requestDef := GenReqDefForListTranscodingTask()
	return &ListTranscodingTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) CreateTransTemplate(request *model.CreateTransTemplateRequest) (*model.CreateTransTemplateResponse, error) {
	requestDef := GenReqDefForCreateTransTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTransTemplateResponse), nil
	}
}

func (c *MpcClient) CreateTransTemplateInvoker(request *model.CreateTransTemplateRequest) *CreateTransTemplateInvoker {
	requestDef := GenReqDefForCreateTransTemplate()
	return &CreateTransTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) DeleteTemplate(request *model.DeleteTemplateRequest) (*model.DeleteTemplateResponse, error) {
	requestDef := GenReqDefForDeleteTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTemplateResponse), nil
	}
}

func (c *MpcClient) DeleteTemplateInvoker(request *model.DeleteTemplateRequest) *DeleteTemplateInvoker {
	requestDef := GenReqDefForDeleteTemplate()
	return &DeleteTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ListTemplate(request *model.ListTemplateRequest) (*model.ListTemplateResponse, error) {
	requestDef := GenReqDefForListTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplateResponse), nil
	}
}

func (c *MpcClient) ListTemplateInvoker(request *model.ListTemplateRequest) *ListTemplateInvoker {
	requestDef := GenReqDefForListTemplate()
	return &ListTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) UpdateTransTemplate(request *model.UpdateTransTemplateRequest) (*model.UpdateTransTemplateResponse, error) {
	requestDef := GenReqDefForUpdateTransTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTransTemplateResponse), nil
	}
}

func (c *MpcClient) UpdateTransTemplateInvoker(request *model.UpdateTransTemplateRequest) *UpdateTransTemplateInvoker {
	requestDef := GenReqDefForUpdateTransTemplate()
	return &UpdateTransTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) CreateWatermarkTemplate(request *model.CreateWatermarkTemplateRequest) (*model.CreateWatermarkTemplateResponse, error) {
	requestDef := GenReqDefForCreateWatermarkTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWatermarkTemplateResponse), nil
	}
}

func (c *MpcClient) CreateWatermarkTemplateInvoker(request *model.CreateWatermarkTemplateRequest) *CreateWatermarkTemplateInvoker {
	requestDef := GenReqDefForCreateWatermarkTemplate()
	return &CreateWatermarkTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) DeleteWatermarkTemplate(request *model.DeleteWatermarkTemplateRequest) (*model.DeleteWatermarkTemplateResponse, error) {
	requestDef := GenReqDefForDeleteWatermarkTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWatermarkTemplateResponse), nil
	}
}

func (c *MpcClient) DeleteWatermarkTemplateInvoker(request *model.DeleteWatermarkTemplateRequest) *DeleteWatermarkTemplateInvoker {
	requestDef := GenReqDefForDeleteWatermarkTemplate()
	return &DeleteWatermarkTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) ListWatermarkTemplate(request *model.ListWatermarkTemplateRequest) (*model.ListWatermarkTemplateResponse, error) {
	requestDef := GenReqDefForListWatermarkTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWatermarkTemplateResponse), nil
	}
}

func (c *MpcClient) ListWatermarkTemplateInvoker(request *model.ListWatermarkTemplateRequest) *ListWatermarkTemplateInvoker {
	requestDef := GenReqDefForListWatermarkTemplate()
	return &ListWatermarkTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *MpcClient) UpdateWatermarkTemplate(request *model.UpdateWatermarkTemplateRequest) (*model.UpdateWatermarkTemplateResponse, error) {
	requestDef := GenReqDefForUpdateWatermarkTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWatermarkTemplateResponse), nil
	}
}

func (c *MpcClient) UpdateWatermarkTemplateInvoker(request *model.UpdateWatermarkTemplateRequest) *UpdateWatermarkTemplateInvoker {
	requestDef := GenReqDefForUpdateWatermarkTemplate()
	return &UpdateWatermarkTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
