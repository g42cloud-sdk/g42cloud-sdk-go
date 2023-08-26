package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type MigprojectsResponseBody struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	UsePublicIp *bool `json:"use_public_ip,omitempty"`

	Isdefault *bool `json:"isdefault,omitempty"`

	StartTargetServer *bool `json:"start_target_server,omitempty"`

	Region *string `json:"region,omitempty"`

	SpeedLimit *int32 `json:"speed_limit,omitempty"`

	ExistServer *bool `json:"exist_server,omitempty"`

	Description *string `json:"description,omitempty"`

	Type *MigprojectsResponseBodyType `json:"type,omitempty"`

	EnterpriseProject *string `json:"enterprise_project,omitempty"`

	Syncing *bool `json:"syncing,omitempty"`

	StartNetworkCheck *bool `json:"start_network_check,omitempty"`
}

func (o MigprojectsResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigprojectsResponseBody struct{}"
	}

	return strings.Join([]string{"MigprojectsResponseBody", string(data)}, " ")
}

type MigprojectsResponseBodyType struct {
	value string
}

type MigprojectsResponseBodyTypeEnum struct {
	MIGRATE_BLOCK MigprojectsResponseBodyType
	MIGRATE_FILE  MigprojectsResponseBodyType
}

func GetMigprojectsResponseBodyTypeEnum() MigprojectsResponseBodyTypeEnum {
	return MigprojectsResponseBodyTypeEnum{
		MIGRATE_BLOCK: MigprojectsResponseBodyType{
			value: "MIGRATE_BLOCK",
		},
		MIGRATE_FILE: MigprojectsResponseBodyType{
			value: "MIGRATE_FILE",
		},
	}
}

func (c MigprojectsResponseBodyType) Value() string {
	return c.value
}

func (c MigprojectsResponseBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MigprojectsResponseBodyType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
