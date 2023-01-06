package v1

import (
	http_client "github.com/g42cloud-sdk/g42cloud-sdk-go/core"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/invoker"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/services/cbr/v1/model"
)

type CbrClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCbrClient(hcClient *http_client.HcHttpClient) *CbrClient {
	return &CbrClient{HcClient: hcClient}
}

func CbrClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *CbrClient) AddMember(request *model.AddMemberRequest) (*model.AddMemberResponse, error) {
	requestDef := GenReqDefForAddMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddMemberResponse), nil
	}
}

func (c *CbrClient) AddMemberInvoker(request *model.AddMemberRequest) *AddMemberInvoker {
	requestDef := GenReqDefForAddMember()
	return &AddMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) AddVaultResource(request *model.AddVaultResourceRequest) (*model.AddVaultResourceResponse, error) {
	requestDef := GenReqDefForAddVaultResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddVaultResourceResponse), nil
	}
}

func (c *CbrClient) AddVaultResourceInvoker(request *model.AddVaultResourceRequest) *AddVaultResourceInvoker {
	requestDef := GenReqDefForAddVaultResource()
	return &AddVaultResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) AssociateVaultPolicy(request *model.AssociateVaultPolicyRequest) (*model.AssociateVaultPolicyResponse, error) {
	requestDef := GenReqDefForAssociateVaultPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateVaultPolicyResponse), nil
	}
}

func (c *CbrClient) AssociateVaultPolicyInvoker(request *model.AssociateVaultPolicyRequest) *AssociateVaultPolicyInvoker {
	requestDef := GenReqDefForAssociateVaultPolicy()
	return &AssociateVaultPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) BatchCreateAndDeleteVaultTags(request *model.BatchCreateAndDeleteVaultTagsRequest) (*model.BatchCreateAndDeleteVaultTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateAndDeleteVaultTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateAndDeleteVaultTagsResponse), nil
	}
}

func (c *CbrClient) BatchCreateAndDeleteVaultTagsInvoker(request *model.BatchCreateAndDeleteVaultTagsRequest) *BatchCreateAndDeleteVaultTagsInvoker {
	requestDef := GenReqDefForBatchCreateAndDeleteVaultTags()
	return &BatchCreateAndDeleteVaultTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) CopyBackup(request *model.CopyBackupRequest) (*model.CopyBackupResponse, error) {
	requestDef := GenReqDefForCopyBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyBackupResponse), nil
	}
}

func (c *CbrClient) CopyBackupInvoker(request *model.CopyBackupRequest) *CopyBackupInvoker {
	requestDef := GenReqDefForCopyBackup()
	return &CopyBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) CopyCheckpoint(request *model.CopyCheckpointRequest) (*model.CopyCheckpointResponse, error) {
	requestDef := GenReqDefForCopyCheckpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyCheckpointResponse), nil
	}
}

func (c *CbrClient) CopyCheckpointInvoker(request *model.CopyCheckpointRequest) *CopyCheckpointInvoker {
	requestDef := GenReqDefForCopyCheckpoint()
	return &CopyCheckpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) CreateCheckpoint(request *model.CreateCheckpointRequest) (*model.CreateCheckpointResponse, error) {
	requestDef := GenReqDefForCreateCheckpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCheckpointResponse), nil
	}
}

func (c *CbrClient) CreateCheckpointInvoker(request *model.CreateCheckpointRequest) *CreateCheckpointInvoker {
	requestDef := GenReqDefForCreateCheckpoint()
	return &CreateCheckpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) CreatePolicy(request *model.CreatePolicyRequest) (*model.CreatePolicyResponse, error) {
	requestDef := GenReqDefForCreatePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePolicyResponse), nil
	}
}

func (c *CbrClient) CreatePolicyInvoker(request *model.CreatePolicyRequest) *CreatePolicyInvoker {
	requestDef := GenReqDefForCreatePolicy()
	return &CreatePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) CreateVault(request *model.CreateVaultRequest) (*model.CreateVaultResponse, error) {
	requestDef := GenReqDefForCreateVault()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVaultResponse), nil
	}
}

func (c *CbrClient) CreateVaultInvoker(request *model.CreateVaultRequest) *CreateVaultInvoker {
	requestDef := GenReqDefForCreateVault()
	return &CreateVaultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) CreateVaultTags(request *model.CreateVaultTagsRequest) (*model.CreateVaultTagsResponse, error) {
	requestDef := GenReqDefForCreateVaultTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVaultTagsResponse), nil
	}
}

func (c *CbrClient) CreateVaultTagsInvoker(request *model.CreateVaultTagsRequest) *CreateVaultTagsInvoker {
	requestDef := GenReqDefForCreateVaultTags()
	return &CreateVaultTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) DeleteBackup(request *model.DeleteBackupRequest) (*model.DeleteBackupResponse, error) {
	requestDef := GenReqDefForDeleteBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackupResponse), nil
	}
}

func (c *CbrClient) DeleteBackupInvoker(request *model.DeleteBackupRequest) *DeleteBackupInvoker {
	requestDef := GenReqDefForDeleteBackup()
	return &DeleteBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) DeleteMember(request *model.DeleteMemberRequest) (*model.DeleteMemberResponse, error) {
	requestDef := GenReqDefForDeleteMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMemberResponse), nil
	}
}

func (c *CbrClient) DeleteMemberInvoker(request *model.DeleteMemberRequest) *DeleteMemberInvoker {
	requestDef := GenReqDefForDeleteMember()
	return &DeleteMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) DeletePolicy(request *model.DeletePolicyRequest) (*model.DeletePolicyResponse, error) {
	requestDef := GenReqDefForDeletePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePolicyResponse), nil
	}
}

func (c *CbrClient) DeletePolicyInvoker(request *model.DeletePolicyRequest) *DeletePolicyInvoker {
	requestDef := GenReqDefForDeletePolicy()
	return &DeletePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) DeleteVault(request *model.DeleteVaultRequest) (*model.DeleteVaultResponse, error) {
	requestDef := GenReqDefForDeleteVault()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVaultResponse), nil
	}
}

func (c *CbrClient) DeleteVaultInvoker(request *model.DeleteVaultRequest) *DeleteVaultInvoker {
	requestDef := GenReqDefForDeleteVault()
	return &DeleteVaultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) DeleteVaultTag(request *model.DeleteVaultTagRequest) (*model.DeleteVaultTagResponse, error) {
	requestDef := GenReqDefForDeleteVaultTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVaultTagResponse), nil
	}
}

func (c *CbrClient) DeleteVaultTagInvoker(request *model.DeleteVaultTagRequest) *DeleteVaultTagInvoker {
	requestDef := GenReqDefForDeleteVaultTag()
	return &DeleteVaultTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) DisassociateVaultPolicy(request *model.DisassociateVaultPolicyRequest) (*model.DisassociateVaultPolicyResponse, error) {
	requestDef := GenReqDefForDisassociateVaultPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateVaultPolicyResponse), nil
	}
}

func (c *CbrClient) DisassociateVaultPolicyInvoker(request *model.DisassociateVaultPolicyRequest) *DisassociateVaultPolicyInvoker {
	requestDef := GenReqDefForDisassociateVaultPolicy()
	return &DisassociateVaultPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) ImportBackup(request *model.ImportBackupRequest) (*model.ImportBackupResponse, error) {
	requestDef := GenReqDefForImportBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportBackupResponse), nil
	}
}

func (c *CbrClient) ImportBackupInvoker(request *model.ImportBackupRequest) *ImportBackupInvoker {
	requestDef := GenReqDefForImportBackup()
	return &ImportBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) ListBackups(request *model.ListBackupsRequest) (*model.ListBackupsResponse, error) {
	requestDef := GenReqDefForListBackups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackupsResponse), nil
	}
}

func (c *CbrClient) ListBackupsInvoker(request *model.ListBackupsRequest) *ListBackupsInvoker {
	requestDef := GenReqDefForListBackups()
	return &ListBackupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) ListOpLogs(request *model.ListOpLogsRequest) (*model.ListOpLogsResponse, error) {
	requestDef := GenReqDefForListOpLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOpLogsResponse), nil
	}
}

func (c *CbrClient) ListOpLogsInvoker(request *model.ListOpLogsRequest) *ListOpLogsInvoker {
	requestDef := GenReqDefForListOpLogs()
	return &ListOpLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) ListPolicies(request *model.ListPoliciesRequest) (*model.ListPoliciesResponse, error) {
	requestDef := GenReqDefForListPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPoliciesResponse), nil
	}
}

func (c *CbrClient) ListPoliciesInvoker(request *model.ListPoliciesRequest) *ListPoliciesInvoker {
	requestDef := GenReqDefForListPolicies()
	return &ListPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) ListProtectable(request *model.ListProtectableRequest) (*model.ListProtectableResponse, error) {
	requestDef := GenReqDefForListProtectable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProtectableResponse), nil
	}
}

func (c *CbrClient) ListProtectableInvoker(request *model.ListProtectableRequest) *ListProtectableInvoker {
	requestDef := GenReqDefForListProtectable()
	return &ListProtectableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) ListVault(request *model.ListVaultRequest) (*model.ListVaultResponse, error) {
	requestDef := GenReqDefForListVault()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVaultResponse), nil
	}
}

func (c *CbrClient) ListVaultInvoker(request *model.ListVaultRequest) *ListVaultInvoker {
	requestDef := GenReqDefForListVault()
	return &ListVaultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) MigrateVaultResource(request *model.MigrateVaultResourceRequest) (*model.MigrateVaultResourceResponse, error) {
	requestDef := GenReqDefForMigrateVaultResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MigrateVaultResourceResponse), nil
	}
}

func (c *CbrClient) MigrateVaultResourceInvoker(request *model.MigrateVaultResourceRequest) *MigrateVaultResourceInvoker {
	requestDef := GenReqDefForMigrateVaultResource()
	return &MigrateVaultResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) RemoveVaultResource(request *model.RemoveVaultResourceRequest) (*model.RemoveVaultResourceResponse, error) {
	requestDef := GenReqDefForRemoveVaultResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveVaultResourceResponse), nil
	}
}

func (c *CbrClient) RemoveVaultResourceInvoker(request *model.RemoveVaultResourceRequest) *RemoveVaultResourceInvoker {
	requestDef := GenReqDefForRemoveVaultResource()
	return &RemoveVaultResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) RestoreBackup(request *model.RestoreBackupRequest) (*model.RestoreBackupResponse, error) {
	requestDef := GenReqDefForRestoreBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreBackupResponse), nil
	}
}

func (c *CbrClient) RestoreBackupInvoker(request *model.RestoreBackupRequest) *RestoreBackupInvoker {
	requestDef := GenReqDefForRestoreBackup()
	return &RestoreBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) ShowBackup(request *model.ShowBackupRequest) (*model.ShowBackupResponse, error) {
	requestDef := GenReqDefForShowBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackupResponse), nil
	}
}

func (c *CbrClient) ShowBackupInvoker(request *model.ShowBackupRequest) *ShowBackupInvoker {
	requestDef := GenReqDefForShowBackup()
	return &ShowBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) ShowCheckpoint(request *model.ShowCheckpointRequest) (*model.ShowCheckpointResponse, error) {
	requestDef := GenReqDefForShowCheckpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCheckpointResponse), nil
	}
}

func (c *CbrClient) ShowCheckpointInvoker(request *model.ShowCheckpointRequest) *ShowCheckpointInvoker {
	requestDef := GenReqDefForShowCheckpoint()
	return &ShowCheckpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) ShowMemberDetail(request *model.ShowMemberDetailRequest) (*model.ShowMemberDetailResponse, error) {
	requestDef := GenReqDefForShowMemberDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMemberDetailResponse), nil
	}
}

func (c *CbrClient) ShowMemberDetailInvoker(request *model.ShowMemberDetailRequest) *ShowMemberDetailInvoker {
	requestDef := GenReqDefForShowMemberDetail()
	return &ShowMemberDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) ShowMembersDetail(request *model.ShowMembersDetailRequest) (*model.ShowMembersDetailResponse, error) {
	requestDef := GenReqDefForShowMembersDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMembersDetailResponse), nil
	}
}

func (c *CbrClient) ShowMembersDetailInvoker(request *model.ShowMembersDetailRequest) *ShowMembersDetailInvoker {
	requestDef := GenReqDefForShowMembersDetail()
	return &ShowMembersDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) ShowOpLog(request *model.ShowOpLogRequest) (*model.ShowOpLogResponse, error) {
	requestDef := GenReqDefForShowOpLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOpLogResponse), nil
	}
}

func (c *CbrClient) ShowOpLogInvoker(request *model.ShowOpLogRequest) *ShowOpLogInvoker {
	requestDef := GenReqDefForShowOpLog()
	return &ShowOpLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) ShowPolicy(request *model.ShowPolicyRequest) (*model.ShowPolicyResponse, error) {
	requestDef := GenReqDefForShowPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPolicyResponse), nil
	}
}

func (c *CbrClient) ShowPolicyInvoker(request *model.ShowPolicyRequest) *ShowPolicyInvoker {
	requestDef := GenReqDefForShowPolicy()
	return &ShowPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) ShowProtectable(request *model.ShowProtectableRequest) (*model.ShowProtectableResponse, error) {
	requestDef := GenReqDefForShowProtectable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProtectableResponse), nil
	}
}

func (c *CbrClient) ShowProtectableInvoker(request *model.ShowProtectableRequest) *ShowProtectableInvoker {
	requestDef := GenReqDefForShowProtectable()
	return &ShowProtectableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) ShowReplicationCapabilities(request *model.ShowReplicationCapabilitiesRequest) (*model.ShowReplicationCapabilitiesResponse, error) {
	requestDef := GenReqDefForShowReplicationCapabilities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReplicationCapabilitiesResponse), nil
	}
}

func (c *CbrClient) ShowReplicationCapabilitiesInvoker(request *model.ShowReplicationCapabilitiesRequest) *ShowReplicationCapabilitiesInvoker {
	requestDef := GenReqDefForShowReplicationCapabilities()
	return &ShowReplicationCapabilitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) ShowVault(request *model.ShowVaultRequest) (*model.ShowVaultResponse, error) {
	requestDef := GenReqDefForShowVault()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVaultResponse), nil
	}
}

func (c *CbrClient) ShowVaultInvoker(request *model.ShowVaultRequest) *ShowVaultInvoker {
	requestDef := GenReqDefForShowVault()
	return &ShowVaultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) ShowVaultProjectTag(request *model.ShowVaultProjectTagRequest) (*model.ShowVaultProjectTagResponse, error) {
	requestDef := GenReqDefForShowVaultProjectTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVaultProjectTagResponse), nil
	}
}

func (c *CbrClient) ShowVaultProjectTagInvoker(request *model.ShowVaultProjectTagRequest) *ShowVaultProjectTagInvoker {
	requestDef := GenReqDefForShowVaultProjectTag()
	return &ShowVaultProjectTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) ShowVaultResourceInstances(request *model.ShowVaultResourceInstancesRequest) (*model.ShowVaultResourceInstancesResponse, error) {
	requestDef := GenReqDefForShowVaultResourceInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVaultResourceInstancesResponse), nil
	}
}

func (c *CbrClient) ShowVaultResourceInstancesInvoker(request *model.ShowVaultResourceInstancesRequest) *ShowVaultResourceInstancesInvoker {
	requestDef := GenReqDefForShowVaultResourceInstances()
	return &ShowVaultResourceInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) ShowVaultTag(request *model.ShowVaultTagRequest) (*model.ShowVaultTagResponse, error) {
	requestDef := GenReqDefForShowVaultTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVaultTagResponse), nil
	}
}

func (c *CbrClient) ShowVaultTagInvoker(request *model.ShowVaultTagRequest) *ShowVaultTagInvoker {
	requestDef := GenReqDefForShowVaultTag()
	return &ShowVaultTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) UpdateMemberStatus(request *model.UpdateMemberStatusRequest) (*model.UpdateMemberStatusResponse, error) {
	requestDef := GenReqDefForUpdateMemberStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMemberStatusResponse), nil
	}
}

func (c *CbrClient) UpdateMemberStatusInvoker(request *model.UpdateMemberStatusRequest) *UpdateMemberStatusInvoker {
	requestDef := GenReqDefForUpdateMemberStatus()
	return &UpdateMemberStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) UpdatePolicy(request *model.UpdatePolicyRequest) (*model.UpdatePolicyResponse, error) {
	requestDef := GenReqDefForUpdatePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePolicyResponse), nil
	}
}

func (c *CbrClient) UpdatePolicyInvoker(request *model.UpdatePolicyRequest) *UpdatePolicyInvoker {
	requestDef := GenReqDefForUpdatePolicy()
	return &UpdatePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CbrClient) UpdateVault(request *model.UpdateVaultRequest) (*model.UpdateVaultResponse, error) {
	requestDef := GenReqDefForUpdateVault()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVaultResponse), nil
	}
}

func (c *CbrClient) UpdateVaultInvoker(request *model.UpdateVaultRequest) *UpdateVaultInvoker {
	requestDef := GenReqDefForUpdateVault()
	return &UpdateVaultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
