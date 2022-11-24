package v3

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/vpc/v3/model"
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
)

type VpcClient struct {
	HcClient *http_client.HcHttpClient
}

func NewVpcClient(hcClient *http_client.HcHttpClient) *VpcClient {
	return &VpcClient{HcClient: hcClient}
}

func VpcClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *VpcClient) BatchCreateSubNetworkInterface(request *model.BatchCreateSubNetworkInterfaceRequest) (*model.BatchCreateSubNetworkInterfaceResponse, error) {
	requestDef := GenReqDefForBatchCreateSubNetworkInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateSubNetworkInterfaceResponse), nil
	}
}

func (c *VpcClient) BatchCreateSubNetworkInterfaceInvoker(request *model.BatchCreateSubNetworkInterfaceRequest) *BatchCreateSubNetworkInterfaceInvoker {
	requestDef := GenReqDefForBatchCreateSubNetworkInterface()
	return &BatchCreateSubNetworkInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) CreateSecurityGroup(request *model.CreateSecurityGroupRequest) (*model.CreateSecurityGroupResponse, error) {
	requestDef := GenReqDefForCreateSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityGroupResponse), nil
	}
}

func (c *VpcClient) CreateSecurityGroupInvoker(request *model.CreateSecurityGroupRequest) *CreateSecurityGroupInvoker {
	requestDef := GenReqDefForCreateSecurityGroup()
	return &CreateSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) CreateSecurityGroupRule(request *model.CreateSecurityGroupRuleRequest) (*model.CreateSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForCreateSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityGroupRuleResponse), nil
	}
}

func (c *VpcClient) CreateSecurityGroupRuleInvoker(request *model.CreateSecurityGroupRuleRequest) *CreateSecurityGroupRuleInvoker {
	requestDef := GenReqDefForCreateSecurityGroupRule()
	return &CreateSecurityGroupRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) CreateSubNetworkInterface(request *model.CreateSubNetworkInterfaceRequest) (*model.CreateSubNetworkInterfaceResponse, error) {
	requestDef := GenReqDefForCreateSubNetworkInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubNetworkInterfaceResponse), nil
	}
}

func (c *VpcClient) CreateSubNetworkInterfaceInvoker(request *model.CreateSubNetworkInterfaceRequest) *CreateSubNetworkInterfaceInvoker {
	requestDef := GenReqDefForCreateSubNetworkInterface()
	return &CreateSubNetworkInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) DeleteSecurityGroup(request *model.DeleteSecurityGroupRequest) (*model.DeleteSecurityGroupResponse, error) {
	requestDef := GenReqDefForDeleteSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecurityGroupResponse), nil
	}
}

func (c *VpcClient) DeleteSecurityGroupInvoker(request *model.DeleteSecurityGroupRequest) *DeleteSecurityGroupInvoker {
	requestDef := GenReqDefForDeleteSecurityGroup()
	return &DeleteSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) DeleteSecurityGroupRule(request *model.DeleteSecurityGroupRuleRequest) (*model.DeleteSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForDeleteSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecurityGroupRuleResponse), nil
	}
}

func (c *VpcClient) DeleteSecurityGroupRuleInvoker(request *model.DeleteSecurityGroupRuleRequest) *DeleteSecurityGroupRuleInvoker {
	requestDef := GenReqDefForDeleteSecurityGroupRule()
	return &DeleteSecurityGroupRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) DeleteSubNetworkInterface(request *model.DeleteSubNetworkInterfaceRequest) (*model.DeleteSubNetworkInterfaceResponse, error) {
	requestDef := GenReqDefForDeleteSubNetworkInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSubNetworkInterfaceResponse), nil
	}
}

func (c *VpcClient) DeleteSubNetworkInterfaceInvoker(request *model.DeleteSubNetworkInterfaceRequest) *DeleteSubNetworkInterfaceInvoker {
	requestDef := GenReqDefForDeleteSubNetworkInterface()
	return &DeleteSubNetworkInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ListSecurityGroupRules(request *model.ListSecurityGroupRulesRequest) (*model.ListSecurityGroupRulesResponse, error) {
	requestDef := GenReqDefForListSecurityGroupRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityGroupRulesResponse), nil
	}
}

func (c *VpcClient) ListSecurityGroupRulesInvoker(request *model.ListSecurityGroupRulesRequest) *ListSecurityGroupRulesInvoker {
	requestDef := GenReqDefForListSecurityGroupRules()
	return &ListSecurityGroupRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ListSecurityGroups(request *model.ListSecurityGroupsRequest) (*model.ListSecurityGroupsResponse, error) {
	requestDef := GenReqDefForListSecurityGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityGroupsResponse), nil
	}
}

func (c *VpcClient) ListSecurityGroupsInvoker(request *model.ListSecurityGroupsRequest) *ListSecurityGroupsInvoker {
	requestDef := GenReqDefForListSecurityGroups()
	return &ListSecurityGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ListSubNetworkInterfaces(request *model.ListSubNetworkInterfacesRequest) (*model.ListSubNetworkInterfacesResponse, error) {
	requestDef := GenReqDefForListSubNetworkInterfaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubNetworkInterfacesResponse), nil
	}
}

func (c *VpcClient) ListSubNetworkInterfacesInvoker(request *model.ListSubNetworkInterfacesRequest) *ListSubNetworkInterfacesInvoker {
	requestDef := GenReqDefForListSubNetworkInterfaces()
	return &ListSubNetworkInterfacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) MigrateSubNetworkInterface(request *model.MigrateSubNetworkInterfaceRequest) (*model.MigrateSubNetworkInterfaceResponse, error) {
	requestDef := GenReqDefForMigrateSubNetworkInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MigrateSubNetworkInterfaceResponse), nil
	}
}

func (c *VpcClient) MigrateSubNetworkInterfaceInvoker(request *model.MigrateSubNetworkInterfaceRequest) *MigrateSubNetworkInterfaceInvoker {
	requestDef := GenReqDefForMigrateSubNetworkInterface()
	return &MigrateSubNetworkInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ShowSecurityGroup(request *model.ShowSecurityGroupRequest) (*model.ShowSecurityGroupResponse, error) {
	requestDef := GenReqDefForShowSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecurityGroupResponse), nil
	}
}

func (c *VpcClient) ShowSecurityGroupInvoker(request *model.ShowSecurityGroupRequest) *ShowSecurityGroupInvoker {
	requestDef := GenReqDefForShowSecurityGroup()
	return &ShowSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ShowSecurityGroupRule(request *model.ShowSecurityGroupRuleRequest) (*model.ShowSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForShowSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecurityGroupRuleResponse), nil
	}
}

func (c *VpcClient) ShowSecurityGroupRuleInvoker(request *model.ShowSecurityGroupRuleRequest) *ShowSecurityGroupRuleInvoker {
	requestDef := GenReqDefForShowSecurityGroupRule()
	return &ShowSecurityGroupRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ShowSubNetworkInterface(request *model.ShowSubNetworkInterfaceRequest) (*model.ShowSubNetworkInterfaceResponse, error) {
	requestDef := GenReqDefForShowSubNetworkInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSubNetworkInterfaceResponse), nil
	}
}

func (c *VpcClient) ShowSubNetworkInterfaceInvoker(request *model.ShowSubNetworkInterfaceRequest) *ShowSubNetworkInterfaceInvoker {
	requestDef := GenReqDefForShowSubNetworkInterface()
	return &ShowSubNetworkInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ShowSubNetworkInterfacesQuantity(request *model.ShowSubNetworkInterfacesQuantityRequest) (*model.ShowSubNetworkInterfacesQuantityResponse, error) {
	requestDef := GenReqDefForShowSubNetworkInterfacesQuantity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSubNetworkInterfacesQuantityResponse), nil
	}
}

func (c *VpcClient) ShowSubNetworkInterfacesQuantityInvoker(request *model.ShowSubNetworkInterfacesQuantityRequest) *ShowSubNetworkInterfacesQuantityInvoker {
	requestDef := GenReqDefForShowSubNetworkInterfacesQuantity()
	return &ShowSubNetworkInterfacesQuantityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) UpdateSecurityGroup(request *model.UpdateSecurityGroupRequest) (*model.UpdateSecurityGroupResponse, error) {
	requestDef := GenReqDefForUpdateSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecurityGroupResponse), nil
	}
}

func (c *VpcClient) UpdateSecurityGroupInvoker(request *model.UpdateSecurityGroupRequest) *UpdateSecurityGroupInvoker {
	requestDef := GenReqDefForUpdateSecurityGroup()
	return &UpdateSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) UpdateSubNetworkInterface(request *model.UpdateSubNetworkInterfaceRequest) (*model.UpdateSubNetworkInterfaceResponse, error) {
	requestDef := GenReqDefForUpdateSubNetworkInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSubNetworkInterfaceResponse), nil
	}
}

func (c *VpcClient) UpdateSubNetworkInterfaceInvoker(request *model.UpdateSubNetworkInterfaceRequest) *UpdateSubNetworkInterfaceInvoker {
	requestDef := GenReqDefForUpdateSubNetworkInterface()
	return &UpdateSubNetworkInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) CreateAddressGroup(request *model.CreateAddressGroupRequest) (*model.CreateAddressGroupResponse, error) {
	requestDef := GenReqDefForCreateAddressGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAddressGroupResponse), nil
	}
}

func (c *VpcClient) CreateAddressGroupInvoker(request *model.CreateAddressGroupRequest) *CreateAddressGroupInvoker {
	requestDef := GenReqDefForCreateAddressGroup()
	return &CreateAddressGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) DeleteAddressGroup(request *model.DeleteAddressGroupRequest) (*model.DeleteAddressGroupResponse, error) {
	requestDef := GenReqDefForDeleteAddressGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAddressGroupResponse), nil
	}
}

func (c *VpcClient) DeleteAddressGroupInvoker(request *model.DeleteAddressGroupRequest) *DeleteAddressGroupInvoker {
	requestDef := GenReqDefForDeleteAddressGroup()
	return &DeleteAddressGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) DeleteIpAddressGroupForce(request *model.DeleteIpAddressGroupForceRequest) (*model.DeleteIpAddressGroupForceResponse, error) {
	requestDef := GenReqDefForDeleteIpAddressGroupForce()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIpAddressGroupForceResponse), nil
	}
}

func (c *VpcClient) DeleteIpAddressGroupForceInvoker(request *model.DeleteIpAddressGroupForceRequest) *DeleteIpAddressGroupForceInvoker {
	requestDef := GenReqDefForDeleteIpAddressGroupForce()
	return &DeleteIpAddressGroupForceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ListAddressGroup(request *model.ListAddressGroupRequest) (*model.ListAddressGroupResponse, error) {
	requestDef := GenReqDefForListAddressGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAddressGroupResponse), nil
	}
}

func (c *VpcClient) ListAddressGroupInvoker(request *model.ListAddressGroupRequest) *ListAddressGroupInvoker {
	requestDef := GenReqDefForListAddressGroup()
	return &ListAddressGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ShowAddressGroup(request *model.ShowAddressGroupRequest) (*model.ShowAddressGroupResponse, error) {
	requestDef := GenReqDefForShowAddressGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAddressGroupResponse), nil
	}
}

func (c *VpcClient) ShowAddressGroupInvoker(request *model.ShowAddressGroupRequest) *ShowAddressGroupInvoker {
	requestDef := GenReqDefForShowAddressGroup()
	return &ShowAddressGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) UpdateAddressGroup(request *model.UpdateAddressGroupRequest) (*model.UpdateAddressGroupResponse, error) {
	requestDef := GenReqDefForUpdateAddressGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAddressGroupResponse), nil
	}
}

func (c *VpcClient) UpdateAddressGroupInvoker(request *model.UpdateAddressGroupRequest) *UpdateAddressGroupInvoker {
	requestDef := GenReqDefForUpdateAddressGroup()
	return &UpdateAddressGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) AddVpcExtendCidr(request *model.AddVpcExtendCidrRequest) (*model.AddVpcExtendCidrResponse, error) {
	requestDef := GenReqDefForAddVpcExtendCidr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddVpcExtendCidrResponse), nil
	}
}

func (c *VpcClient) AddVpcExtendCidrInvoker(request *model.AddVpcExtendCidrRequest) *AddVpcExtendCidrInvoker {
	requestDef := GenReqDefForAddVpcExtendCidr()
	return &AddVpcExtendCidrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) CreateVpc(request *model.CreateVpcRequest) (*model.CreateVpcResponse, error) {
	requestDef := GenReqDefForCreateVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpcResponse), nil
	}
}

func (c *VpcClient) CreateVpcInvoker(request *model.CreateVpcRequest) *CreateVpcInvoker {
	requestDef := GenReqDefForCreateVpc()
	return &CreateVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) DeleteVpc(request *model.DeleteVpcRequest) (*model.DeleteVpcResponse, error) {
	requestDef := GenReqDefForDeleteVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpcResponse), nil
	}
}

func (c *VpcClient) DeleteVpcInvoker(request *model.DeleteVpcRequest) *DeleteVpcInvoker {
	requestDef := GenReqDefForDeleteVpc()
	return &DeleteVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ListVpcs(request *model.ListVpcsRequest) (*model.ListVpcsResponse, error) {
	requestDef := GenReqDefForListVpcs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcsResponse), nil
	}
}

func (c *VpcClient) ListVpcsInvoker(request *model.ListVpcsRequest) *ListVpcsInvoker {
	requestDef := GenReqDefForListVpcs()
	return &ListVpcsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) RemoveVpcExtendCidr(request *model.RemoveVpcExtendCidrRequest) (*model.RemoveVpcExtendCidrResponse, error) {
	requestDef := GenReqDefForRemoveVpcExtendCidr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveVpcExtendCidrResponse), nil
	}
}

func (c *VpcClient) RemoveVpcExtendCidrInvoker(request *model.RemoveVpcExtendCidrRequest) *RemoveVpcExtendCidrInvoker {
	requestDef := GenReqDefForRemoveVpcExtendCidr()
	return &RemoveVpcExtendCidrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ShowVpc(request *model.ShowVpcRequest) (*model.ShowVpcResponse, error) {
	requestDef := GenReqDefForShowVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVpcResponse), nil
	}
}

func (c *VpcClient) ShowVpcInvoker(request *model.ShowVpcRequest) *ShowVpcInvoker {
	requestDef := GenReqDefForShowVpc()
	return &ShowVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) UpdateVpc(request *model.UpdateVpcRequest) (*model.UpdateVpcResponse, error) {
	requestDef := GenReqDefForUpdateVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpcResponse), nil
	}
}

func (c *VpcClient) UpdateVpcInvoker(request *model.UpdateVpcRequest) *UpdateVpcInvoker {
	requestDef := GenReqDefForUpdateVpc()
	return &UpdateVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
