package v2

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/evs/v2/model"
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
)

type EvsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewEvsClient(hcClient *http_client.HcHttpClient) *EvsClient {
	return &EvsClient{HcClient: hcClient}
}

func EvsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *EvsClient) BatchCreateVolumeTags(request *model.BatchCreateVolumeTagsRequest) (*model.BatchCreateVolumeTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateVolumeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateVolumeTagsResponse), nil
	}
}

func (c *EvsClient) BatchCreateVolumeTagsInvoker(request *model.BatchCreateVolumeTagsRequest) *BatchCreateVolumeTagsInvoker {
	requestDef := GenReqDefForBatchCreateVolumeTags()
	return &BatchCreateVolumeTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) BatchDeleteVolumeTags(request *model.BatchDeleteVolumeTagsRequest) (*model.BatchDeleteVolumeTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteVolumeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteVolumeTagsResponse), nil
	}
}

func (c *EvsClient) BatchDeleteVolumeTagsInvoker(request *model.BatchDeleteVolumeTagsRequest) *BatchDeleteVolumeTagsInvoker {
	requestDef := GenReqDefForBatchDeleteVolumeTags()
	return &BatchDeleteVolumeTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) CinderAcceptVolumeTransfer(request *model.CinderAcceptVolumeTransferRequest) (*model.CinderAcceptVolumeTransferResponse, error) {
	requestDef := GenReqDefForCinderAcceptVolumeTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderAcceptVolumeTransferResponse), nil
	}
}

func (c *EvsClient) CinderAcceptVolumeTransferInvoker(request *model.CinderAcceptVolumeTransferRequest) *CinderAcceptVolumeTransferInvoker {
	requestDef := GenReqDefForCinderAcceptVolumeTransfer()
	return &CinderAcceptVolumeTransferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) CinderCreateVolumeTransfer(request *model.CinderCreateVolumeTransferRequest) (*model.CinderCreateVolumeTransferResponse, error) {
	requestDef := GenReqDefForCinderCreateVolumeTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderCreateVolumeTransferResponse), nil
	}
}

func (c *EvsClient) CinderCreateVolumeTransferInvoker(request *model.CinderCreateVolumeTransferRequest) *CinderCreateVolumeTransferInvoker {
	requestDef := GenReqDefForCinderCreateVolumeTransfer()
	return &CinderCreateVolumeTransferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) CinderDeleteVolumeTransfer(request *model.CinderDeleteVolumeTransferRequest) (*model.CinderDeleteVolumeTransferResponse, error) {
	requestDef := GenReqDefForCinderDeleteVolumeTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderDeleteVolumeTransferResponse), nil
	}
}

func (c *EvsClient) CinderDeleteVolumeTransferInvoker(request *model.CinderDeleteVolumeTransferRequest) *CinderDeleteVolumeTransferInvoker {
	requestDef := GenReqDefForCinderDeleteVolumeTransfer()
	return &CinderDeleteVolumeTransferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) CinderListAvailabilityZones(request *model.CinderListAvailabilityZonesRequest) (*model.CinderListAvailabilityZonesResponse, error) {
	requestDef := GenReqDefForCinderListAvailabilityZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderListAvailabilityZonesResponse), nil
	}
}

func (c *EvsClient) CinderListAvailabilityZonesInvoker(request *model.CinderListAvailabilityZonesRequest) *CinderListAvailabilityZonesInvoker {
	requestDef := GenReqDefForCinderListAvailabilityZones()
	return &CinderListAvailabilityZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) CinderListQuotas(request *model.CinderListQuotasRequest) (*model.CinderListQuotasResponse, error) {
	requestDef := GenReqDefForCinderListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderListQuotasResponse), nil
	}
}

func (c *EvsClient) CinderListQuotasInvoker(request *model.CinderListQuotasRequest) *CinderListQuotasInvoker {
	requestDef := GenReqDefForCinderListQuotas()
	return &CinderListQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) CinderListVolumeTransfers(request *model.CinderListVolumeTransfersRequest) (*model.CinderListVolumeTransfersResponse, error) {
	requestDef := GenReqDefForCinderListVolumeTransfers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderListVolumeTransfersResponse), nil
	}
}

func (c *EvsClient) CinderListVolumeTransfersInvoker(request *model.CinderListVolumeTransfersRequest) *CinderListVolumeTransfersInvoker {
	requestDef := GenReqDefForCinderListVolumeTransfers()
	return &CinderListVolumeTransfersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) CinderListVolumeTypes(request *model.CinderListVolumeTypesRequest) (*model.CinderListVolumeTypesResponse, error) {
	requestDef := GenReqDefForCinderListVolumeTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderListVolumeTypesResponse), nil
	}
}

func (c *EvsClient) CinderListVolumeTypesInvoker(request *model.CinderListVolumeTypesRequest) *CinderListVolumeTypesInvoker {
	requestDef := GenReqDefForCinderListVolumeTypes()
	return &CinderListVolumeTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) CinderShowVolumeTransfer(request *model.CinderShowVolumeTransferRequest) (*model.CinderShowVolumeTransferResponse, error) {
	requestDef := GenReqDefForCinderShowVolumeTransfer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CinderShowVolumeTransferResponse), nil
	}
}

func (c *EvsClient) CinderShowVolumeTransferInvoker(request *model.CinderShowVolumeTransferRequest) *CinderShowVolumeTransferInvoker {
	requestDef := GenReqDefForCinderShowVolumeTransfer()
	return &CinderShowVolumeTransferInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) CreateSnapshot(request *model.CreateSnapshotRequest) (*model.CreateSnapshotResponse, error) {
	requestDef := GenReqDefForCreateSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSnapshotResponse), nil
	}
}

func (c *EvsClient) CreateSnapshotInvoker(request *model.CreateSnapshotRequest) *CreateSnapshotInvoker {
	requestDef := GenReqDefForCreateSnapshot()
	return &CreateSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) CreateVolume(request *model.CreateVolumeRequest) (*model.CreateVolumeResponse, error) {
	requestDef := GenReqDefForCreateVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVolumeResponse), nil
	}
}

func (c *EvsClient) CreateVolumeInvoker(request *model.CreateVolumeRequest) *CreateVolumeInvoker {
	requestDef := GenReqDefForCreateVolume()
	return &CreateVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) DeleteSnapshot(request *model.DeleteSnapshotRequest) (*model.DeleteSnapshotResponse, error) {
	requestDef := GenReqDefForDeleteSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSnapshotResponse), nil
	}
}

func (c *EvsClient) DeleteSnapshotInvoker(request *model.DeleteSnapshotRequest) *DeleteSnapshotInvoker {
	requestDef := GenReqDefForDeleteSnapshot()
	return &DeleteSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) DeleteVolume(request *model.DeleteVolumeRequest) (*model.DeleteVolumeResponse, error) {
	requestDef := GenReqDefForDeleteVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVolumeResponse), nil
	}
}

func (c *EvsClient) DeleteVolumeInvoker(request *model.DeleteVolumeRequest) *DeleteVolumeInvoker {
	requestDef := GenReqDefForDeleteVolume()
	return &DeleteVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) ListSnapshots(request *model.ListSnapshotsRequest) (*model.ListSnapshotsResponse, error) {
	requestDef := GenReqDefForListSnapshots()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotsResponse), nil
	}
}

func (c *EvsClient) ListSnapshotsInvoker(request *model.ListSnapshotsRequest) *ListSnapshotsInvoker {
	requestDef := GenReqDefForListSnapshots()
	return &ListSnapshotsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) ListVolumeTags(request *model.ListVolumeTagsRequest) (*model.ListVolumeTagsResponse, error) {
	requestDef := GenReqDefForListVolumeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVolumeTagsResponse), nil
	}
}

func (c *EvsClient) ListVolumeTagsInvoker(request *model.ListVolumeTagsRequest) *ListVolumeTagsInvoker {
	requestDef := GenReqDefForListVolumeTags()
	return &ListVolumeTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) ListVolumes(request *model.ListVolumesRequest) (*model.ListVolumesResponse, error) {
	requestDef := GenReqDefForListVolumes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVolumesResponse), nil
	}
}

func (c *EvsClient) ListVolumesInvoker(request *model.ListVolumesRequest) *ListVolumesInvoker {
	requestDef := GenReqDefForListVolumes()
	return &ListVolumesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) ListVolumesByTags(request *model.ListVolumesByTagsRequest) (*model.ListVolumesByTagsResponse, error) {
	requestDef := GenReqDefForListVolumesByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVolumesByTagsResponse), nil
	}
}

func (c *EvsClient) ListVolumesByTagsInvoker(request *model.ListVolumesByTagsRequest) *ListVolumesByTagsInvoker {
	requestDef := GenReqDefForListVolumesByTags()
	return &ListVolumesByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) ResizeVolume(request *model.ResizeVolumeRequest) (*model.ResizeVolumeResponse, error) {
	requestDef := GenReqDefForResizeVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeVolumeResponse), nil
	}
}

func (c *EvsClient) ResizeVolumeInvoker(request *model.ResizeVolumeRequest) *ResizeVolumeInvoker {
	requestDef := GenReqDefForResizeVolume()
	return &ResizeVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) RollbackSnapshot(request *model.RollbackSnapshotRequest) (*model.RollbackSnapshotResponse, error) {
	requestDef := GenReqDefForRollbackSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RollbackSnapshotResponse), nil
	}
}

func (c *EvsClient) RollbackSnapshotInvoker(request *model.RollbackSnapshotRequest) *RollbackSnapshotInvoker {
	requestDef := GenReqDefForRollbackSnapshot()
	return &RollbackSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) ShowJob(request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	requestDef := GenReqDefForShowJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResponse), nil
	}
}

func (c *EvsClient) ShowJobInvoker(request *model.ShowJobRequest) *ShowJobInvoker {
	requestDef := GenReqDefForShowJob()
	return &ShowJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) ShowSnapshot(request *model.ShowSnapshotRequest) (*model.ShowSnapshotResponse, error) {
	requestDef := GenReqDefForShowSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSnapshotResponse), nil
	}
}

func (c *EvsClient) ShowSnapshotInvoker(request *model.ShowSnapshotRequest) *ShowSnapshotInvoker {
	requestDef := GenReqDefForShowSnapshot()
	return &ShowSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) ShowVolume(request *model.ShowVolumeRequest) (*model.ShowVolumeResponse, error) {
	requestDef := GenReqDefForShowVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVolumeResponse), nil
	}
}

func (c *EvsClient) ShowVolumeInvoker(request *model.ShowVolumeRequest) *ShowVolumeInvoker {
	requestDef := GenReqDefForShowVolume()
	return &ShowVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) ShowVolumeTags(request *model.ShowVolumeTagsRequest) (*model.ShowVolumeTagsResponse, error) {
	requestDef := GenReqDefForShowVolumeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVolumeTagsResponse), nil
	}
}

func (c *EvsClient) ShowVolumeTagsInvoker(request *model.ShowVolumeTagsRequest) *ShowVolumeTagsInvoker {
	requestDef := GenReqDefForShowVolumeTags()
	return &ShowVolumeTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) UpdateSnapshot(request *model.UpdateSnapshotRequest) (*model.UpdateSnapshotResponse, error) {
	requestDef := GenReqDefForUpdateSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSnapshotResponse), nil
	}
}

func (c *EvsClient) UpdateSnapshotInvoker(request *model.UpdateSnapshotRequest) *UpdateSnapshotInvoker {
	requestDef := GenReqDefForUpdateSnapshot()
	return &UpdateSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) UpdateVolume(request *model.UpdateVolumeRequest) (*model.UpdateVolumeResponse, error) {
	requestDef := GenReqDefForUpdateVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVolumeResponse), nil
	}
}

func (c *EvsClient) UpdateVolumeInvoker(request *model.UpdateVolumeRequest) *UpdateVolumeInvoker {
	requestDef := GenReqDefForUpdateVolume()
	return &UpdateVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) ListVersions(request *model.ListVersionsRequest) (*model.ListVersionsResponse, error) {
	requestDef := GenReqDefForListVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVersionsResponse), nil
	}
}

func (c *EvsClient) ListVersionsInvoker(request *model.ListVersionsRequest) *ListVersionsInvoker {
	requestDef := GenReqDefForListVersions()
	return &ListVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *EvsClient) ShowVersion(request *model.ShowVersionRequest) (*model.ShowVersionResponse, error) {
	requestDef := GenReqDefForShowVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVersionResponse), nil
	}
}

func (c *EvsClient) ShowVersionInvoker(request *model.ShowVersionRequest) *ShowVersionInvoker {
	requestDef := GenReqDefForShowVersion()
	return &ShowVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
