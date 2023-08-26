package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type MigProject struct {
	Id *string `json:"id,omitempty"`

	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Isdefault *bool `json:"isdefault,omitempty"`

	Template *TemplateResponseBody `json:"template,omitempty"`

	Region string `json:"region"`

	StartTargetServer *bool `json:"start_target_server,omitempty"`

	SpeedLimit *int32 `json:"speed_limit,omitempty"`

	UsePublicIp bool `json:"use_public_ip"`

	ExistServer bool `json:"exist_server"`

	Type MigProjectType `json:"type"`

	EnterpriseProject *string `json:"enterprise_project,omitempty"`

	Syncing bool `json:"syncing"`

	StartNetworkCheck *bool `json:"start_network_check,omitempty"`
}

func (o MigProject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigProject struct{}"
	}

	return strings.Join([]string{"MigProject", string(data)}, " ")
}

type MigProjectType struct {
	value string
}

type MigProjectTypeEnum struct {
	MIGRATE_BLOCK MigProjectType
	MIGRATE_FILE  MigProjectType
}

func GetMigProjectTypeEnum() MigProjectTypeEnum {
	return MigProjectTypeEnum{
		MIGRATE_BLOCK: MigProjectType{
			value: "MIGRATE_BLOCK",
		},
		MIGRATE_FILE: MigProjectType{
			value: "MIGRATE_FILE",
		},
	}
}

func (c MigProjectType) Value() string {
	return c.value
}

func (c MigProjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MigProjectType) UnmarshalJSON(b []byte) error {
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
