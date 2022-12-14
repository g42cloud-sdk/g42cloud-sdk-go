package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/sdktime"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AddressGroup struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	IpSet []string `json:"ip_set"`

	IpVersion int32 `json:"ip_version"`

	CreatedAt *sdktime.SdkTime `json:"created_at"`

	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	TenantId string `json:"tenant_id"`
}

func (o AddressGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressGroup struct{}"
	}

	return strings.Join([]string{"AddressGroup", string(data)}, " ")
}
