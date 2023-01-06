package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ResourceResp struct {
	ExtraInfo *ResourceExtraInfo `json:"extra_info,omitempty"`

	Id string `json:"id"`

	Name string `json:"name"`

	ProtectStatus *ResourceRespProtectStatus `json:"protect_status,omitempty"`

	Size *int32 `json:"size,omitempty"`

	Type string `json:"type"`

	BackupSize *int32 `json:"backup_size,omitempty"`

	BackupCount *int32 `json:"backup_count,omitempty"`
}

func (o ResourceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceResp struct{}"
	}

	return strings.Join([]string{"ResourceResp", string(data)}, " ")
}

type ResourceRespProtectStatus struct {
	value string
}

type ResourceRespProtectStatusEnum struct {
	AVAILABLE  ResourceRespProtectStatus
	ERROR      ResourceRespProtectStatus
	PROTECTING ResourceRespProtectStatus
	RESTORING  ResourceRespProtectStatus
	REMOVING   ResourceRespProtectStatus
}

func GetResourceRespProtectStatusEnum() ResourceRespProtectStatusEnum {
	return ResourceRespProtectStatusEnum{
		AVAILABLE: ResourceRespProtectStatus{
			value: "available",
		},
		ERROR: ResourceRespProtectStatus{
			value: "error",
		},
		PROTECTING: ResourceRespProtectStatus{
			value: "protecting",
		},
		RESTORING: ResourceRespProtectStatus{
			value: "restoring",
		},
		REMOVING: ResourceRespProtectStatus{
			value: "removing",
		},
	}
}

func (c ResourceRespProtectStatus) Value() string {
	return c.value
}

func (c ResourceRespProtectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceRespProtectStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
