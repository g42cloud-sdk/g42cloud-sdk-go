package v3

import (
	http_client "github.com/g42cloud-sdk/g42cloud-sdk-go/core"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/invoker"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/services/sms/v3/model"
)

type SmsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewSmsClient(hcClient *http_client.HcHttpClient) *SmsClient {
	return &SmsClient{HcClient: hcClient}
}

func SmsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *SmsClient) CheckNetAcl(request *model.CheckNetAclRequest) (*model.CheckNetAclResponse, error) {
	requestDef := GenReqDefForCheckNetAcl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckNetAclResponse), nil
	}
}

func (c *SmsClient) CheckNetAclInvoker(request *model.CheckNetAclRequest) *CheckNetAclInvoker {
	requestDef := GenReqDefForCheckNetAcl()
	return &CheckNetAclInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) CollectLog(request *model.CollectLogRequest) (*model.CollectLogResponse, error) {
	requestDef := GenReqDefForCollectLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectLogResponse), nil
	}
}

func (c *SmsClient) CollectLogInvoker(request *model.CollectLogRequest) *CollectLogInvoker {
	requestDef := GenReqDefForCollectLog()
	return &CollectLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) CreateMigproject(request *model.CreateMigprojectRequest) (*model.CreateMigprojectResponse, error) {
	requestDef := GenReqDefForCreateMigproject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMigprojectResponse), nil
	}
}

func (c *SmsClient) CreateMigprojectInvoker(request *model.CreateMigprojectRequest) *CreateMigprojectInvoker {
	requestDef := GenReqDefForCreateMigproject()
	return &CreateMigprojectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) CreateTask(request *model.CreateTaskRequest) (*model.CreateTaskResponse, error) {
	requestDef := GenReqDefForCreateTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTaskResponse), nil
	}
}

func (c *SmsClient) CreateTaskInvoker(request *model.CreateTaskRequest) *CreateTaskInvoker {
	requestDef := GenReqDefForCreateTask()
	return &CreateTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) CreateTemplate(request *model.CreateTemplateRequest) (*model.CreateTemplateResponse, error) {
	requestDef := GenReqDefForCreateTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTemplateResponse), nil
	}
}

func (c *SmsClient) CreateTemplateInvoker(request *model.CreateTemplateRequest) *CreateTemplateInvoker {
	requestDef := GenReqDefForCreateTemplate()
	return &CreateTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) DeleteMigproject(request *model.DeleteMigprojectRequest) (*model.DeleteMigprojectResponse, error) {
	requestDef := GenReqDefForDeleteMigproject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMigprojectResponse), nil
	}
}

func (c *SmsClient) DeleteMigprojectInvoker(request *model.DeleteMigprojectRequest) *DeleteMigprojectInvoker {
	requestDef := GenReqDefForDeleteMigproject()
	return &DeleteMigprojectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) DeleteServer(request *model.DeleteServerRequest) (*model.DeleteServerResponse, error) {
	requestDef := GenReqDefForDeleteServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServerResponse), nil
	}
}

func (c *SmsClient) DeleteServerInvoker(request *model.DeleteServerRequest) *DeleteServerInvoker {
	requestDef := GenReqDefForDeleteServer()
	return &DeleteServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) DeleteServers(request *model.DeleteServersRequest) (*model.DeleteServersResponse, error) {
	requestDef := GenReqDefForDeleteServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServersResponse), nil
	}
}

func (c *SmsClient) DeleteServersInvoker(request *model.DeleteServersRequest) *DeleteServersInvoker {
	requestDef := GenReqDefForDeleteServers()
	return &DeleteServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) DeleteTask(request *model.DeleteTaskRequest) (*model.DeleteTaskResponse, error) {
	requestDef := GenReqDefForDeleteTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTaskResponse), nil
	}
}

func (c *SmsClient) DeleteTaskInvoker(request *model.DeleteTaskRequest) *DeleteTaskInvoker {
	requestDef := GenReqDefForDeleteTask()
	return &DeleteTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) DeleteTasks(request *model.DeleteTasksRequest) (*model.DeleteTasksResponse, error) {
	requestDef := GenReqDefForDeleteTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTasksResponse), nil
	}
}

func (c *SmsClient) DeleteTasksInvoker(request *model.DeleteTasksRequest) *DeleteTasksInvoker {
	requestDef := GenReqDefForDeleteTasks()
	return &DeleteTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) DeleteTemplate(request *model.DeleteTemplateRequest) (*model.DeleteTemplateResponse, error) {
	requestDef := GenReqDefForDeleteTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTemplateResponse), nil
	}
}

func (c *SmsClient) DeleteTemplateInvoker(request *model.DeleteTemplateRequest) *DeleteTemplateInvoker {
	requestDef := GenReqDefForDeleteTemplate()
	return &DeleteTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) DeleteTemplates(request *model.DeleteTemplatesRequest) (*model.DeleteTemplatesResponse, error) {
	requestDef := GenReqDefForDeleteTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTemplatesResponse), nil
	}
}

func (c *SmsClient) DeleteTemplatesInvoker(request *model.DeleteTemplatesRequest) *DeleteTemplatesInvoker {
	requestDef := GenReqDefForDeleteTemplates()
	return &DeleteTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) ListErrorServers(request *model.ListErrorServersRequest) (*model.ListErrorServersResponse, error) {
	requestDef := GenReqDefForListErrorServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListErrorServersResponse), nil
	}
}

func (c *SmsClient) ListErrorServersInvoker(request *model.ListErrorServersRequest) *ListErrorServersInvoker {
	requestDef := GenReqDefForListErrorServers()
	return &ListErrorServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) ListMigprojects(request *model.ListMigprojectsRequest) (*model.ListMigprojectsResponse, error) {
	requestDef := GenReqDefForListMigprojects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMigprojectsResponse), nil
	}
}

func (c *SmsClient) ListMigprojectsInvoker(request *model.ListMigprojectsRequest) *ListMigprojectsInvoker {
	requestDef := GenReqDefForListMigprojects()
	return &ListMigprojectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) ListServers(request *model.ListServersRequest) (*model.ListServersResponse, error) {
	requestDef := GenReqDefForListServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServersResponse), nil
	}
}

func (c *SmsClient) ListServersInvoker(request *model.ListServersRequest) *ListServersInvoker {
	requestDef := GenReqDefForListServers()
	return &ListServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) ListTasks(request *model.ListTasksRequest) (*model.ListTasksResponse, error) {
	requestDef := GenReqDefForListTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTasksResponse), nil
	}
}

func (c *SmsClient) ListTasksInvoker(request *model.ListTasksRequest) *ListTasksInvoker {
	requestDef := GenReqDefForListTasks()
	return &ListTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) ListTemplates(request *model.ListTemplatesRequest) (*model.ListTemplatesResponse, error) {
	requestDef := GenReqDefForListTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplatesResponse), nil
	}
}

func (c *SmsClient) ListTemplatesInvoker(request *model.ListTemplatesRequest) *ListTemplatesInvoker {
	requestDef := GenReqDefForListTemplates()
	return &ListTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) RegisterServer(request *model.RegisterServerRequest) (*model.RegisterServerResponse, error) {
	requestDef := GenReqDefForRegisterServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterServerResponse), nil
	}
}

func (c *SmsClient) RegisterServerInvoker(request *model.RegisterServerRequest) *RegisterServerInvoker {
	requestDef := GenReqDefForRegisterServer()
	return &RegisterServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) ShowCertKey(request *model.ShowCertKeyRequest) (*model.ShowCertKeyResponse, error) {
	requestDef := GenReqDefForShowCertKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertKeyResponse), nil
	}
}

func (c *SmsClient) ShowCertKeyInvoker(request *model.ShowCertKeyRequest) *ShowCertKeyInvoker {
	requestDef := GenReqDefForShowCertKey()
	return &ShowCertKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) ShowCommand(request *model.ShowCommandRequest) (*model.ShowCommandResponse, error) {
	requestDef := GenReqDefForShowCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCommandResponse), nil
	}
}

func (c *SmsClient) ShowCommandInvoker(request *model.ShowCommandRequest) *ShowCommandInvoker {
	requestDef := GenReqDefForShowCommand()
	return &ShowCommandInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) ShowConfigSetting(request *model.ShowConfigSettingRequest) (*model.ShowConfigSettingResponse, error) {
	requestDef := GenReqDefForShowConfigSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigSettingResponse), nil
	}
}

func (c *SmsClient) ShowConfigSettingInvoker(request *model.ShowConfigSettingRequest) *ShowConfigSettingInvoker {
	requestDef := GenReqDefForShowConfigSetting()
	return &ShowConfigSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) ShowMigproject(request *model.ShowMigprojectRequest) (*model.ShowMigprojectResponse, error) {
	requestDef := GenReqDefForShowMigproject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMigprojectResponse), nil
	}
}

func (c *SmsClient) ShowMigprojectInvoker(request *model.ShowMigprojectRequest) *ShowMigprojectInvoker {
	requestDef := GenReqDefForShowMigproject()
	return &ShowMigprojectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) ShowOverview(request *model.ShowOverviewRequest) (*model.ShowOverviewResponse, error) {
	requestDef := GenReqDefForShowOverview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOverviewResponse), nil
	}
}

func (c *SmsClient) ShowOverviewInvoker(request *model.ShowOverviewRequest) *ShowOverviewInvoker {
	requestDef := GenReqDefForShowOverview()
	return &ShowOverviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) ShowPassphrase(request *model.ShowPassphraseRequest) (*model.ShowPassphraseResponse, error) {
	requestDef := GenReqDefForShowPassphrase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPassphraseResponse), nil
	}
}

func (c *SmsClient) ShowPassphraseInvoker(request *model.ShowPassphraseRequest) *ShowPassphraseInvoker {
	requestDef := GenReqDefForShowPassphrase()
	return &ShowPassphraseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) ShowServer(request *model.ShowServerRequest) (*model.ShowServerResponse, error) {
	requestDef := GenReqDefForShowServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerResponse), nil
	}
}

func (c *SmsClient) ShowServerInvoker(request *model.ShowServerRequest) *ShowServerInvoker {
	requestDef := GenReqDefForShowServer()
	return &ShowServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) ShowSha256(request *model.ShowSha256Request) (*model.ShowSha256Response, error) {
	requestDef := GenReqDefForShowSha256()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSha256Response), nil
	}
}

func (c *SmsClient) ShowSha256Invoker(request *model.ShowSha256Request) *ShowSha256Invoker {
	requestDef := GenReqDefForShowSha256()
	return &ShowSha256Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) ShowTargetPassword(request *model.ShowTargetPasswordRequest) (*model.ShowTargetPasswordResponse, error) {
	requestDef := GenReqDefForShowTargetPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTargetPasswordResponse), nil
	}
}

func (c *SmsClient) ShowTargetPasswordInvoker(request *model.ShowTargetPasswordRequest) *ShowTargetPasswordInvoker {
	requestDef := GenReqDefForShowTargetPassword()
	return &ShowTargetPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) ShowTask(request *model.ShowTaskRequest) (*model.ShowTaskResponse, error) {
	requestDef := GenReqDefForShowTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskResponse), nil
	}
}

func (c *SmsClient) ShowTaskInvoker(request *model.ShowTaskRequest) *ShowTaskInvoker {
	requestDef := GenReqDefForShowTask()
	return &ShowTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) ShowTemplate(request *model.ShowTemplateRequest) (*model.ShowTemplateResponse, error) {
	requestDef := GenReqDefForShowTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateResponse), nil
	}
}

func (c *SmsClient) ShowTemplateInvoker(request *model.ShowTemplateRequest) *ShowTemplateInvoker {
	requestDef := GenReqDefForShowTemplate()
	return &ShowTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) ShowsSpeedLimits(request *model.ShowsSpeedLimitsRequest) (*model.ShowsSpeedLimitsResponse, error) {
	requestDef := GenReqDefForShowsSpeedLimits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowsSpeedLimitsResponse), nil
	}
}

func (c *SmsClient) ShowsSpeedLimitsInvoker(request *model.ShowsSpeedLimitsRequest) *ShowsSpeedLimitsInvoker {
	requestDef := GenReqDefForShowsSpeedLimits()
	return &ShowsSpeedLimitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) UnlockTargetEcs(request *model.UnlockTargetEcsRequest) (*model.UnlockTargetEcsResponse, error) {
	requestDef := GenReqDefForUnlockTargetEcs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnlockTargetEcsResponse), nil
	}
}

func (c *SmsClient) UnlockTargetEcsInvoker(request *model.UnlockTargetEcsRequest) *UnlockTargetEcsInvoker {
	requestDef := GenReqDefForUnlockTargetEcs()
	return &UnlockTargetEcsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) UpdateCommandResult(request *model.UpdateCommandResultRequest) (*model.UpdateCommandResultResponse, error) {
	requestDef := GenReqDefForUpdateCommandResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCommandResultResponse), nil
	}
}

func (c *SmsClient) UpdateCommandResultInvoker(request *model.UpdateCommandResultRequest) *UpdateCommandResultInvoker {
	requestDef := GenReqDefForUpdateCommandResult()
	return &UpdateCommandResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) UpdateCopyState(request *model.UpdateCopyStateRequest) (*model.UpdateCopyStateResponse, error) {
	requestDef := GenReqDefForUpdateCopyState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCopyStateResponse), nil
	}
}

func (c *SmsClient) UpdateCopyStateInvoker(request *model.UpdateCopyStateRequest) *UpdateCopyStateInvoker {
	requestDef := GenReqDefForUpdateCopyState()
	return &UpdateCopyStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) UpdateDefaultMigproject(request *model.UpdateDefaultMigprojectRequest) (*model.UpdateDefaultMigprojectResponse, error) {
	requestDef := GenReqDefForUpdateDefaultMigproject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDefaultMigprojectResponse), nil
	}
}

func (c *SmsClient) UpdateDefaultMigprojectInvoker(request *model.UpdateDefaultMigprojectRequest) *UpdateDefaultMigprojectInvoker {
	requestDef := GenReqDefForUpdateDefaultMigproject()
	return &UpdateDefaultMigprojectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) UpdateDiskInfo(request *model.UpdateDiskInfoRequest) (*model.UpdateDiskInfoResponse, error) {
	requestDef := GenReqDefForUpdateDiskInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDiskInfoResponse), nil
	}
}

func (c *SmsClient) UpdateDiskInfoInvoker(request *model.UpdateDiskInfoRequest) *UpdateDiskInfoInvoker {
	requestDef := GenReqDefForUpdateDiskInfo()
	return &UpdateDiskInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) UpdateMigproject(request *model.UpdateMigprojectRequest) (*model.UpdateMigprojectResponse, error) {
	requestDef := GenReqDefForUpdateMigproject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMigprojectResponse), nil
	}
}

func (c *SmsClient) UpdateMigprojectInvoker(request *model.UpdateMigprojectRequest) *UpdateMigprojectInvoker {
	requestDef := GenReqDefForUpdateMigproject()
	return &UpdateMigprojectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) UpdateNetworkCheckInfo(request *model.UpdateNetworkCheckInfoRequest) (*model.UpdateNetworkCheckInfoResponse, error) {
	requestDef := GenReqDefForUpdateNetworkCheckInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNetworkCheckInfoResponse), nil
	}
}

func (c *SmsClient) UpdateNetworkCheckInfoInvoker(request *model.UpdateNetworkCheckInfoRequest) *UpdateNetworkCheckInfoInvoker {
	requestDef := GenReqDefForUpdateNetworkCheckInfo()
	return &UpdateNetworkCheckInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) UpdateServerName(request *model.UpdateServerNameRequest) (*model.UpdateServerNameResponse, error) {
	requestDef := GenReqDefForUpdateServerName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServerNameResponse), nil
	}
}

func (c *SmsClient) UpdateServerNameInvoker(request *model.UpdateServerNameRequest) *UpdateServerNameInvoker {
	requestDef := GenReqDefForUpdateServerName()
	return &UpdateServerNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) UpdateSpeed(request *model.UpdateSpeedRequest) (*model.UpdateSpeedResponse, error) {
	requestDef := GenReqDefForUpdateSpeed()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSpeedResponse), nil
	}
}

func (c *SmsClient) UpdateSpeedInvoker(request *model.UpdateSpeedRequest) *UpdateSpeedInvoker {
	requestDef := GenReqDefForUpdateSpeed()
	return &UpdateSpeedInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) UpdateTask(request *model.UpdateTaskRequest) (*model.UpdateTaskResponse, error) {
	requestDef := GenReqDefForUpdateTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTaskResponse), nil
	}
}

func (c *SmsClient) UpdateTaskInvoker(request *model.UpdateTaskRequest) *UpdateTaskInvoker {
	requestDef := GenReqDefForUpdateTask()
	return &UpdateTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) UpdateTaskSpeed(request *model.UpdateTaskSpeedRequest) (*model.UpdateTaskSpeedResponse, error) {
	requestDef := GenReqDefForUpdateTaskSpeed()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTaskSpeedResponse), nil
	}
}

func (c *SmsClient) UpdateTaskSpeedInvoker(request *model.UpdateTaskSpeedRequest) *UpdateTaskSpeedInvoker {
	requestDef := GenReqDefForUpdateTaskSpeed()
	return &UpdateTaskSpeedInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) UpdateTaskStatus(request *model.UpdateTaskStatusRequest) (*model.UpdateTaskStatusResponse, error) {
	requestDef := GenReqDefForUpdateTaskStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTaskStatusResponse), nil
	}
}

func (c *SmsClient) UpdateTaskStatusInvoker(request *model.UpdateTaskStatusRequest) *UpdateTaskStatusInvoker {
	requestDef := GenReqDefForUpdateTaskStatus()
	return &UpdateTaskStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) UpdateTemplate(request *model.UpdateTemplateRequest) (*model.UpdateTemplateResponse, error) {
	requestDef := GenReqDefForUpdateTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTemplateResponse), nil
	}
}

func (c *SmsClient) UpdateTemplateInvoker(request *model.UpdateTemplateRequest) *UpdateTemplateInvoker {
	requestDef := GenReqDefForUpdateTemplate()
	return &UpdateTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) UploadSpecialConfigurationSetting(request *model.UploadSpecialConfigurationSettingRequest) (*model.UploadSpecialConfigurationSettingResponse, error) {
	requestDef := GenReqDefForUploadSpecialConfigurationSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadSpecialConfigurationSettingResponse), nil
	}
}

func (c *SmsClient) UploadSpecialConfigurationSettingInvoker(request *model.UploadSpecialConfigurationSettingRequest) *UploadSpecialConfigurationSettingInvoker {
	requestDef := GenReqDefForUploadSpecialConfigurationSetting()
	return &UploadSpecialConfigurationSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SmsClient) ShowConfig(request *model.ShowConfigRequest) (*model.ShowConfigResponse, error) {
	requestDef := GenReqDefForShowConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigResponse), nil
	}
}

func (c *SmsClient) ShowConfigInvoker(request *model.ShowConfigRequest) *ShowConfigInvoker {
	requestDef := GenReqDefForShowConfig()
	return &ShowConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
