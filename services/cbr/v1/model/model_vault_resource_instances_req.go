package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type VaultResourceInstancesReq struct {
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`

	Tags *[]TagsReq `json:"tags,omitempty"`

	TagsAny *[]TagsReq `json:"tags_any,omitempty"`

	NotTags *[]TagsReq `json:"not_tags,omitempty"`

	NotTagsAny *[]TagsReq `json:"not_tags_any,omitempty"`

	SysTags *[]SysTags `json:"sys_tags,omitempty"`

	Limit *string `json:"limit,omitempty"`

	Offset *string `json:"offset,omitempty"`

	Action string `json:"action"`

	Matches *[]Match `json:"matches,omitempty"`

	CloudType *VaultResourceInstancesReqCloudType `json:"cloud_type,omitempty"`

	ObjectType *VaultResourceInstancesReqObjectType `json:"object_type,omitempty"`
}

func (o VaultResourceInstancesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultResourceInstancesReq struct{}"
	}

	return strings.Join([]string{"VaultResourceInstancesReq", string(data)}, " ")
}

type VaultResourceInstancesReqCloudType struct {
	value string
}

type VaultResourceInstancesReqCloudTypeEnum struct {
	PUBLIC VaultResourceInstancesReqCloudType
	HYBRID VaultResourceInstancesReqCloudType
}

func GetVaultResourceInstancesReqCloudTypeEnum() VaultResourceInstancesReqCloudTypeEnum {
	return VaultResourceInstancesReqCloudTypeEnum{
		PUBLIC: VaultResourceInstancesReqCloudType{
			value: "public",
		},
		HYBRID: VaultResourceInstancesReqCloudType{
			value: " hybrid",
		},
	}
}

func (c VaultResourceInstancesReqCloudType) Value() string {
	return c.value
}

func (c VaultResourceInstancesReqCloudType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VaultResourceInstancesReqCloudType) UnmarshalJSON(b []byte) error {
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

type VaultResourceInstancesReqObjectType struct {
	value string
}

type VaultResourceInstancesReqObjectTypeEnum struct {
	SERVER VaultResourceInstancesReqObjectType
	DISK   VaultResourceInstancesReqObjectType
}

func GetVaultResourceInstancesReqObjectTypeEnum() VaultResourceInstancesReqObjectTypeEnum {
	return VaultResourceInstancesReqObjectTypeEnum{
		SERVER: VaultResourceInstancesReqObjectType{
			value: "server",
		},
		DISK: VaultResourceInstancesReqObjectType{
			value: "disk",
		},
	}
}

func (c VaultResourceInstancesReqObjectType) Value() string {
	return c.value
}

func (c VaultResourceInstancesReqObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VaultResourceInstancesReqObjectType) UnmarshalJSON(b []byte) error {
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
