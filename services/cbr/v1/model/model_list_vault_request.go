package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListVaultRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Name *string `json:"name,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	CloudType *ListVaultRequestCloudType `json:"cloud_type,omitempty"`

	ProtectType *ListVaultRequestProtectType `json:"protect_type,omitempty"`

	ObjectType *string `json:"object_type,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Id *string `json:"id,omitempty"`

	PolicyId *string `json:"policy_id,omitempty"`

	Status *string `json:"status,omitempty"`

	ResourceIds *string `json:"resource_ids,omitempty"`
}

func (o ListVaultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVaultRequest struct{}"
	}

	return strings.Join([]string{"ListVaultRequest", string(data)}, " ")
}

type ListVaultRequestCloudType struct {
	value string
}

type ListVaultRequestCloudTypeEnum struct {
	PUBLIC ListVaultRequestCloudType
	HYBRID ListVaultRequestCloudType
}

func GetListVaultRequestCloudTypeEnum() ListVaultRequestCloudTypeEnum {
	return ListVaultRequestCloudTypeEnum{
		PUBLIC: ListVaultRequestCloudType{
			value: "public",
		},
		HYBRID: ListVaultRequestCloudType{
			value: "hybrid",
		},
	}
}

func (c ListVaultRequestCloudType) Value() string {
	return c.value
}

func (c ListVaultRequestCloudType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVaultRequestCloudType) UnmarshalJSON(b []byte) error {
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

type ListVaultRequestProtectType struct {
	value string
}

type ListVaultRequestProtectTypeEnum struct {
	BACKUP      ListVaultRequestProtectType
	REPLICATION ListVaultRequestProtectType
}

func GetListVaultRequestProtectTypeEnum() ListVaultRequestProtectTypeEnum {
	return ListVaultRequestProtectTypeEnum{
		BACKUP: ListVaultRequestProtectType{
			value: "backup",
		},
		REPLICATION: ListVaultRequestProtectType{
			value: "replication",
		},
	}
}

func (c ListVaultRequestProtectType) Value() string {
	return c.value
}

func (c ListVaultRequestProtectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVaultRequestProtectType) UnmarshalJSON(b []byte) error {
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
